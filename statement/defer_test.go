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

func TestDeferStatement(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple defer statement",
			val: &statement.Defer{
				Call: expression.NewCallFunction("fmt", "Println", expression.StringVal("foo")),
			},
			want: `defer fmt.Println("foo")`,
		},
		{
			name: "defer statement with function literal",
			val: &statement.Defer{
				Call: &expression.CallExpression{
					Function: &expression.FuncLit{
						Body: []astgen.ASTStmt{
							statement.NewExpression(expression.NewCallFunction("fmt", "Println", expression.StringVal("foo"))),
						},
					},
				},
			},
			want: `defer func() {
	fmt.Println("foo")
}()`,
		},
	})
}
