// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement

import (
	"go/ast"
	"go/token"

	"github.com/palantir/goastwriter/astgen"
)

type Range struct {
	Key, Value astgen.ASTExpr
	Tok        token.Token
	Expr       astgen.ASTExpr
	Body       []astgen.ASTStmt
}

func (r *Range) ASTStmt() ast.Stmt {
	var value ast.Expr
	if r.Value != nil {
		value = r.Value.ASTExpr()
	}

	var body []ast.Stmt
	for _, stmter := range r.Body {
		body = append(body, stmter.ASTStmt())
	}

	return &ast.RangeStmt{
		Key:   r.Key.ASTExpr(),
		Value: value,
		Tok:   r.Tok,
		X:     r.Expr.ASTExpr(),
		Body: &ast.BlockStmt{
			List: body,
		},
	}
}
