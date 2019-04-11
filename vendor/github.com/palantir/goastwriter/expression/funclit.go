// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type FuncLit struct {
	Type FuncType
	Body []astgen.ASTStmt
}

func NewFuncLit(typ FuncType, body ...astgen.ASTStmt) *FuncLit {
	return &FuncLit{
		Type: typ,
		Body: body,
	}
}

func (f *FuncLit) ASTExpr() ast.Expr {
	body := &ast.BlockStmt{}
	if len(f.Body) > 0 {
		var stmts []ast.Stmt
		for _, currStmter := range f.Body {
			stmts = append(stmts, currStmter.ASTStmt())
		}
		body.List = stmts
	}
	return &ast.FuncLit{
		Type: f.Type.ASTExpr().(*ast.FuncType),
		Body: body,
	}
}
