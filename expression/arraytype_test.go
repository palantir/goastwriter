// Copyright 2020 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestArrayTypeExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "slice type",
			val: &expression.ArrayType{
				Elt: expression.EmptyInterfaceType,
			},
			want: `[]interface{}`,
		},
		{
			name: "array type",
			val: &expression.ArrayType{
				Len: expression.IntVal(2),
				Elt: expression.BoolType,
			},
			want: `[2]bool`,
		},
		{
			name: "ellipsis array type",
			val: &expression.ArrayType{
				Len: expression.Ellipsis{},
				Elt: expression.BoolType,
			},
			want: `[...]bool`,
		},
	})
}
