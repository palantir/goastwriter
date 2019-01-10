// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"
)

type Type string

const (
	BoolType           Type = "bool"
	StringType         Type = "string"
	IntType            Type = "int"
	ErrorType          Type = "error"
	EmptyInterfaceType Type = "interface{}"
	ByteSliceType      Type = "[]byte"
)

// Pointer returns a new type that is a pointer to the current type (prepends a "*").
func (t Type) Pointer() Type {
	return Type("*" + string(t))
}

func (t Type) ToIdent() *ast.Ident {
	return ast.NewIdent(string(t))
}

func (t Type) ASTExpr() ast.Expr {
	return t.ToIdent()
}

type Types []Type

func (t Types) ToFieldList() *ast.FieldList {
	var fields []*ast.Field
	for _, curr := range t {
		fields = append(fields, &ast.Field{
			Type: ast.NewIdent(string(curr)),
		})
	}
	return &ast.FieldList{
		List: fields,
	}
}
