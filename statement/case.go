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

package statement

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type CaseClause struct {
	Exprs []astgen.ASTExpr
	Body  []astgen.ASTStmt
}

func NewCaseClause(expr astgen.ASTExpr, body ...astgen.ASTStmt) *CaseClause {
	return &CaseClause{
		Exprs: []astgen.ASTExpr{
			expr,
		},
		Body: body,
	}
}

func (c *CaseClause) ASTStmt() ast.Stmt {
	astCase := &ast.CaseClause{}

	if c.Exprs != nil {
		var exprs []ast.Expr
		for _, v := range c.Exprs {
			exprs = append(exprs, v.ASTExpr())
		}
		astCase.List = exprs
	}

	if c.Body != nil {
		var stmts []ast.Stmt
		for _, v := range c.Body {
			stmts = append(stmts, v.ASTStmt())
		}
		astCase.Body = stmts
	}

	return astCase
}
