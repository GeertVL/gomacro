// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
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
)

func (c *Comp) AddressOf(node *ast.UnaryExpr) *Expr {
	place := c.PlaceOrAddress(node.X, true)
	if place.Fun != nil {
		return c.badUnaryExpr("unimplemented: address of non-identifier", node, nil)
	}
	return place.Var.Address(c.Depth)
}
func (c *Comp) AddressOfVar(name string) *Expr {
	upn, bind := c.Resolve(name)
	class := bind.Desc.Class()
	switch class {
	default:
		c.Errorf("cannot take the address of %s: %v", class, name)
		return nil
	case VarBind, IntBind:
		va := bind.AsVar(upn)
		return va.Address(c.Depth)
	}
}
func (va *Var) Address(maxdepth int) *Expr {
	upn := va.Upn
	k := va.Type.Kind()
	index := va.Desc.Index()
	if index == NoIndex {
		base.Errorf("cannot take the address of %s: _", va.Desc.Class())
		return nil
	}
	var ret I
	switch upn {
	case 0:
		switch k {
		case r.Bool:

			ret = func(env *Env) *bool {

				env.AddressTaken = true
				return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int:

			ret = func(env *Env) *int {

				env.AddressTaken = true
				return (*int)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int8:

			ret = func(env *Env) *int8 {

				env.AddressTaken = true
				return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int16:

			ret = func(env *Env) *int16 {

				env.AddressTaken = true
				return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int32:

			ret = func(env *Env) *int32 {

				env.AddressTaken = true
				return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int64:

			ret = func(env *Env) *int64 {

				env.AddressTaken = true
				return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint:

			ret = func(env *Env) *uint {

				env.AddressTaken = true
				return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint8:

			ret = func(env *Env) *uint8 {

				env.AddressTaken = true
				return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint16:

			ret = func(env *Env) *uint16 {

				env.AddressTaken = true
				return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint32:

			ret = func(env *Env) *uint32 {

				env.AddressTaken = true
				return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint64:

			ret = func(env *Env) *uint64 {

				env.AddressTaken = true
				return &env.IntBinds[index]

			}
		case r.Uintptr:

			ret = func(env *Env) *uintptr {

				env.AddressTaken = true
				return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float32:

			ret = func(env *Env) *float32 {

				env.AddressTaken = true
				return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float64:

			ret = func(env *Env) *float64 {

				env.AddressTaken = true
				return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Complex64:

			ret = func(env *Env) *complex64 {

				env.AddressTaken = true
				return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		default:

			ret = func(env *Env) r.Value {

				env.AddressTaken = true
				return env.Binds[index].Addr()

			}
		}
	case 1:
		switch k {
		case r.Bool:

			ret = func(env *Env) *bool {
				env = env.
					Outer

				env.AddressTaken = true
				return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int:

			ret = func(env *Env) *int {
				env = env.
					Outer

				env.AddressTaken = true
				return (*int)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int8:

			ret = func(env *Env) *int8 {
				env = env.
					Outer

				env.AddressTaken = true
				return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int16:

			ret = func(env *Env) *int16 {
				env = env.
					Outer

				env.AddressTaken = true
				return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int32:

			ret = func(env *Env) *int32 {
				env = env.
					Outer

				env.AddressTaken = true
				return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int64:

			ret = func(env *Env) *int64 {
				env = env.
					Outer

				env.AddressTaken = true
				return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint:

			ret = func(env *Env) *uint {
				env = env.
					Outer

				env.AddressTaken = true
				return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint8:

			ret = func(env *Env) *uint8 {
				env = env.
					Outer

				env.AddressTaken = true
				return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint16:

			ret = func(env *Env) *uint16 {
				env = env.
					Outer

				env.AddressTaken = true
				return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint32:

			ret = func(env *Env) *uint32 {
				env = env.
					Outer

				env.AddressTaken = true
				return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint64:

			ret = func(env *Env) *uint64 {
				env = env.
					Outer

				env.AddressTaken = true
				return &env.IntBinds[index]

			}
		case r.Uintptr:

			ret = func(env *Env) *uintptr {
				env = env.
					Outer

				env.AddressTaken = true
				return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float32:

			ret = func(env *Env) *float32 {
				env = env.
					Outer

				env.AddressTaken = true
				return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float64:

			ret = func(env *Env) *float64 {
				env = env.
					Outer

				env.AddressTaken = true
				return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Complex64:

			ret = func(env *Env) *complex64 {
				env = env.
					Outer

				env.AddressTaken = true
				return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		default:

			ret = func(env *Env) r.Value {
				env = env.
					Outer

				env.AddressTaken = true
				return env.Binds[index].Addr()

			}
		}
	case 2:
		switch k {
		case r.Bool:

			ret = func(env *Env) *bool {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int:

			ret = func(env *Env) *int {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*int)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int8:

			ret = func(env *Env) *int8 {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int16:

			ret = func(env *Env) *int16 {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int32:

			ret = func(env *Env) *int32 {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int64:

			ret = func(env *Env) *int64 {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint:

			ret = func(env *Env) *uint {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint8:

			ret = func(env *Env) *uint8 {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint16:

			ret = func(env *Env) *uint16 {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint32:

			ret = func(env *Env) *uint32 {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint64:

			ret = func(env *Env) *uint64 {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return &env.IntBinds[index]

			}
		case r.Uintptr:

			ret = func(env *Env) *uintptr {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float32:

			ret = func(env *Env) *float32 {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float64:

			ret = func(env *Env) *float64 {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Complex64:

			ret = func(env *Env) *complex64 {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		default:

			ret = func(env *Env) r.Value {
				env = env.
					Outer.Outer

				env.AddressTaken = true
				return env.Binds[index].Addr()

			}
		}
	default:
		switch k {
		case r.Bool:

			ret = func(env *Env) *bool {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int:

			ret = func(env *Env) *int {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*int)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int8:

			ret = func(env *Env) *int8 {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int16:

			ret = func(env *Env) *int16 {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int32:

			ret = func(env *Env) *int32 {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int64:

			ret = func(env *Env) *int64 {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint:

			ret = func(env *Env) *uint {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint8:

			ret = func(env *Env) *uint8 {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint16:

			ret = func(env *Env) *uint16 {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint32:

			ret = func(env *Env) *uint32 {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint64:

			ret = func(env *Env) *uint64 {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return &env.IntBinds[index]

			}
		case r.Uintptr:

			ret = func(env *Env) *uintptr {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float32:

			ret = func(env *Env) *float32 {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float64:

			ret = func(env *Env) *float64 {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Complex64:

			ret = func(env *Env) *complex64 {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		default:

			ret = func(env *Env) r.Value {
				env = env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					env = env.Outer
				}

				env.AddressTaken = true
				return env.Binds[index].Addr()

			}
		}
	case maxdepth - 1:
		switch k {
		case r.Bool:

			ret = func(env *Env) *bool {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int:

			ret = func(env *Env) *int {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*int)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int8:

			ret = func(env *Env) *int8 {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int16:

			ret = func(env *Env) *int16 {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int32:

			ret = func(env *Env) *int32 {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int64:

			ret = func(env *Env) *int64 {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint:

			ret = func(env *Env) *uint {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint8:

			ret = func(env *Env) *uint8 {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint16:

			ret = func(env *Env) *uint16 {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint32:

			ret = func(env *Env) *uint32 {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint64:

			ret = func(env *Env) *uint64 {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return &env.IntBinds[index]

			}
		case r.Uintptr:

			ret = func(env *Env) *uintptr {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float32:

			ret = func(env *Env) *float32 {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float64:

			ret = func(env *Env) *float64 {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Complex64:

			ret = func(env *Env) *complex64 {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		default:

			ret = func(env *Env) r.Value {
				env = env.ThreadGlobals.FileEnv

				env.AddressTaken = true
				return env.Binds[index].Addr()

			}
		}
	case maxdepth:
		switch k {
		case r.Bool:

			ret = func(env *Env) *bool {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*bool)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int:

			ret = func(env *Env) *int {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*int)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int8:

			ret = func(env *Env) *int8 {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*int8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int16:

			ret = func(env *Env) *int16 {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*int16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int32:

			ret = func(env *Env) *int32 {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*int32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Int64:

			ret = func(env *Env) *int64 {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*int64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint:

			ret = func(env *Env) *uint {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*uint)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint8:

			ret = func(env *Env) *uint8 {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*uint8)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint16:

			ret = func(env *Env) *uint16 {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*uint16)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint32:

			ret = func(env *Env) *uint32 {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*uint32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Uint64:

			ret = func(env *Env) *uint64 {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return &env.IntBinds[index]

			}
		case r.Uintptr:

			ret = func(env *Env) *uintptr {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*uintptr)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float32:

			ret = func(env *Env) *float32 {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*float32)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Float64:

			ret = func(env *Env) *float64 {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*float64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		case r.Complex64:

			ret = func(env *Env) *complex64 {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return (*complex64)(unsafe.Pointer(&env.IntBinds[index]))

			}
		default:

			ret = func(env *Env) r.Value {
				env = env.ThreadGlobals.TopEnv

				env.AddressTaken = true
				return env.Binds[index].Addr()

			}
		}
	}
	return &Expr{Lit: Lit{Type: r.PtrTo(va.Type)}, Fun: ret}
}
