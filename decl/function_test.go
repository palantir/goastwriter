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

package decl_test

import (
	"go/token"
	"testing"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/decl"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/statement"
)

func TestFunctions(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple function",
			val: &decl.Function{
				Name: "GetDataset",
				FuncType: expression.FuncType{
					Params: expression.FuncParams([]*expression.FuncParam{
						expression.NewFuncParam("datasetRID", expression.StringType),
					}),
					ReturnTypes: expression.Types([]expression.Type{
						expression.IntType,
						expression.ErrorType,
					}),
				},
				Comment: "Comment for function",
			},
			want: `func
// Comment for function
GetDataset(datasetRID string) (int, error) {
}`,
		},
		{
			name: "function with body",
			val: &decl.Function{
				Name: "Add",
				FuncType: expression.FuncType{
					Params: expression.FuncParams([]*expression.FuncParam{
						{
							Names: []string{"x", "y"},
							Type:  expression.IntType,
						},
					}),
					ReturnTypes: expression.Types([]expression.Type{
						expression.IntType,
					}),
				},
				Body: []astgen.ASTStmt{
					statement.NewAssignment(
						expression.VariableVal("total"),
						token.DEFINE,
						&expression.Binary{
							LHS: expression.VariableVal("x"),
							Op:  token.ADD,
							RHS: expression.VariableVal("y"),
						},
					),
					&statement.Return{
						Values: []astgen.ASTExpr{
							expression.VariableVal("total"),
						},
					},
				},
				Comment: "Add returns the result of adding the two inputs",
			},
			want: `func
// Add returns the result of adding the two inputs
Add(x, y int) int {
	total := x + y
	return total
}`,
		},
	})
}
