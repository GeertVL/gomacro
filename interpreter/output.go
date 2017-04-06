/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * output.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"fmt"
	"go/ast"
	"io"
	r "reflect"
	"sort"

	. "github.com/cosmos72/gomacro/base"
)

var (
	nilEnv *Env
	NilEnv = []r.Value{r.ValueOf(nilEnv)}
)

func (ir *InterpreterCommon) showHelp(out io.Writer) {
	fmt.Fprint(out, `// interpreter commands:
:env [name]     show available functions, variables and constants
                in current package, or from imported package "name"
:help           print this help
:inspect EXPR   inspect expression interactively
:options [OPTS] show or toggle interpreter options
:quit           quit the interpreter
:write [FILE]   write collected declarations and/or statements to standard output or to file
                use :o Declarations and/or :o Statements to start collecting them
`)
}

func (env *Env) showStack() {
	frames := env.CallStack.Frames
	n := len(frames)
	for i := 1; i < n; i++ {
		frame := &frames[i]
		name := ""
		if frame.FuncEnv != nil {
			name = frame.FuncEnv.Name
		}
		if frame.panicking {
			env.Debugf("%d:\t     %v, runningDefers = %v, panic = %v", i, name, frame.runningDefers, frame.panick)
		} else {
			env.Debugf("%d:\t     %v, runningDefers = %v, panic is nil", i, name, frame.runningDefers)
		}
	}
}

func (env *Env) showPackage(packageName string) {
	out := env.Stdout
	e := env
	path := env.Path
	pkg := &env.Package
	if len(packageName) != 0 {
		bind := env.evalIdentifier(&ast.Ident{Name: packageName})
		if bind == None || bind == Nil {
			env.Warnf("not an imported package: %q", packageName)
			return
		}
		switch val := bind.Interface().(type) {
		case *PackageRef:
			e = nil
			pkg = &val.Package
			path = packageName
		default:
			env.Warnf("not an imported package: %q = %v <%v>", packageName, val, typeOf(bind))
			return
		}
	}
	spaces15 := "               "
Loop:
	binds := pkg.Binds
	if len(binds) > 0 {
		fmt.Fprintf(out, "// ----- %s binds -----\n", path)

		keys := make([]string, len(binds))
		i := 0
		for k := range binds {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			n := len(k) & 15
			fmt.Fprintf(out, "%s%s = ", k, spaces15[n:])
			bind := binds[k]
			if bind != Nil {
				switch bind := bind.Interface().(type) {
				case *Env:
					fmt.Fprintf(out, "%p <%v>\n", bind, r.TypeOf(bind))
					continue
				}
			}
			env.FprintValue(out, bind)
		}
		fmt.Fprintln(out)
	}
	types := pkg.Types
	if len(types) > 0 {
		fmt.Fprintf(out, "// ----- %s types -----\n", path)

		keys := make([]string, len(types))
		i := 0
		for k := range types {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			n := len(k) & 15
			t := types[k]
			fmt.Fprintf(out, "%s%s %v <%v>\n", k, spaces15[n:], t.Kind(), t)
		}
		fmt.Fprintln(out)
	}
	if e != nil {
		if e = e.Outer; e != nil {
			path = e.Path
			pkg = &e.Package
			goto Loop
		}
	}
}
