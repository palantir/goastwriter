// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement_test

import (
	"go/token"
	"testing"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/statement"
)

func TestAssignmentStatement(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "variable declaration and assignment",
			val: &statement.Assignment{
				LHS: []astgen.ASTExpr{
					expression.VariableVal("foo"),
				},
				Tok: token.DEFINE,
				RHS: expression.StringVal("FOO"),
			},
			want: `foo := "FOO"`,
		},
		{
			name: "variable assignment only",
			val:  statement.NewAssignment(expression.VariableVal("answer"), token.ASSIGN, expression.IntVal(42)),
			want: `answer = 42`,
		},
		{
			name: "variable declaration and assignment multi-value LHS",
			val: &statement.Assignment{
				LHS: []astgen.ASTExpr{
					expression.VariableVal("foo"),
					expression.VariableVal("err"),
				},
				Tok: token.DEFINE,
				RHS: &expression.CallExpression{
					Function: &expression.Selector{
						Receiver: expression.VariableVal("bar"),
						Selector: "Method",
					},
				},
			},
			want: `foo, err := bar.Method()`,
		},
	})
}
