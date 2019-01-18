// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type Decl struct {
	Decl astgen.ASTDecl
}

func NewDecl(decl astgen.ASTDecl) *Decl {
	return &Decl{
		Decl: decl,
	}
}

func (d *Decl) ASTStmt() ast.Stmt {
	return &ast.DeclStmt{
		Decl: d.Decl.ASTDecl(),
	}
}
