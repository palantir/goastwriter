// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type Return struct {
	Values []astgen.ASTExpr
}

func NewReturn(values ...astgen.ASTExpr) *Return {
	return &Return{
		Values: values,
	}
}

func (r *Return) ASTStmt() ast.Stmt {
	var exprs []ast.Expr
	for _, expr := range r.Values {
		exprs = append(exprs, expr.ASTExpr())
	}
	return &ast.ReturnStmt{
		Results: exprs,
	}
}
