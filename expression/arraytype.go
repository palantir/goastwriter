// Copyright 2020 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type ArrayType struct {
	Len astgen.ASTExpr
	Elt Type
}

func (a *ArrayType) ASTExpr() ast.Expr {
	var lenExpr ast.Expr
	if a.Len != nil {
		lenExpr = a.Len.ASTExpr()
	}
	return &ast.ArrayType{
		Len: lenExpr,
		Elt: a.Elt.ASTExpr(),
	}
}
