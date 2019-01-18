// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type Star struct {
	Expression astgen.ASTExpr
}

func NewStar(expr astgen.ASTExpr) *Star {
	return &Star{
		Expression: expr,
	}
}

func (s *Star) ASTExpr() ast.Expr {
	return &ast.StarExpr{
		X: s.Expression.ASTExpr(),
	}
}
