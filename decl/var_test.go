// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decl_test

import (
	"testing"

	"github.com/palantir/goastwriter/decl"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/spec"
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
		{
			name: "var constructor without value",
			val:  decl.NewVar("sortedKeys", expression.Type("[]string")),
			want: `var sortedKeys []string`,
		},
		{
			name: "single var with implied type",
			val: decl.NewVarWithValues(
				spec.NewValue("key", "", expression.StringVal("value")),
			),
			want: `var key = "value"`,
		},
		{
			name: "single var with explicit type",
			val: decl.NewVarWithValues(
				spec.NewValue("key", expression.StringType, expression.StringVal("value")),
			),
			want: `var key string = "value"`,
		},
		{
			name: "multiple var block with explicit types",
			val: decl.NewVarWithValues(
				spec.NewValue("key1", expression.StringType, expression.StringVal("value1")),
				spec.NewValue("key2", expression.StringType, expression.StringVal("value2")),
			),
			want: `var (
	key1	string	= "value1"
	key2	string	= "value2"
)`,
		},
		{
			name: "multiple var block with implied types",
			val: decl.NewVarWithValues(
				spec.NewValue("key1", "", expression.StringVal("value1")),
				spec.NewValue("key2", "", expression.StringVal("value2")),
			),
			want: `var (
	key1	= "value1"
	key2	= "value2"
)`,
		},
	})
}
