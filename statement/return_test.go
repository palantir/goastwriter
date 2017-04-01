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

func TestReturnStatement(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple return statement",
			val: &statement.Return{
				Values: []astgen.ASTExpr{
					expression.IntVal(42),
				},
			},
			want: `return 42`,
		},
		{
			name: "multi-value return statement",
			val: &statement.Return{
				Values: []astgen.ASTExpr{
					expression.IntVal(42),
					expression.VariableVal("nil"),
				},
			},
			want: `return 42, nil`,
		},
	})
}
