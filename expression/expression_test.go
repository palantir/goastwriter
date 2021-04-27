// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"testing"

	"github.com/palantir/goastwriter/astgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name string
	val  astgen.ASTExpr
	want string
}

func runTest(t *testing.T, cases []testCase) {
	for caseNum, currCase := range cases {
		var buf bytes.Buffer
		err := printer.Fprint(&buf, token.NewFileSet(), &ast.ExprStmt{
			X: currCase.val.ASTExpr(),
		})
		src := buf.String()

		require.NoError(t, err, "Case %d: %s", caseNum, currCase.name)
		assert.Equal(t, currCase.want, src, "Case %d: %s\n%s", caseNum, currCase.name, src)
	}
}
