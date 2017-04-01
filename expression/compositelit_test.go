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
