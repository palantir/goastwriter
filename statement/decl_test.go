// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement_test

import (
	"testing"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/decl"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/spec"
	"github.com/palantir/goastwriter/statement"
)

func TestDecls(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "single const declaration",
			val: &statement.Decl{
				Decl: decl.NewConstSingleValue("foo", expression.StringType, expression.StringVal("FOO")),
			},
			want: `const foo string = "FOO"`,
		},
		{
			name: "single const declaration with inferred type",
			val: &statement.Decl{
				Decl: decl.NewConstSingleValue("foo", expression.Type(""), expression.IntVal(42)),
			},
			want: `const foo = 42`,
		},
		{
			name: "group constant declaration",
			val: &statement.Decl{
				Decl: &decl.Const{
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
			},
			want: `const (
	foo	= "FOO"
	bar	= "BAR"
)`,
		},
		{
			name: "group constant declaration with types",
			val: &statement.Decl{
				Decl: &decl.Const{
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
			},
			want: `const (
	foo	Alias	= "FOO"
	bar	Alias	= "BAR"
)`,
		},
	})
}
