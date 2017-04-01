// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"go/token"
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestUnaryExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple unary expression",
			val: &expression.Unary{
				Op:       token.AND,
				Receiver: expression.VariableVal("foo"),
			},
			want: `&foo`,
		},
	})
}
