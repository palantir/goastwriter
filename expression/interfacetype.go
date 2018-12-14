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

	"github.com/palantir/goastwriter/comment"
)

type InterfaceType struct {
	Functions InterfaceFunctionDecls
}

func NewInterfaceType(functions ...*InterfaceFunctionDecl) *InterfaceType {
	return &InterfaceType{
		Functions: functions,
	}
}

func (i *InterfaceType) ASTExpr() ast.Expr {
	return &ast.InterfaceType{
		Methods: i.Functions.ToFieldList(),
	}
}

type InterfaceFunctionDecls []*InterfaceFunctionDecl

func (i InterfaceFunctionDecls) ToFieldList() *ast.FieldList {
	var fields []*ast.Field
	for _, curr := range i {
		fields = append(fields, curr.ToASTField())
	}
	return &ast.FieldList{
		List: fields,
	}
}

type InterfaceFunctionDecl struct {
	Name        string
	Params      FuncParams
	ReturnTypes Types
	Comment     string
}

func (i *InterfaceFunctionDecl) ToASTField() *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(i.Name),
		},
		Type: &ast.FuncType{
			Params:  i.Params.ToFieldList(),
			Results: i.ReturnTypes.ToFieldList(),
		},
		Doc: comment.ToComment(i.Comment),
	}
}
