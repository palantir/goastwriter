// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestStructTypeExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "struct type",
			val: &expression.StructType{
				Fields: []*expression.StructField{
					{
						Name:    "Bar",
						Type:    "string",
						Comment: "Bar is a field",
					},
					{
						Name:    "baz",
						Type:    expression.Type("bool").Pointer(),
						Comment: "Baz is a field",
					},
				},
			},
			want: `struct {
// Bar is a field
	Bar	string
// Baz is a field
	baz	*bool
}`,
		},
	})
}
