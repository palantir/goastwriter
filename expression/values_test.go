// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestValueExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "string value expression",
			val:  expression.StringVal("FOO"),
			want: `"FOO"`,
		},
		{
			name: "int value expression",
			val:  expression.IntVal(42),
			want: `42`,
		},
		{
			name: "variable value expression",
			val:  expression.VariableVal("myVar"),
			want: `myVar`,
		},
	})
}
