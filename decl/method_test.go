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
	"testing"

	"github.com/palantir/goastwriter/decl"
	"github.com/palantir/goastwriter/expression"
)

func TestMethods(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple method",
			val: &decl.Method{
				Function: decl.Function{
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
					Comment: "Comment for method",
				},
				ReceiverName: "c",
				ReceiverType: expression.Type("*serviceClient"),
			},
			want: `func
// Comment for method
(c *serviceClient) GetDataset(datasetRID string) (int, error) {
}`,
		},
	})
}
