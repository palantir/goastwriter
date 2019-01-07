// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/expression"
)

func TestCallExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "string value expression",
			val: &expression.CallExpression{
				Function: &expression.Selector{
					Receiver: expression.VariableVal("strings"),
					Selector: "HasPrefix",
				},
				Args: []astgen.ASTExpr{
					expression.VariableVal("s"),
					expression.StringVal("prefix-"),
				},
			},
			want: `strings.HasPrefix(s, "prefix-")`,
		},
		{
			name: "string value expression with convenience constructor",
			val: expression.NewCallFunction("strings", "HasPrefix",
				expression.VariableVal("s"),
				expression.StringVal("prefix-")),
			want: `strings.HasPrefix(s, "prefix-")`,
		},
	})
}
