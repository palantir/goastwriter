// Copyright 2018 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package statement

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

// Block represents a block of statements wrapped by braces ("{" and "}").
// Useful for the Else block of an If statement.
type Block struct {
	Statments []astgen.ASTStmt
}

func NewBlock(stmts ...astgen.ASTStmt) *Block {
	return &Block{Statments: stmts}
}

func (b *Block) ASTStmt() ast.Stmt {
	block := ast.BlockStmt{List: make([]ast.Stmt, len(b.Statments))}
	for i, stmt := range b.Statments {
		block.List[i] = stmt.ASTStmt()
	}
	return &block
}
