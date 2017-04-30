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
 * callnret0.go
 *
 *  Created on Apr 15, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"

	. "github.com/cosmos72/gomacro/base"
)

func call0ret0(c *Call, maxdepth int) func(env *Env) {
	expr := c.Fun
	funsym := expr.Sym
	if funsym == nil {
		exprfun := expr.AsX1()
		return func(env *Env) {
			fun := exprfun(env).Interface().(func())
			fun()
		}
	}

	var cachedfunv r.Value
	var cachedfun func()

	funupn := funsym.Upn
	funindex := funsym.Desc.Index()
	switch funupn {
	case maxdepth - 1:
		return func(env *Env) {
			funv := env.ThreadGlobals.FileEnv.Binds[funindex]
			if cachedfunv != funv {
				cachedfunv = funv
				cachedfun = funv.Interface().(func())
			}
			cachedfun()
		}
	case 0:
		return func(env *Env) {
			fun := env.Binds[funindex].Interface().(func())
			fun()
		}
	case 1:
		return func(env *Env) {
			fun := env.Outer.Binds[funindex].Interface().(func())
			fun()
		}
	case 2:
		return func(env *Env) {
			fun := env.Outer.Outer.Binds[funindex].Interface().(func())
			fun()
		}
	default:
		return func(env *Env) {
			env = env.Outer.Outer.Outer.Outer
			for i := 3; i < funupn; i++ {
				env = env.Outer
			}

			fun := env.Binds[funindex].Interface().(func())
			fun()
		}
	}
}
func call1ret0(c *Call, maxdepth int) func(env *Env) {
	expr := c.Fun
	exprfun := expr.AsX1()
	funsym := expr.Sym
	funupn, funindex := -1, -1
	if funsym != nil {
		funupn = funsym.Upn
		funindex = funsym.Desc.Index()
		if funindex == NoIndex {
			Errorf("internal error: call1ret0() invoked for constant function %#v. use call_builtin() instead", expr)
		}

	}
	args := c.Args
	argfunsX1 := c.MakeArgfunsX1()

	var cachedfunv r.Value
	var call func(env *Env)

	t := expr.Type.In(0)
	switch t.Kind() {
	case r.Bool:

		if args[0].Const() {
			argconst := args[0].Value.(bool)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(bool)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(bool))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(bool))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) bool)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(bool)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(bool))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(bool))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Int:

		if args[0].Const() {
			argconst := args[0].Value.(int)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(int))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) int)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(int))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Int8:

		if args[0].Const() {
			argconst := args[0].Value.(int8)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int8)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(int8))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int8))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) int8)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int8)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(int8))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int8))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Int16:
		if args[0].Const() {
			argconst := args[0].Value.(int16)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int16)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(int16))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int16))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) int16)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int16)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(int16))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int16))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Int32:
		if args[0].Const() {
			argconst := args[0].Value.(int32)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int32)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(int32))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int32))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) int32)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int32)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(int32))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int32))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Int64:
		if args[0].Const() {
			argconst := args[0].Value.(int64)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int64)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(int64))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int64))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) int64)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(int64)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(int64))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(int64))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Uint:
		if args[0].Const() {
			argconst := args[0].Value.(uint)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uint))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) uint)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uint))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Uint8:
		if args[0].Const() {
			argconst := args[0].Value.(uint8)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint8)
				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uint8))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint8))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) uint8)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint8)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uint8))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint8))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Uint16:
		if args[0].Const() {
			argconst := args[0].Value.(uint16)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint16)
				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uint16))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint16))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) uint16)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint16)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uint16))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint16))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Uint32:
		if args[0].Const() {
			argconst := args[0].Value.(uint32)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint32)
				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uint32))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint32))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) uint32)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint32)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uint32))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint32))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Uint64:
		if args[0].Const() {
			argconst := args[0].Value.(uint64)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint64)
				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uint64))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint64))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) uint64)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uint64)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uint64))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uint64))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Uintptr:
		if args[0].Const() {
			argconst := args[0].Value.(uintptr)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uintptr)
				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uintptr))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uintptr))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) uintptr)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(uintptr)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(uintptr))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(uintptr))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Float32:
		if args[0].Const() {
			argconst := args[0].Value.(float32)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float32)
				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(float32))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(float32))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) float32)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float32)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(float32))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(float32))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Float64:
		if args[0].Const() {
			argconst := args[0].Value.(float64)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float64)
				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(float64))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(float64))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) float64)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(float64)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(float64))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(float64))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Complex64:
		if args[0].Const() {
			argconst := args[0].Value.(complex64)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex64)
				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(complex64))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(complex64))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) complex64)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex64)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(complex64))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(complex64))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.Complex128:
		if args[0].Const() {
			argconst := args[0].Value.(complex128)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex128)
				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(complex128))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(complex128))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) complex128)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(complex128)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(complex128))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(complex128))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	case r.String:
		if args[0].Const() {
			argconst := args[0].Value.(string)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(string)
				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(string))
					}

					cachedfun(argconst)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(string))

					fun(argconst)
				}
			}

		} else {
			argfun := args[0].Fun.(func(env *Env) string)
			if funsym != nil && funupn == maxdepth-1 {
				var cachedfun func(string)

				call = func(env *Env) {
					funv := env.ThreadGlobals.FileEnv.Binds[funindex]
					if cachedfunv != funv {
						cachedfunv = funv
						cachedfun = funv.Interface().(func(string))
					}

					arg := argfun(env)

					cachedfun(arg)
				}
			} else {
				call = func(env *Env) {
					fun := exprfun(env).Interface().(func(string))
					arg := argfun(env)

					fun(arg)
				}
			}

		}
	default:
		{
			argfun := argfunsX1[0]
			call = func(env *Env) {
				funv := exprfun(env)

				argv := []r.Value{
					argfun(env),
				}
				funv.Call(argv)
			}
		}

	}
	return call
}
func call2ret0(c *Call, maxdepth int) func(env *Env) {
	expr := c.Fun
	exprfun := expr.AsX1()
	funsym := expr.Sym
	funupn, funindex := -1, -1
	if funsym != nil {
		funupn = funsym.Upn
		funindex = funsym.Desc.Index()
		if funindex == NoIndex {
			Errorf("internal error: call2ret0() invoked for constant function %#v. use call_builtin() instead", expr)
		}

	}
	args := c.Args
	argfunsX1 := c.MakeArgfunsX1()
	var cachedfunv r.Value
	var call func(env *Env)

	if expr.Type.In(0) == expr.Type.In(1) {
		t := expr.Type.In(0)
		switch t.Kind() {
		case r.Bool:

			{
				arg0fun := args[0].WithFun().(func(*Env) bool)
				arg1fun := args[1].WithFun().(func(*Env) bool)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(bool, bool)

					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(bool, bool))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(bool, bool))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Int:

			{
				arg0fun := args[0].WithFun().(func(*Env) int)
				arg1fun := args[1].WithFun().(func(*Env) int)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int, int)

					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int, int))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(int, int))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Int8:
			{
				arg0fun := args[0].WithFun().(func(*Env) int8)
				arg1fun := args[1].WithFun().(func(*Env) int8)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int8, int8)

					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int8, int8))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(int8, int8))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Int16:
			{
				arg0fun := args[0].WithFun().(func(*Env) int16)
				arg1fun := args[1].WithFun().(func(*Env) int16)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int16, int16)

					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int16, int16))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(int16, int16))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Int32:
			{
				arg0fun := args[0].WithFun().(func(*Env) int32)
				arg1fun := args[1].WithFun().(func(*Env) int32)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int32, int32)

					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int32, int32))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(int32, int32))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Int64:
			{
				arg0fun := args[0].WithFun().(func(*Env) int64)
				arg1fun := args[1].WithFun().(func(*Env) int64)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(int64, int64)

					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(int64, int64))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(int64, int64))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uint:
			{
				arg0fun := args[0].WithFun().(func(*Env) uint)
				arg1fun := args[1].WithFun().(func(*Env) uint)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uint, uint)

					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint, uint))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint, uint))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uint8:
			{
				arg0fun := args[0].WithFun().(func(*Env) uint8)
				arg1fun := args[1].WithFun().(func(*Env) uint8)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uint8, uint8)

					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint8, uint8))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint8, uint8))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uint16:
			{
				arg0fun := args[0].WithFun().(func(*Env) uint16)
				arg1fun := args[1].WithFun().(func(*Env) uint16)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uint16, uint16)
					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint16, uint16))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint16, uint16))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uint32:
			{
				arg0fun := args[0].WithFun().(func(*Env) uint32)
				arg1fun := args[1].WithFun().(func(*Env) uint32)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uint32, uint32)
					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint32, uint32))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint32, uint32))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uint64:
			{
				arg0fun := args[0].WithFun().(func(*Env) uint64)
				arg1fun := args[1].WithFun().(func(*Env) uint64)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uint64, uint64)
					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uint64, uint64))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(uint64, uint64))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Uintptr:
			{
				arg0fun := args[0].WithFun().(func(*Env) uintptr)
				arg1fun := args[1].WithFun().(func(*Env) uintptr)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(uintptr, uintptr)
					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(uintptr, uintptr))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(uintptr, uintptr))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Float32:
			{
				arg0fun := args[0].WithFun().(func(*Env) float32)
				arg1fun := args[1].WithFun().(func(*Env) float32)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(float32, float32)
					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float32, float32))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(float32, float32))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Float64:
			{
				arg0fun := args[0].WithFun().(func(*Env) float64)
				arg1fun := args[1].WithFun().(func(*Env) float64)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(float64, float64)
					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(float64, float64))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(float64, float64))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Complex64:
			{
				arg0fun := args[0].WithFun().(func(*Env) complex64)
				arg1fun := args[1].WithFun().(func(*Env) complex64)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(complex64, complex64)
					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex64, complex64))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(complex64, complex64))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.Complex128:
			{
				arg0fun := args[0].WithFun().(func(*Env) complex128)
				arg1fun := args[1].WithFun().(func(*Env) complex128)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(complex128, complex128)
					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(complex128, complex128))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(complex128, complex128))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		case r.String:
			{
				arg0fun := args[0].WithFun().(func(*Env) string)
				arg1fun := args[1].WithFun().(func(*Env) string)

				if funsym != nil && funupn == maxdepth-1 {
					var cachedfun func(string, string)
					call = func(env *Env) {
						funv := env.ThreadGlobals.FileEnv.Binds[funindex]
						if cachedfunv != funv {
							cachedfunv = funv
							cachedfun = funv.Interface().(func(string, string))
						}

						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						cachedfun(arg0, arg1)
					}
				} else {
					call = func(env *Env) {
						fun := exprfun(env).Interface().(func(string, string))
						arg0 := arg0fun(env)
						arg1 := arg1fun(env)
						fun(arg0, arg1)
					}
				}

			}
		}
	}
	if call == nil {
		call = func(env *Env) {
			funv := exprfun(env)

			argv := []r.Value{
				argfunsX1[0](env),
				argfunsX1[1](env),
			}
			funv.Call(argv)
		}
	}
	return call
}
