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

func TestRangeStatement(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple range statement",
			val: &statement.Range{
				Key:  expression.VariableVal("k"),
				Tok:  token.DEFINE,
				Expr: expression.VariableVal("vals"),
			},
			want: `for k := range vals {
}`,
		},
		{
			name: "range statement with body",
			val: &statement.Range{
				Key:   expression.VariableVal("k"),
				Value: expression.VariableVal("v"),
				Tok:   token.DEFINE,
				Expr:  expression.VariableVal("vals"),
				Body: []astgen.ASTStmt{
					statement.NewDecl(decl.NewVar("foo", expression.IntType)),
				},
			},
			want: `for k, v := range vals {
	var foo int
}`,
		},
	})
}
