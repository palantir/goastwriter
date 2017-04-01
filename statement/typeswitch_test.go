// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement_test

import (
	"testing"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/statement"
)

func TestTypeSwitches(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "type switch statement with only default case",
			val: &statement.TypeSwitch{
				Assign: &statement.Expression{
					Expr: &expression.TypeAssert{
						Receiver: expression.VariableVal("foo"),
					},
				},
				Cases: []statement.CaseClause{
					{
						Body: []astgen.ASTStmt{
							&statement.Return{
								Values: []astgen.ASTExpr{
									expression.StringVal("foo"),
								},
							},
						},
					},
				},
			},
			want: `switch foo.(type) {
default:
	return "foo"
}`,
		},
		{
			name: "switch statement with only non-default case",
			val: &statement.TypeSwitch{
				Assign: &statement.Expression{
					Expr: &expression.TypeAssert{
						Receiver: expression.VariableVal("foo"),
					},
				},
				Cases: []statement.CaseClause{
					*statement.NewCaseClause(
						expression.VariableVal(expression.StringType),
						&statement.Return{
							Values: []astgen.ASTExpr{
								expression.StringVal("foo"),
							},
						}),
				},
			},
			want: `switch foo.(type) {
case string:
	return "foo"
}`,
		},
	})
}
