// Copyright 2018 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement_test

import (
	"go/token"
	"testing"

	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/statement"
)

func TestBlockStatement(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple block expression",
			val: statement.NewBlock(
				statement.NewAssignment(expression.VariableVal("foo"), token.ASSIGN, expression.StringVal("foo")),
				statement.NewAssignment(expression.VariableVal("bar"), token.ASSIGN, expression.StringVal("bar")),
			),
			want: `{
	foo = "foo"
	bar = "bar"
}`,
		},
	})
}
