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
