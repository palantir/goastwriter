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
	"go/token"

	"github.com/palantir/goastwriter/astgen"
)

type Range struct {
	Key, Value astgen.ASTExpr
	Tok        token.Token
	Expr       astgen.ASTExpr
	Body       []astgen.ASTStmt
}

func (r *Range) ASTStmt() ast.Stmt {
	var value ast.Expr
	if r.Value != nil {
		value = r.Value.ASTExpr()
	}

	var body []ast.Stmt
	for _, stmter := range r.Body {
		body = append(body, stmter.ASTStmt())
	}

	return &ast.RangeStmt{
		Key:   r.Key.ASTExpr(),
		Value: value,
		Tok:   r.Tok,
		X:     r.Expr.ASTExpr(),
		Body: &ast.BlockStmt{
			List: body,
		},
	}
}
