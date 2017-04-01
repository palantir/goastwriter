// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement_test

import (
	"bytes"
	"go/printer"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/palantir/goastwriter/astgen"
)

type testCase struct {
	name string
	val  astgen.ASTStmt
	want string
}

func runTest(t *testing.T, cases []testCase) {
	for caseNum, currCase := range cases {
		var buf bytes.Buffer
		err := printer.Fprint(&buf, token.NewFileSet(), currCase.val.ASTStmt())
		src := buf.String()

		require.NoError(t, err, "Case %d: %s", caseNum, currCase.name)
		assert.Equal(t, currCase.want, src, "Case %d: %s\n%s", caseNum, currCase.name, src)
	}
}
