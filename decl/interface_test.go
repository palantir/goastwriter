// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decl_test

import (
	"testing"

	"github.com/palantir/goastwriter/decl"
	"github.com/palantir/goastwriter/expression"
)

func TestInterfaces(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple interface",
			val: &decl.Interface{
				Name: "Writer",
				InterfaceType: expression.InterfaceType{
					Functions: expression.InterfaceFunctionDecls([]*expression.InterfaceFunctionDecl{
						{
							Name: "GetDataset",
							Params: expression.FuncParams([]*expression.FuncParam{
								expression.NewFuncParam("datasetRID", expression.StringType),
							}),
							ReturnTypes: expression.Types([]expression.Type{
								expression.IntType,
							}),
						},
					}),
				},
				Comment: "Comment for interface",
			},
			want: `type
// Comment for interface
Writer interface {
	GetDataset(datasetRID string) int
}`,
		},
	})
}
