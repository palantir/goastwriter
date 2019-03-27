// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goastwriter

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/printer"
	"go/token"

	"github.com/palantir/goastwriter/astgen"
)

func Write(pkgName string, components ...astgen.ASTDecl) ([]byte, error) {
	if pkgName == "" {
		return nil, fmt.Errorf("PackageName must be non-empty")
	}

	var decls []ast.Decl
	for _, currDecler := range components {
		decls = append(decls, currDecler.ASTDecl())
	}

	var buf bytes.Buffer
	if err := printer.Fprint(&buf, token.NewFileSet(), &ast.File{
		Name:  ast.NewIdent(pkgName),
		Decls: decls,
	}); err != nil {
		return nil, fmt.Errorf("failed to convert AST to source code: %v", err)
	}
	return format.Source(buf.Bytes())
}
