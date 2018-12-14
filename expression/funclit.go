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

	"github.com/palantir/goastwriter/astgen"
)

type FuncLit struct {
	Type FuncType
	Body []astgen.ASTStmt
}

func NewFuncLit(typ FuncType, body ...astgen.ASTStmt) *FuncLit {
	return &FuncLit{
		Type: typ,
		Body: body,
	}
}

func (f *FuncLit) ASTExpr() ast.Expr {
	body := &ast.BlockStmt{}
	if len(f.Body) > 0 {
		var stmts []ast.Stmt
		for _, currStmter := range f.Body {
			stmts = append(stmts, currStmter.ASTStmt())
		}
		body.List = stmts
	}
	return &ast.FuncLit{
		Type: f.Type.ASTExpr().(*ast.FuncType),
		Body: body,
	}
}
