// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type Switch struct {
	Init       astgen.ASTStmt
	Expression astgen.ASTExpr
	Cases      []CaseClause
}

func (s *Switch) ASTStmt() ast.Stmt {
	ss := &ast.SwitchStmt{}

	if s.Init != nil {
		ss.Init = s.Init.ASTStmt()
	}

	if s.Expression != nil {
		ss.Tag = s.Expression.ASTExpr()
	}

	if len(s.Cases) > 0 {
		var cases []ast.Stmt
		for _, v := range s.Cases {
			cases = append(cases, v.ASTStmt())
		}
		ss.Body = &ast.BlockStmt{
			List: cases,
		}
	}

	return ss
}
