// Copyright (c) 2018 Palantir Technologies. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
