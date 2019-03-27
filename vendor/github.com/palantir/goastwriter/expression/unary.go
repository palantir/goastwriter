// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"
	"go/token"

	"github.com/palantir/goastwriter/astgen"
)

type Unary struct {
	Op       token.Token
	Receiver astgen.ASTExpr
}

func NewUnary(op token.Token, receiver astgen.ASTExpr) *Unary {
	return &Unary{
		Op:       op,
		Receiver: receiver,
	}
}

func (u *Unary) ASTExpr() ast.Expr {
	return &ast.UnaryExpr{
		Op: u.Op,
		X:  u.Receiver.ASTExpr(),
	}
}
