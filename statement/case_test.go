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

func TestCases(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "case clause with single expression and body",
			val: &statement.CaseClause{
				Exprs: []astgen.ASTExpr{
					expression.StringVal("foo"),
				},
				Body: []astgen.ASTStmt{
					&statement.Return{
						Values: []astgen.ASTExpr{
							expression.StringVal("foo"),
						},
					},
				},
			},
			want: `case "foo":
	return "foo"`,
		},
		{
			name: "case clause with multiple expressions and single body",
			val: &statement.CaseClause{
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
			want: `case "foo", "bar":
	return "foo"`,
		},
	})
}
