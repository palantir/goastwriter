// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decl_test

import (
	"testing"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/decl"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/spec"
)

func TestConsts(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "single const declaration",
			val:  decl.NewConstSingleValue("foo", expression.StringType, expression.StringVal("FOO")),
			want: `const foo string = "FOO"`,
		},
		{
			name: "single const declaration with inferred type",
			val:  decl.NewConstSingleValue("foo", expression.Type(""), expression.IntVal(42)),
			want: `const foo = 42`,
		},
		{
			name: "group constant declaration",
			val: &decl.Const{
				Values: []*spec.Value{
					{
						Names:  []string{"foo"},
						Values: []astgen.ASTExpr{expression.StringVal("FOO")},
					},
					{
						Names:  []string{"bar"},
						Values: []astgen.ASTExpr{expression.StringVal("BAR")},
					},
				},
			},
			want: `const (
	foo	= "FOO"
	bar	= "BAR"
)`,
		},
		{
			name: "group constant declaration with types",
			val: &decl.Const{
				Values: []*spec.Value{
					{
						Names:  []string{"foo"},
						Type:   expression.Type("Alias"),
						Values: []astgen.ASTExpr{expression.StringVal("FOO")},
					},
					{
						Names:  []string{"bar"},
						Type:   expression.Type("Alias"),
						Values: []astgen.ASTExpr{expression.StringVal("BAR")},
					},
				},
			},
			want: `const (
	foo	Alias	= "FOO"
	bar	Alias	= "BAR"
)`,
		},
		{
			name: "group constant declaration with comments",
			val: &decl.Const{
				Values: []*spec.Value{
					{
						Comment: "foo docs",
						Names:   []string{"foo"},
						Type:    expression.Type("Alias"),
						Values:  []astgen.ASTExpr{expression.StringVal("FOO")},
					},
					{
						Comment: "bar multiline docs\nsecond line docs",
						Names:   []string{"bar"},
						Type:    expression.Type("Alias"),
						Values:  []astgen.ASTExpr{expression.StringVal("BAR")},
					},
				},
			},
			want: "const ( \n// foo docs\n\tfoo\tAlias\t= \"FOO\"\n// bar multiline docs\n\t// second line docs\n\tbar\tAlias\t= \"BAR\"\n)",
		},
	})
}
