// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestIndexExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple slice index expression",
			val: &expression.Index{
				Receiver: expression.VariableVal("foo"),
				Index:    expression.VariableVal("0"),
			},
			want: `foo[0]`,
		},
		{
			name: "simple map index expression",
			val: &expression.Index{
				Receiver: expression.VariableVal("foo"),
				Index:    expression.VariableVal(`"key"`),
			},
			want: `foo["key"]`,
		},
	})
}
