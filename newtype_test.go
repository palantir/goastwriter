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
