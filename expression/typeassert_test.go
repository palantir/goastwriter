// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestTypeAssertExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "type assertion",
			val: &expression.TypeAssert{
				Receiver: expression.VariableVal("fooVar"),
			},
			want: `fooVar.(type)`,
		},
		{
			name: "type assertion with type specified",
			val: &expression.TypeAssert{
				Receiver: expression.VariableVal("fooVar"),
				Type:     expression.VariableVal(expression.BoolType),
			},
			want: `fooVar.(bool)`,
		},
	})
}
