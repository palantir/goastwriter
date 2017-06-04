// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decl_test

import (
	"testing"

	"github.com/palantir/goastwriter/decl"
	"github.com/palantir/goastwriter/expression"
)

func TestStructs(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple struct",
			val: &decl.Struct{
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
							Type:    expression.Type("bool").Pointer(),
							Comment: "Baz is a field",
						},
					},
				},
			},
			want: `type
// Foo is a struct
Foo struct {
// Bar is a field
	Bar	string
// Baz is a field
	baz	*bool
}`,
		},
		{
			name: "struct with multi-line comment",
			val: &decl.Struct{
				Name:    "Foo",
				Comment: "Foo is a struct\nWith a multi-line comment",
				StructType: expression.StructType{
					Fields: []*expression.StructField{
						{
							Name:    "Bar",
							Type:    "string",
							Comment: "Bar is a field\nWith a multi-line comment",
						},
					},
				},
			},
			want: `type
// Foo is a struct
// With a multi-line comment
Foo struct {
// Bar is a field
	// With a multi-line comment
	Bar string
}`,
		},
		{
			name: "struct with tag",
			val: &decl.Struct{
				Name:    "Bar",
				Comment: "Bar is another struct",
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
							Tag:     `json:"myName"`,
						},
					},
				},
			},
			want: `type
// Bar is another struct
Bar struct {
// Bar is a field
	Bar	string
// Baz is a field
	baz	*bool` + "	`json:\"myName\"`" + `
}`,
		},
	})
}
