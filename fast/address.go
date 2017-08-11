// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *
 * address.go
 *
 *  Created on Apr 05, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"
	"unsafe"

	"github.com/cosmos72/gomacro/base"
	xr "github.com/cosmos72/gomacro/xreflect"
)

func (c *Comp) AddressOf(node *ast.UnaryExpr) *Expr { return c.addressOf(node.X) }
func (c *Comp) addressOf(expr ast.Expr) *Expr {
	for {
		switch e := expr.(type) {
		case *ast.ParenExpr:
			expr = e.X
			continue
		case *ast.StarExpr:

			ret := c.Expr1(e.X)
			if ret.Type.Kind() != r.Ptr {
				c.Errorf("unary operation * on non-pointer <%v>: %v", ret.Type, e)
			}

		}
		break
	}
	place := c.placeOrAddress(expr, PlaceAddress)

	if place.IsVar() {
		va := place.Var
		return va.Address(c.Depth, c.CompileOptions)
	} else if place.Addr == nil {
		c.Errorf("cannot take the address of %v <%v>", expr, place.Type)
		return nil
	} else {

		t := c.Universe.PtrTo(place.Type)
		return exprX1(t, place.Addr)
	}
}
func (c *Comp) AddressOfVar(name string) *Expr {
	sym := c.Resolve(name)
	va := sym.AsVar(PlaceAddress)
	return va.Address(c.Depth, c.CompileOptions)
}
func (va *Var) Address(maxdepth int, opts CompileOptions) *Expr {
	upn := va.Upn
	k := va.Type.Kind()
	index := va.Desc.Index()
	if index == NoIndex {
		base.Errorf("cannot take the address of %s: _", va.Desc.Class())
		return nil
	}
	var ret I
	compiled := opts&OptIsCompiled != 0
	switch upn {
	case 0:
		switch k {
		case r.Bool:

			if compiled {
				ret = func(env *Env) *bool {
					return env.Binds[index].Addr().Interface().(*bool)
				}
			} else {
				ret = func(env *Env) *bool {
					env.AddressTaken = true
					return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int:

			if compiled {
				ret = func(env *Env) *int {
					return env.Binds[index].Addr().Interface().(*int)
				}
			} else {
				ret = func(env *Env) *int {
					env.AddressTaken = true
					return (*int)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int8:

			if compiled {
				ret = func(env *Env) *int8 {
					return env.Binds[index].Addr().Interface().(*int8)
				}
			} else {
				ret = func(env *Env) *int8 {
					env.AddressTaken = true
					return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int16:

			if compiled {
				ret = func(env *Env) *int16 {
					return env.Binds[index].Addr().Interface().(*int16)
				}
			} else {
				ret = func(env *Env) *int16 {
					env.AddressTaken = true
					return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int32:

			if compiled {
				ret = func(env *Env) *int32 {
					return env.Binds[index].Addr().Interface().(*int32)
				}
			} else {
				ret = func(env *Env) *int32 {
					env.AddressTaken = true
					return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int64:

			if compiled {
				ret = func(env *Env) *int64 {
					return env.Binds[index].Addr().Interface().(*int64)
				}
			} else {
				ret = func(env *Env) *int64 {
					env.AddressTaken = true
					return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint:

			if compiled {
				ret = func(env *Env) *uint {
					return env.Binds[index].Addr().Interface().(*uint)
				}
			} else {
				ret = func(env *Env) *uint {
					env.AddressTaken = true
					return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint8:

			if compiled {
				ret = func(env *Env) *uint8 {
					return env.Binds[index].Addr().Interface().(*uint8)
				}
			} else {
				ret = func(env *Env) *uint8 {
					env.AddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint16:

			if compiled {
				ret = func(env *Env) *uint16 {
					return env.Binds[index].Addr().Interface().(*uint16)
				}
			} else {
				ret = func(env *Env) *uint16 {
					env.AddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint32:

			if compiled {
				ret = func(env *Env) *uint32 {
					return env.Binds[index].Addr().Interface().(*uint32)
				}
			} else {
				ret = func(env *Env) *uint32 {
					env.AddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint64:

			if compiled {
				ret = func(env *Env) *uint64 {
					return env.Binds[index].Addr().Interface().(*uint64)
				}
			} else {
				ret = func(env *Env) *uint64 {
					env.AddressTaken = true
					return &env.IntBinds[index]

				}
			}
		case r.Uintptr:

			if compiled {
				ret = func(env *Env) *uintptr {
					return env.Binds[index].Addr().Interface().(*uintptr)
				}
			} else {
				ret = func(env *Env) *uintptr {
					env.AddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float32:

			if compiled {
				ret = func(env *Env) *float32 {
					return env.Binds[index].Addr().Interface().(*float32)
				}
			} else {
				ret = func(env *Env) *float32 {
					env.AddressTaken = true
					return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float64:

			if compiled {
				ret = func(env *Env) *float64 {
					return env.Binds[index].Addr().Interface().(*float64)
				}
			} else {
				ret = func(env *Env) *float64 {
					env.AddressTaken = true
					return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Complex64:

			if compiled {
				ret = func(env *Env) *complex64 {
					return env.Binds[index].Addr().Interface().(*complex64)
				}
			} else {
				ret = func(env *Env) *complex64 {
					env.AddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		default:

			if compiled {
				ret = func(env *Env) r.Value {
					return env.Binds[index].Addr().Interface().(r.Value)
				}
			} else {
				ret = func(env *Env) r.Value {
					return env.Binds[index].Addr()
				}
			}
		}
	case 1:
		switch k {
		case r.Bool:

			if compiled {
				ret = func(env *Env) *bool {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*bool)
				}
			} else {
				ret = func(env *Env) *bool {
					env = env.
						Outer

					env.AddressTaken = true
					return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int:

			if compiled {
				ret = func(env *Env) *int {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*int)
				}
			} else {
				ret = func(env *Env) *int {
					env = env.
						Outer

					env.AddressTaken = true
					return (*int)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int8:

			if compiled {
				ret = func(env *Env) *int8 {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*int8)
				}
			} else {
				ret = func(env *Env) *int8 {
					env = env.
						Outer

					env.AddressTaken = true
					return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int16:

			if compiled {
				ret = func(env *Env) *int16 {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*int16)
				}
			} else {
				ret = func(env *Env) *int16 {
					env = env.
						Outer

					env.AddressTaken = true
					return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int32:

			if compiled {
				ret = func(env *Env) *int32 {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*int32)
				}
			} else {
				ret = func(env *Env) *int32 {
					env = env.
						Outer

					env.AddressTaken = true
					return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int64:

			if compiled {
				ret = func(env *Env) *int64 {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*int64)
				}
			} else {
				ret = func(env *Env) *int64 {
					env = env.
						Outer

					env.AddressTaken = true
					return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint:

			if compiled {
				ret = func(env *Env) *uint {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*uint)
				}
			} else {
				ret = func(env *Env) *uint {
					env = env.
						Outer

					env.AddressTaken = true
					return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint8:

			if compiled {
				ret = func(env *Env) *uint8 {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*uint8)
				}
			} else {
				ret = func(env *Env) *uint8 {
					env = env.
						Outer

					env.AddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint16:

			if compiled {
				ret = func(env *Env) *uint16 {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*uint16)
				}
			} else {
				ret = func(env *Env) *uint16 {
					env = env.
						Outer

					env.AddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint32:

			if compiled {
				ret = func(env *Env) *uint32 {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*uint32)
				}
			} else {
				ret = func(env *Env) *uint32 {
					env = env.
						Outer

					env.AddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint64:

			if compiled {
				ret = func(env *Env) *uint64 {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*uint64)
				}
			} else {
				ret = func(env *Env) *uint64 {
					env = env.
						Outer

					env.AddressTaken = true
					return &env.IntBinds[index]

				}
			}
		case r.Uintptr:

			if compiled {
				ret = func(env *Env) *uintptr {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*uintptr)
				}
			} else {
				ret = func(env *Env) *uintptr {
					env = env.
						Outer

					env.AddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float32:

			if compiled {
				ret = func(env *Env) *float32 {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*float32)
				}
			} else {
				ret = func(env *Env) *float32 {
					env = env.
						Outer

					env.AddressTaken = true
					return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float64:

			if compiled {
				ret = func(env *Env) *float64 {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*float64)
				}
			} else {
				ret = func(env *Env) *float64 {
					env = env.
						Outer

					env.AddressTaken = true
					return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Complex64:

			if compiled {
				ret = func(env *Env) *complex64 {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(*complex64)
				}
			} else {
				ret = func(env *Env) *complex64 {
					env = env.
						Outer

					env.AddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		default:

			if compiled {
				ret = func(env *Env) r.Value {
					env = env.
						Outer
					return env.Binds[index].Addr().Interface().(r.Value)
				}
			} else {
				ret = func(env *Env) r.Value {
					env = env.
						Outer
					return env.Binds[index].Addr()

				}
			}
		}
	case 2:
		switch k {
		case r.Bool:

			if compiled {
				ret = func(env *Env) *bool {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*bool)
				}
			} else {
				ret = func(env *Env) *bool {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int:

			if compiled {
				ret = func(env *Env) *int {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*int)
				}
			} else {
				ret = func(env *Env) *int {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*int)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int8:

			if compiled {
				ret = func(env *Env) *int8 {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*int8)
				}
			} else {
				ret = func(env *Env) *int8 {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int16:

			if compiled {
				ret = func(env *Env) *int16 {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*int16)
				}
			} else {
				ret = func(env *Env) *int16 {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int32:

			if compiled {
				ret = func(env *Env) *int32 {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*int32)
				}
			} else {
				ret = func(env *Env) *int32 {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int64:

			if compiled {
				ret = func(env *Env) *int64 {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*int64)
				}
			} else {
				ret = func(env *Env) *int64 {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint:

			if compiled {
				ret = func(env *Env) *uint {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*uint)
				}
			} else {
				ret = func(env *Env) *uint {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint8:

			if compiled {
				ret = func(env *Env) *uint8 {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*uint8)
				}
			} else {
				ret = func(env *Env) *uint8 {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint16:

			if compiled {
				ret = func(env *Env) *uint16 {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*uint16)
				}
			} else {
				ret = func(env *Env) *uint16 {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint32:

			if compiled {
				ret = func(env *Env) *uint32 {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*uint32)
				}
			} else {
				ret = func(env *Env) *uint32 {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint64:

			if compiled {
				ret = func(env *Env) *uint64 {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*uint64)
				}
			} else {
				ret = func(env *Env) *uint64 {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return &env.IntBinds[index]

				}
			}
		case r.Uintptr:

			if compiled {
				ret = func(env *Env) *uintptr {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*uintptr)
				}
			} else {
				ret = func(env *Env) *uintptr {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float32:

			if compiled {
				ret = func(env *Env) *float32 {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*float32)
				}
			} else {
				ret = func(env *Env) *float32 {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float64:

			if compiled {
				ret = func(env *Env) *float64 {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*float64)
				}
			} else {
				ret = func(env *Env) *float64 {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Complex64:

			if compiled {
				ret = func(env *Env) *complex64 {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(*complex64)
				}
			} else {
				ret = func(env *Env) *complex64 {
					env = env.
						Outer.Outer

					env.AddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		default:

			if compiled {
				ret = func(env *Env) r.Value {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr().Interface().(r.Value)
				}
			} else {
				ret = func(env *Env) r.Value {
					env = env.
						Outer.Outer
					return env.Binds[index].Addr()

				}
			}
		}
	default:
		switch k {
		case r.Bool:

			if compiled {
				ret = func(env *Env) *bool {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*bool)
				}
			} else {
				ret = func(env *Env) *bool {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int:

			if compiled {
				ret = func(env *Env) *int {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*int)
				}
			} else {
				ret = func(env *Env) *int {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*int)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int8:

			if compiled {
				ret = func(env *Env) *int8 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*int8)
				}
			} else {
				ret = func(env *Env) *int8 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int16:

			if compiled {
				ret = func(env *Env) *int16 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*int16)
				}
			} else {
				ret = func(env *Env) *int16 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int32:

			if compiled {
				ret = func(env *Env) *int32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*int32)
				}
			} else {
				ret = func(env *Env) *int32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int64:

			if compiled {
				ret = func(env *Env) *int64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*int64)
				}
			} else {
				ret = func(env *Env) *int64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint:

			if compiled {
				ret = func(env *Env) *uint {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*uint)
				}
			} else {
				ret = func(env *Env) *uint {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint8:

			if compiled {
				ret = func(env *Env) *uint8 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*uint8)
				}
			} else {
				ret = func(env *Env) *uint8 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint16:

			if compiled {
				ret = func(env *Env) *uint16 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*uint16)
				}
			} else {
				ret = func(env *Env) *uint16 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint32:

			if compiled {
				ret = func(env *Env) *uint32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*uint32)
				}
			} else {
				ret = func(env *Env) *uint32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint64:

			if compiled {
				ret = func(env *Env) *uint64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*uint64)
				}
			} else {
				ret = func(env *Env) *uint64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return &env.IntBinds[index]

				}
			}
		case r.Uintptr:

			if compiled {
				ret = func(env *Env) *uintptr {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*uintptr)
				}
			} else {
				ret = func(env *Env) *uintptr {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float32:

			if compiled {
				ret = func(env *Env) *float32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*float32)
				}
			} else {
				ret = func(env *Env) *float32 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float64:

			if compiled {
				ret = func(env *Env) *float64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*float64)
				}
			} else {
				ret = func(env *Env) *float64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Complex64:

			if compiled {
				ret = func(env *Env) *complex64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(*complex64)
				}
			} else {
				ret = func(env *Env) *complex64 {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}

					env.AddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		default:

			if compiled {
				ret = func(env *Env) r.Value {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr().Interface().(r.Value)
				}
			} else {
				ret = func(env *Env) r.Value {
					env = env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						env = env.Outer
					}
					return env.Binds[index].Addr()

				}
			}
		}
	case maxdepth - 1:
		switch k {
		case r.Bool:

			if compiled {
				ret = func(env *Env) *bool {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*bool)
				}
			} else {
				ret = func(env *Env) *bool {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int:

			if compiled {
				ret = func(env *Env) *int {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*int)
				}
			} else {
				ret = func(env *Env) *int {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*int)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int8:

			if compiled {
				ret = func(env *Env) *int8 {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*int8)
				}
			} else {
				ret = func(env *Env) *int8 {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int16:

			if compiled {
				ret = func(env *Env) *int16 {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*int16)
				}
			} else {
				ret = func(env *Env) *int16 {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int32:

			if compiled {
				ret = func(env *Env) *int32 {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*int32)
				}
			} else {
				ret = func(env *Env) *int32 {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int64:

			if compiled {
				ret = func(env *Env) *int64 {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*int64)
				}
			} else {
				ret = func(env *Env) *int64 {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint:

			if compiled {
				ret = func(env *Env) *uint {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*uint)
				}
			} else {
				ret = func(env *Env) *uint {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint8:

			if compiled {
				ret = func(env *Env) *uint8 {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*uint8)
				}
			} else {
				ret = func(env *Env) *uint8 {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint16:

			if compiled {
				ret = func(env *Env) *uint16 {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*uint16)
				}
			} else {
				ret = func(env *Env) *uint16 {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint32:

			if compiled {
				ret = func(env *Env) *uint32 {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*uint32)
				}
			} else {
				ret = func(env *Env) *uint32 {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint64:

			if compiled {
				ret = func(env *Env) *uint64 {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*uint64)
				}
			} else {
				ret = func(env *Env) *uint64 {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return &env.IntBinds[index]

				}
			}
		case r.Uintptr:

			if compiled {
				ret = func(env *Env) *uintptr {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*uintptr)
				}
			} else {
				ret = func(env *Env) *uintptr {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float32:

			if compiled {
				ret = func(env *Env) *float32 {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*float32)
				}
			} else {
				ret = func(env *Env) *float32 {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float64:

			if compiled {
				ret = func(env *Env) *float64 {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*float64)
				}
			} else {
				ret = func(env *Env) *float64 {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Complex64:

			if compiled {
				ret = func(env *Env) *complex64 {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(*complex64)
				}
			} else {
				ret = func(env *Env) *complex64 {
					env = env.ThreadGlobals.FileEnv

					env.AddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		default:

			if compiled {
				ret = func(env *Env) r.Value {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr().Interface().(r.Value)
				}
			} else {
				ret = func(env *Env) r.Value {
					env = env.ThreadGlobals.FileEnv
					return env.Binds[index].Addr()

				}
			}
		}
	case maxdepth:
		switch k {
		case r.Bool:

			if compiled {
				ret = func(env *Env) *bool {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*bool)
				}
			} else {
				ret = func(env *Env) *bool {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int:

			if compiled {
				ret = func(env *Env) *int {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*int)
				}
			} else {
				ret = func(env *Env) *int {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*int)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int8:

			if compiled {
				ret = func(env *Env) *int8 {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*int8)
				}
			} else {
				ret = func(env *Env) *int8 {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int16:

			if compiled {
				ret = func(env *Env) *int16 {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*int16)
				}
			} else {
				ret = func(env *Env) *int16 {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int32:

			if compiled {
				ret = func(env *Env) *int32 {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*int32)
				}
			} else {
				ret = func(env *Env) *int32 {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Int64:

			if compiled {
				ret = func(env *Env) *int64 {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*int64)
				}
			} else {
				ret = func(env *Env) *int64 {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint:

			if compiled {
				ret = func(env *Env) *uint {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*uint)
				}
			} else {
				ret = func(env *Env) *uint {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint8:

			if compiled {
				ret = func(env *Env) *uint8 {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*uint8)
				}
			} else {
				ret = func(env *Env) *uint8 {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint16:

			if compiled {
				ret = func(env *Env) *uint16 {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*uint16)
				}
			} else {
				ret = func(env *Env) *uint16 {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint32:

			if compiled {
				ret = func(env *Env) *uint32 {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*uint32)
				}
			} else {
				ret = func(env *Env) *uint32 {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Uint64:

			if compiled {
				ret = func(env *Env) *uint64 {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*uint64)
				}
			} else {
				ret = func(env *Env) *uint64 {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return &env.IntBinds[index]

				}
			}
		case r.Uintptr:

			if compiled {
				ret = func(env *Env) *uintptr {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*uintptr)
				}
			} else {
				ret = func(env *Env) *uintptr {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float32:

			if compiled {
				ret = func(env *Env) *float32 {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*float32)
				}
			} else {
				ret = func(env *Env) *float32 {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Float64:

			if compiled {
				ret = func(env *Env) *float64 {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*float64)
				}
			} else {
				ret = func(env *Env) *float64 {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		case r.Complex64:

			if compiled {
				ret = func(env *Env) *complex64 {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(*complex64)
				}
			} else {
				ret = func(env *Env) *complex64 {
					env = env.ThreadGlobals.TopEnv

					env.AddressTaken = true
					return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

				}
			}
		default:

			if compiled {
				ret = func(env *Env) r.Value {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr().Interface().(r.Value)
				}
			} else {
				ret = func(env *Env) r.Value {
					env = env.ThreadGlobals.TopEnv
					return env.Binds[index].Addr()

				}
			}
		}
	}
	return &Expr{Lit: Lit{Type: xr.PtrTo(va.Type)}, Fun: ret}
}
