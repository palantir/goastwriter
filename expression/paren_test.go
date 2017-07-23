// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestParenExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "paren expression",
			val: &expression.Paren{
				Expression: expression.VariableVal("foo"),
			},
			want: `(foo)`,
		},
	})
}
