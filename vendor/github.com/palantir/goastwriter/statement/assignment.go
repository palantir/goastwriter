// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement

import (
	"go/ast"
	"go/token"

	"github.com/palantir/goastwriter/astgen"
)

type Assignment struct {
	LHS []astgen.ASTExpr
	Tok token.Token
	RHS astgen.ASTExpr
}

func NewAssignment(lhs astgen.ASTExpr, tok token.Token, rhs astgen.ASTExpr) *Assignment {
	return &Assignment{
		LHS: []astgen.ASTExpr{
			lhs,
		},
		Tok: tok,
		RHS: rhs,
	}
}

func (a *Assignment) ASTStmt() ast.Stmt {
	var lhs []ast.Expr
	for _, expr := range a.LHS {
		lhs = append(lhs, expr.ASTExpr())
	}
	return &ast.AssignStmt{
		Lhs: lhs,
		Tok: a.Tok,
		Rhs: []ast.Expr{
			a.RHS.ASTExpr(),
		},
	}
}
