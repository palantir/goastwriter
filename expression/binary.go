// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"
	"go/token"

	"github.com/palantir/goastwriter/astgen"
)

type Binary struct {
	LHS astgen.ASTExpr
	Op  token.Token
	RHS astgen.ASTExpr
}

func NewBinary(lhs astgen.ASTExpr, op token.Token, rhs astgen.ASTExpr) *Binary {
	return &Binary{
		LHS: lhs,
		Op:  op,
		RHS: rhs,
	}
}

func (b *Binary) ASTExpr() ast.Expr {
	return &ast.BinaryExpr{
		X:  b.LHS.ASTExpr(),
		Op: b.Op,
		Y:  b.RHS.ASTExpr(),
	}
}
