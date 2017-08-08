// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement

import (
	"go/ast"

	"github.com/palantir/goastwriter/expression"
)

type Defer struct {
	Call *expression.CallExpression
}

func (d *Defer) ASTStmt() ast.Stmt {
	return &ast.DeferStmt{
		Call: d.Call.ASTExpr().(*ast.CallExpr),
	}
}
