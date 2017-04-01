// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type If struct {
	Init astgen.ASTStmt
	Cond astgen.ASTExpr
	Body []astgen.ASTStmt
	Else astgen.ASTStmt
}

func (i *If) ASTStmt() ast.Stmt {
	var body []ast.Stmt
	for _, stmter := range i.Body {
		body = append(body, stmter.ASTStmt())
	}

	var initStmt ast.Stmt
	if i.Init != nil {
		initStmt = i.Init.ASTStmt()
	}

	var condExpr ast.Expr
	if i.Cond != nil {
		condExpr = i.Cond.ASTExpr()
	}

	var elseStmt ast.Stmt
	if i.Else != nil {
		elseStmt = i.Else.ASTStmt()
	}

	return &ast.IfStmt{
		Init: initStmt,
		Cond: condExpr,
		Body: &ast.BlockStmt{
			List: body,
		},
		Else: elseStmt,
	}
}
