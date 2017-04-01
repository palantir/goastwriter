// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"go/token"
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestBinaryExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple binary expression",
			val: &expression.Binary{
				LHS: expression.IntVal(42),
				Op:  token.ADD,
				RHS: expression.NewCallFunction("consts", "SomeInt"),
			},
			want: `42 + consts.SomeInt()`,
		},
	})
}
