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

func TestSwitches(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "switch statement with only default case",
			val: &statement.Switch{
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
			want: `switch {
default:
	return "foo"
}`,
		},
		{
			name: "switch statement with only non-default case",
			val: &statement.Switch{
				Cases: []statement.CaseClause{
					*statement.NewCaseClause(
						expression.StringVal("foo"),
						&statement.Return{
							Values: []astgen.ASTExpr{
								expression.StringVal("foo"),
							},
						}),
				},
			},
			want: `switch {
case "foo":
	return "foo"
}`,
		},
		{
			name: "switch statement with multiple cases including default",
			val: &statement.Switch{
				Cases: []statement.CaseClause{
					{
						Exprs: []astgen.ASTExpr{
							expression.StringVal("foo"),
							expression.StringVal("bar"),
						},
						Body: []astgen.ASTStmt{
							&statement.Return{
								Values: []astgen.ASTExpr{
									expression.StringVal("foo"),
								},
							},
						},
					},
					*statement.NewCaseClause(
						expression.StringVal("baz"),
						&statement.Return{
							Values: []astgen.ASTExpr{
								expression.StringVal("baz"),
							},
						},
					),
					{
						Body: []astgen.ASTStmt{
							&statement.Return{
								Values: []astgen.ASTExpr{
									expression.StringVal("default"),
								},
							},
						},
					},
				},
			},
			want: `switch {
case "foo", "bar":
	return "foo"
case "baz":
	return "baz"
default:
	return "default"
}`,
		},
	})
}
