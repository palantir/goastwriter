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

package goastwriter_test

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintAST(t *testing.T) {
	src := `package sample

func Foo() {
	json.Marshal(&struct {
			Typ string      ` + "`json:\"type\"`" + `
			Val interface{} ` + "`json:\"string\"`" + `
		}{
			Typ: "string",
			Val: t,
		})
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	require.NoError(t, err)

	out := &bytes.Buffer{}
	err = ast.Fprint(out, fset, f, ast.NotNilFilter)
	require.NoError(t, err)
}
