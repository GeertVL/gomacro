// +build !go1.8

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
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 *
 * plugin_dummy.go
 *
 *  Created on Mar 21, 2017
 *      Author Massimiliano Ghilardi
 */

package base

import (
	"io"
	"os"
)

func getGoPath() string {
	return os.Getenv("GOPATH")
}

func getGoSrcPath() string {
	return getGoPath() + "/src"
}

func (ir *InterpreterCommon) compilePlugin(filename string, stdout io.Writer, stderr io.Writer) string {
	errorf("gomacro compiled on Go version < 1.8. No support to load plugins - cannot import packages at runtime")
	return ""
}

func loadPlugin(soname string, symbolName string) interface{} {
	errorf("gomacro compiled on Go version < 1.8. No support to load plugins - cannot import packages at runtime")
	return nil
}
