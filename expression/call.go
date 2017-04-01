// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type CallExpression struct {
	Function astgen.ASTExpr
	Args     []astgen.ASTExpr
}

func NewCallFunction(receiver, function string, args ...astgen.ASTExpr) *CallExpression {
	return &CallExpression{
		Function: &Selector{
			Receiver: VariableVal(receiver),
			Selector: function,
		},
		Args: args,
	}
}

func (c *CallExpression) ASTExpr() ast.Expr {
	var args []ast.Expr
	for _, curr := range c.Args {
		args = append(args, curr.ASTExpr())
	}
	return &ast.CallExpr{
		Fun:  c.Function.ASTExpr(),
		Args: args,
	}
}
