// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type Expression struct {
	Expr astgen.ASTExpr
}

func NewExpression(expr astgen.ASTExpr) *Expression {
	return &Expression{
		Expr: expr,
	}
}

func (e *Expression) ASTStmt() ast.Stmt {
	return &ast.ExprStmt{
		X: e.Expr.ASTExpr(),
	}
}
