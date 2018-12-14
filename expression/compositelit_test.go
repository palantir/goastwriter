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
)

func TestCompositeLitExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple composite literal expression",
			val: &expression.CompositeLit{
				Type: expression.Type("http.Client"),
			},
			want: `http.Client{}`,
		},
		{
			name: "composite literal expression with elements",
			val: &expression.CompositeLit{
				Type: expression.Type("http.Client"),
				Elements: []astgen.ASTExpr{
					expression.NewKeyValue("CheckRedirect", expression.Nil),
				},
			},
			want: `http.Client{CheckRedirect: nil}`,
		},
	})
}
