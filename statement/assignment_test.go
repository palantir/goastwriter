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
	"go/token"
	"testing"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/statement"
)

func TestAssignmentStatement(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "variable declaration and assignment",
			val: &statement.Assignment{
				LHS: []astgen.ASTExpr{
					expression.VariableVal("foo"),
				},
				Tok: token.DEFINE,
				RHS: expression.StringVal("FOO"),
			},
			want: `foo := "FOO"`,
		},
		{
			name: "variable assignment only",
			val:  statement.NewAssignment(expression.VariableVal("answer"), token.ASSIGN, expression.IntVal(42)),
			want: `answer = 42`,
		},
		{
			name: "variable declaration and assignment multi-value LHS",
			val: &statement.Assignment{
				LHS: []astgen.ASTExpr{
					expression.VariableVal("foo"),
					expression.VariableVal("err"),
				},
				Tok: token.DEFINE,
				RHS: &expression.CallExpression{
					Function: &expression.Selector{
						Receiver: expression.VariableVal("bar"),
						Selector: "Method",
					},
				},
			},
			want: `foo, err := bar.Method()`,
		},
	})
}
