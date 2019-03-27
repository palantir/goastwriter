// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type CompositeLit struct {
	Type     astgen.ASTExpr
	Elements []astgen.ASTExpr
}

func NewCompositeLit(typ astgen.ASTExpr, elems ...astgen.ASTExpr) *CompositeLit {
	return &CompositeLit{
		Type:     typ,
		Elements: elems,
	}
}

func (c *CompositeLit) ASTExpr() ast.Expr {
	var elts []ast.Expr
	for _, el := range c.Elements {
		elts = append(elts, el.ASTExpr())
	}
	return &ast.CompositeLit{
		Type: c.Type.ASTExpr(),
		Elts: elts,
	}
}
