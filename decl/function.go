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

package decl

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/comment"
	"github.com/palantir/goastwriter/expression"
)

type Function struct {
	Name     string
	FuncType expression.FuncType
	Body     []astgen.ASTStmt
	Comment  string
}

func (f *Function) ASTDecl() ast.Decl {
	return f.funcDecl(nil)
}

func (f *Function) funcDecl(receiver *ast.Field) *ast.FuncDecl {
	var fieldList *ast.FieldList
	if receiver != nil {
		fieldList = &ast.FieldList{
			List: []*ast.Field{
				receiver,
			},
		}
	}
	body := &ast.BlockStmt{}
	if len(f.Body) > 0 {
		var stmts []ast.Stmt
		for _, currStmter := range f.Body {
			stmts = append(stmts, currStmter.ASTStmt())
		}
		body.List = stmts
	}
	return &ast.FuncDecl{
		Doc:  comment.ToComment(f.Comment),
		Recv: fieldList,
		Name: ast.NewIdent(f.Name),
		Type: f.FuncType.ASTExpr().(*ast.FuncType),
		Body: body,
	}
}
