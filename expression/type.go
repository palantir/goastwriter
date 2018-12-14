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
