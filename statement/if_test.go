// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement_test

import (
	"go/token"
	"testing"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/decl"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/statement"
)

func TestIfStatement(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple if statement",
			val: &statement.If{
				Cond: &expression.Binary{
					LHS: expression.VariableVal("err"),
					Op:  token.NEQ,
					RHS: expression.VariableVal("nil"),
				},
				Body: []astgen.ASTStmt{
					statement.NewDecl(decl.NewVar("foo", expression.IntType)),
				},
			},
			want: `if err != nil {
	var foo int
}`,
		},
		{
			name: "if statement with initialization",
			val: &statement.If{
				Init: statement.NewAssignment(expression.VariableVal("err"), token.DEFINE, expression.NewCallFunction("fmt", "Errorf")),
				Cond: &expression.Binary{
					LHS: expression.VariableVal("err"),
					Op:  token.NEQ,
					RHS: expression.VariableVal("nil"),
				},
				Body: []astgen.ASTStmt{
					statement.NewDecl(decl.NewVar("foo", expression.IntType)),
				},
			},
			want: `if err := fmt.Errorf(); err != nil {
	var foo int
}`,
		},
		{
			name: "if statement with else",
			val: &statement.If{
				Cond: &expression.Binary{
					LHS: expression.VariableVal("err"),
					Op:  token.NEQ,
					RHS: expression.VariableVal("nil"),
				},
				Body: []astgen.ASTStmt{
					statement.NewDecl(decl.NewVar("foo", expression.IntType)),
				},
				Else: &statement.Return{
					Values: []astgen.ASTExpr{
						expression.VariableVal("nil"),
					},
				},
			},
			want: `if err != nil {
	var foo int
} else {
	return nil
}`,
		},
	})
}
