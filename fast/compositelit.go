/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * compositelit.go
 *
 *  Created on May 28, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) CompositeLit(node *ast.CompositeLit) *Expr {
	t, ellipsis := c.compileType2(node.Type, true)
	// array and slice: "index must be non-negative integer constant"
	switch t.Kind() {
	case r.Map:
		return c.compositeLitMap(t, node)
	case r.Array, r.Slice:
		return c.compositeLitSlice(t, ellipsis, node)
	case r.Struct:
		return c.compositeLitStruct(t, node)
	default:
		c.Errorf("invalid type for composite literal: <%v> %v", t, node.Type)
		return nil
	}
}

func (c *Comp) compositeLitMap(t xr.Type, node *ast.CompositeLit) *Expr {
	rtype := t.ReflectType()
	n := len(node.Elts)
	if n == 0 {
		return exprX1(t, func(env *Env) r.Value {
			return r.MakeMap(rtype)
		})
	}
	tkey := t.Key()
	tval := t.Elem()

	seen := make(map[interface{}]bool) // constant keys already seen
	funkeys := make([]func(*Env) r.Value, n)
	funvals := make([]func(*Env) r.Value, n)

	for i, el := range node.Elts {
		switch elkv := el.(type) {
		case *ast.KeyValueExpr:
			ekey := c.Expr1(elkv.Key)
			eval := c.Expr1(elkv.Value)
			if ekey.Const() {
				ekey.ConstTo(tkey)
				if seen[ekey.Value] {
					c.Errorf("duplicate key %v in map literal", elkv.Key)
				}
				seen[ekey.Value] = true
			} else if !ekey.Type.AssignableTo(tkey) {
				c.Errorf("cannot use %v <%v> as type <%v> in map key", elkv.Key, ekey.Type, tkey)
			}
			if eval.Const() {
				eval.ConstTo(tval)
			} else if !eval.Type.AssignableTo(tval) {
				c.Errorf("cannot use %v <%v> as type <%v> in map value", elkv.Value, eval.Type, tval)
			}
			funkeys[i] = ekey.AsX1()
			funvals[i] = eval.AsX1()

		default:
			c.Errorf("missing key in map literal: %v", el)
		}
	}
	rtkey, rtval := rtype.Key(), rtype.Elem()
	zerokey, zeroval := r.Zero(rtkey), r.Zero(rtval)
	return exprX1(t, func(env *Env) r.Value {
		obj := r.MakeMap(rtype)
		var key, val r.Value
		for i, funkey := range funkeys {
			key = funkey(env)
			if key == Nil || key == None {
				key = zerokey
			} else if key.Type() != rtkey {
				key = key.Convert(rtkey)
			}
			val = funvals[i](env)
			if val == Nil || val == None {
				val = zeroval
			} else if val.Type() != rtval {
				val = val.Convert(rtval)
			}
			obj.SetMapIndex(key, val)
		}
		return obj
	})
}

func (c *Comp) compositeLitSlice(t xr.Type, ellipsis bool, node *ast.CompositeLit) *Expr {
	c.Errorf("unimplemented: array/slice composite literal: %v", node)
	return nil
}

func (c *Comp) compositeLitStruct(t xr.Type, node *ast.CompositeLit) *Expr {
	rtype := t.ReflectType()
	n := len(node.Elts)
	if n == 0 {
		return exprX1(t, func(env *Env) r.Value {
			return r.New(rtype).Elem()
		})
	}

	var seen map[string]bool
	var all map[string]xr.StructField
	inits := make([]func(*Env) r.Value, n)
	indexes := make([]int, n)
	var flagkv, flagv bool

	for i, el := range node.Elts {
		switch elkv := el.(type) {
		case *ast.KeyValueExpr:
			flagkv = true
			if flagv {
				c.Errorf("mixture of field:value and value in struct literal: %v", node)
			}
			switch k := elkv.Key.(type) {
			case *ast.Ident:
				name := k.Name
				if seen[name] {
					c.Errorf("duplicate field name in struct literal: %v", name)
				} else if seen == nil {
					seen = make(map[string]bool)
					all = listStructFields(t, c.PackagePath)
				}
				field, ok := all[name]
				if !ok {
					c.Errorf("unknown field '%v' in struct literal of type %v", name, t)
				}
				expr := c.Expr1(elkv.Value)
				if expr.Const() {
					expr.ConstTo(field.Type)
				} else if !expr.Type.AssignableTo(field.Type) {
					c.Errorf("cannot use %v <%v> as type <%v> in field value", elkv.Value, expr.Type, field.Type)
				}
				inits[i] = expr.AsX1()
				indexes[i] = field.Index[0]
			default:
				c.Errorf("invalid field name '%v' in struct literal", k)
			}
		default:
			flagv = true
			if flagkv {
				c.Errorf("mixture of field:value and value in struct literal: %v", node)
			}
			field := t.Field(i)
			expr := c.Expr1(el)
			if expr.Const() {
				expr.ConstTo(field.Type)
			} else if !expr.Type.AssignableTo(field.Type) {
				c.Errorf("cannot use %v <%v> as type <%v> in field value", el, expr.Type, field.Type)
			}
			if !ast.IsExported(field.Name) && field.Pkg.Path() != c.PackagePath {
				c.Errorf("implicit assignment of unexported field '%v' in struct literal <%v>", field.Name, t)
			}
			inits[i] = expr.AsX1()
			indexes[i] = field.Index[0]
		}
	}
	if nfield := t.NumField(); flagv && n != nfield {
		var label, plural = "few", "s"
		if n > nfield {
			label = "many"
		} else if n == 1 {
			plural = ""
		}
		c.Errorf("too %s values in struct initializer: <%v> has %d fields, found %d initializer%s",
			label, t, nfield, n, plural)
	}
	return exprX1(t, func(env *Env) r.Value {
		obj := r.New(rtype).Elem()
		var val, field r.Value
		var tfield r.Type
		for i, init := range inits {
			val = init(env)
			if val == Nil || val == None {
				continue
			}
			field = obj.Field(indexes[i])
			tfield = field.Type()
			if val.Type() != tfield {
				val = val.Convert(tfield)
			}
			field.Set(val)
		}
		return obj
	})
}

// listStructFields lists the field names of a struct. It ignores embedded fields.
// Unexported fields are listed only if their package's path matches given pkgpath
func listStructFields(t xr.Type, pkgpath string) map[string]xr.StructField {
	list := make(map[string]xr.StructField)
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		if ast.IsExported(f.Name) || f.Pkg.Path() == pkgpath {
			list[f.Name] = f
		}
	}
	return list
}
