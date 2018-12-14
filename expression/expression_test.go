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
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/palantir/goastwriter/astgen"
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
