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
				Type: expression.FunctionType{},
			},
			want: `func() {
}`,
		},
		{
			name: "function literal with params",
			val: &expression.FuncLit{
				Type: expression.FunctionType{
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
				Type: expression.FunctionType{
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
