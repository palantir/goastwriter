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

type FuncType struct {
	Params      FuncParams
	ReturnTypes Types
}

func (f *FuncType) ASTExpr() ast.Expr {
	return &ast.FuncType{
		Params:  f.Params.ToFieldList(),
		Results: f.ReturnTypes.ToFieldList(),
	}
}

type FuncParams []*FuncParam

func (f FuncParams) ToFieldList() *ast.FieldList {
	var fields []*ast.Field
	for _, curr := range f {
		fields = append(fields, curr.ToASTField())
	}
	return &ast.FieldList{
		List: fields,
	}
}

type FuncParam struct {
	Names []string
	Type  Type
}

func NewFuncParam(name string, t Type) *FuncParam {
	return &FuncParam{
		Names: []string{name},
		Type:  t,
	}
}

func (f *FuncParam) ToASTField() *ast.Field {
	var names []*ast.Ident
	for _, name := range f.Names {
		names = append(names, ast.NewIdent(name))
	}
	return &ast.Field{
		Names: names,
		Type:  ast.NewIdent(string(f.Type)),
	}
}
