// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type TypeAssert struct {
	Receiver astgen.ASTExpr
	Type     astgen.ASTExpr
}

func NewTypeAssert(receiver astgen.ASTExpr, typ astgen.ASTExpr) *TypeAssert {
	return &TypeAssert{
		Receiver: receiver,
		Type:     typ,
	}
}

func (t *TypeAssert) ASTExpr() ast.Expr {
	expr := &ast.TypeAssertExpr{
		X: t.Receiver.ASTExpr(),
	}
	if t.Type != nil {
		expr.Type = t.Type.ASTExpr()
	}
	return expr
}
