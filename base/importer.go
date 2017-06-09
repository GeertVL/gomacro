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
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * importer.go
 *
 *  Created on Feb 27, 2017
 *      Author Massimiliano Ghilardi
 */

package base

import (
	"bytes"
	"fmt"
	"go/importer"
	"go/types"
	"io/ioutil"
	"os"
	r "reflect"
	"strings"

	"github.com/cosmos72/gomacro/imports"
)

type ImportMode int

const (
	ImSharedLib ImportMode = iota
	ImBuiltin
	ImInception
)

type Importer struct {
	from   types.ImporterFrom
	compat types.Importer
	srcDir string
	mode   types.ImportMode
}

func DefaultImporter() *Importer {
	imp := &Importer{}
	compat := importer.Default()
	if from, ok := compat.(types.ImporterFrom); ok {
		imp.from = from
	} else {
		imp.compat = compat
	}
	return imp
}

func (imp *Importer) Import(path string) (*types.Package, error) {
	if imp.from != nil {
		return imp.from.ImportFrom(path, imp.srcDir, imp.mode)
	} else {
		return imp.compat.Import(path)
	}
}

func (imp *Importer) ImportFrom(path string, srcDir string, mode types.ImportMode) (*types.Package, error) {
	if imp.from != nil {
		return imp.from.ImportFrom(path, srcDir, mode)
	} else {
		return imp.compat.Import(path)
	}
}

// LookupPackage returns a package if already present in cache
func (g *Globals) LookupPackage(name, path string) *PackageRef {
	pkg, found := imports.Packages[path]
	if !found {
		return nil
	}
	return &PackageRef{Package: pkg, Name: name, Path: path}
}

func (g *Globals) ImportPackage(name, path string) *PackageRef {
	ref := g.LookupPackage(name, path)
	if ref != nil {
		return ref
	}
	gpkg, err := g.Importer.Import(path) // loads names and types, not the values!
	if err != nil {
		g.Errorf("error loading package %q metadata, maybe you need to download (go get), compile (go build) and install (go install) it? %v", path, err)
	}
	var mode ImportMode
	switch name {
	case "_b":
		mode = ImBuiltin
	case "_i":
		mode = ImInception
	}
	file := g.createImportFile(path, gpkg, mode)
	if mode != ImSharedLib {
		return nil
	}
	ref = &PackageRef{Name: name, Path: path}
	if len(file) == 0 {
		// empty package. still cache it for future use.
		imports.Packages[path] = ref.Package
		return ref
	}
	soname := g.compilePlugin(file, g.Stdout, g.Stderr)
	ifun := g.loadPlugin(soname, "Exports")
	fun := ifun.(func() (map[string]r.Value, map[string]r.Type, map[string]r.Type, map[string]string, map[string][]string))
	binds, types, proxies, untypeds, wrappers := fun()

	// done. cache package for future use.
	ref.Package = imports.Package{
		Binds:    binds,
		Types:    types,
		Proxies:  proxies,
		Untypeds: untypeds,
		Wrappers: wrappers,
	}
	imports.Packages[path] = ref.Package
	return ref
}

func (g *Globals) createImportFile(path string, pkg *types.Package, mode ImportMode) string {
	buf := bytes.Buffer{}
	isEmpty := g.writeImportFile(&buf, path, pkg, mode)
	if isEmpty {
		g.Warnf("package %q exports zero constants, functions, types and variables", path)
		return ""
	}

	file := computeImportFilename(path, mode)
	err := ioutil.WriteFile(file, buf.Bytes(), os.FileMode(0666))
	if err != nil {
		g.Errorf("error writing file %q: %v", file, err)
	}
	if mode != ImSharedLib {
		g.Warnf("created file %q, recompile gomacro to use it", file)
	} else {
		g.Debugf("created file %q...", file)
	}
	return file
}

func sanitizeIdentifier(str string) string {
	return sanitizeIdentifier2(str, '_')
}

func sanitizeIdentifier2(str string, replacement rune) string {
	runes := []rune(str)
	for i, ch := range runes {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || ch == '_' ||
			(i != 0 && ch >= '0' && ch <= '9') {
			continue
		}
		runes[i] = replacement
	}
	return string(runes)
}

const gomacro_dir = "github.com/cosmos72/gomacro"

func computeImportFilename(path string, mode ImportMode) string {
	srcdir := getGoSrcPath()

	switch mode {
	case ImBuiltin:
		return fmt.Sprintf("%s/%s/imports/%s.go", srcdir, gomacro_dir, sanitizeIdentifier(path))
	case ImInception:
		return fmt.Sprintf("%s/%s/x_package.go", srcdir, path)
	}

	file := path[1+strings.LastIndexByte(path, '/'):]
	file = fmt.Sprintf("%s/gomacro_imports/%s/%s.go", srcdir, path, file)
	dir := file[0 : 1+strings.LastIndexByte(file, '/')]
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		Errorf("error creating directory %q: %v", dir, err)
	}
	return file
}
