// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type TypeSwitch struct {
	Init   astgen.ASTStmt
	Assign astgen.ASTStmt
	Cases  []CaseClause
}

func (t *TypeSwitch) ASTStmt() ast.Stmt {
	ss := &ast.TypeSwitchStmt{
		Assign: t.Assign.ASTStmt(),
	}

	if t.Init != nil {
		ss.Init = t.Init.ASTStmt()
	}

	if len(t.Cases) > 0 {
		var cases []ast.Stmt
		for _, v := range t.Cases {
			cases = append(cases, v.ASTStmt())
		}
		ss.Body = &ast.BlockStmt{
			List: cases,
		}
	}

	return ss
}
