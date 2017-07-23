// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goastwriter_test

import (
	"fmt"
	"testing"

	"github.com/palantir/goastwriter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/decl"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/statement"
)

func TestPackage_Source(t *testing.T) {
	for caseNum, currCase := range []struct {
		name  string
		pkg   string
		parts []astgen.ASTDecl
		want  string
	}{
		{
			name: "package with multiple different types",
			pkg:  "testpkg",
			parts: []astgen.ASTDecl{
				decl.NewImports(map[string]string{
					"fmt":       "",
					"go/format": "gofmt",
				}),
				&decl.Struct{
					Name:    "Foo",
					Comment: "Foo is a struct",
					StructType: expression.StructType{
						Fields: []*expression.StructField{
							{
								Name:    "Bar",
								Type:    "string",
								Comment: "Bar is a field",
							},
							{
								Name:    "baz",
								Type:    "*bool",
								Comment: "Baz is a field",
							},
						},
					},
				},
				&decl.Function{
					Name: "Bar",
					FuncType: expression.FuncType{
						Params: []*expression.FuncParam{
							expression.NewFuncParam("input", expression.Type("Foo").Pointer()),
						},
						ReturnTypes: []expression.Type{
							expression.Type("Foo").Pointer(),
							expression.ErrorType,
						},
					},
					Body: []astgen.ASTStmt{
						&statement.Expression{
							Expr: expression.NewCallFunction("fmt", "Println"),
						},
						&statement.Expression{
							Expr: expression.NewCallFunction("gofmt", "Source", expression.Nil),
						},
						&statement.Return{
							Values: []astgen.ASTExpr{
								expression.VariableVal("input"),
								expression.Nil,
							},
						},
					},
				},
			},
			want: `package testpkg

import (
	"fmt"
	gofmt "go/format"
)

// Foo is a struct
type Foo struct {
	// Bar is a field
	Bar string
	// Baz is a field
	baz *bool
}

func Bar(input *Foo) (*Foo, error) {
	fmt.Println()
	gofmt.Source(nil)
	return input, nil
}
`,
		},
	} {
		got, err := goastwriter.Write(currCase.pkg, currCase.parts...)
		require.NoError(t, err, "Case %d: %s", caseNum, currCase.name)
		assert.Equal(t, currCase.want, string(got), "Case %d: %s", caseNum, currCase.name)
	}
}

func ExampleWrite() {
	out, _ := goastwriter.Write("testpkg",
		decl.NewImports(map[string]string{
			"fmt":       "",
			"go/format": "gofmt",
		}),
		&decl.Struct{
			Name:    "Foo",
			Comment: "Foo is a struct",
			StructType: expression.StructType{
				Fields: []*expression.StructField{
					{
						Name:    "Bar",
						Type:    expression.StringType,
						Comment: "Bar is a field",
					},
					{
						Name:    "baz",
						Type:    expression.BoolType.Pointer(),
						Comment: "Baz is a field",
					},
				},
			},
		},
		&decl.Function{
			Name: "Bar",
			FuncType: expression.FuncType{
				Params: []*expression.FuncParam{
					expression.NewFuncParam("input", expression.Type("Foo").Pointer()),
				},
				ReturnTypes: []expression.Type{
					expression.Type("Foo").Pointer(),
					expression.ErrorType,
				},
			},
			Body: []astgen.ASTStmt{
				&statement.Expression{
					Expr: expression.NewCallFunction("fmt", "Println"),
				},
				&statement.Expression{
					Expr: expression.NewCallFunction("gofmt", "Source", expression.Nil),
				},
				&statement.Return{
					Values: []astgen.ASTExpr{
						expression.VariableVal("input"),
						expression.Nil,
					},
				},
			},
		},
	)
	fmt.Println(string(out))
	// Output:
	// package testpkg
	//
	// import (
	// 	"fmt"
	// 	gofmt "go/format"
	// )
	//
	// // Foo is a struct
	// type Foo struct {
	// 	// Bar is a field
	// 	Bar string
	// 	// Baz is a field
	// 	baz *bool
	// }
	//
	// func Bar(input *Foo) (*Foo, error) {
	// 	fmt.Println()
	// 	gofmt.Source(nil)
	// 	return input, nil
	// }
}
