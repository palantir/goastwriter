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

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/statement"
)

func TestFuncLitExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple function literal",
			val: &expression.FuncLit{
				Type: expression.FuncType{},
			},
			want: `func() {
}`,
		},
		{
			name: "function literal with params",
			val: &expression.FuncLit{
				Type: expression.FuncType{
					Params: []*expression.FuncParam{
						expression.NewFuncParam("foo", expression.StringType),
					},
					ReturnTypes: []expression.Type{
						expression.IntType,
					},
				},
			},
			want: `func(foo string) int {
}`,
		},
		{
			name: "function literal with params and body",
			val: &expression.FuncLit{
				Type: expression.FuncType{
					Params: []*expression.FuncParam{
						expression.NewFuncParam("foo", expression.StringType),
					},
					ReturnTypes: []expression.Type{
						expression.IntType,
					},
				},
				Body: []astgen.ASTStmt{
					statement.NewReturn(expression.IntVal(42)),
				},
			},
			want: `func(foo string) int {
	return 42
}`,
		},
	})
}
