// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type Paren struct {
	Expression astgen.ASTExpr
}

func NewParen(expr astgen.ASTExpr) *Paren {
	return &Paren{
		Expression: expr,
	}
}

func (p *Paren) ASTExpr() ast.Expr {
	return &ast.ParenExpr{
		X: p.Expression.ASTExpr(),
	}
}
