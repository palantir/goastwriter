// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decl_test

import (
	"testing"

	"github.com/palantir/goastwriter/decl"
	"github.com/palantir/goastwriter/expression"
)

func TestVars(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple var declaration",
			val: &decl.Var{
				Name: "sortedKeys",
				Type: expression.Type("[]string"),
			},
			want: `var sortedKeys []string`,
		},
		{
			name: "var declaration with value",
			val: &decl.Var{
				Name:  "key",
				Type:  "string",
				Value: expression.StringVal("value"),
			},
			want: `var key string = "value"`,
		},
		{
			name: "var declaration with value implied type",
			val: &decl.Var{
				Name:  "key",
				Value: expression.StringVal("value"),
			},
			want: `var key = "value"`,
		},
	})
}
