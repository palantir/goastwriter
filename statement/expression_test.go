// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement_test

import (
	"testing"

	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/statement"
)

func TestExpressionStatement(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple expression statement",
			val: &statement.Expression{
				Expr: expression.VariableVal("foo"),
			},
			want: `foo`,
		},
	})
}
