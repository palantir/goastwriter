// Copyright 2018 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

// Index represents accessing an array or map by index, e.g. `myMap["myKey"]` or `myArr[0]`
type Index struct {
	Receiver astgen.ASTExpr
	Index    astgen.ASTExpr
}

func NewIndex(receiver astgen.ASTExpr, index astgen.ASTExpr) *Index {
	return &Index{
		Receiver: receiver,
		Index:    index,
	}
}

func (i *Index) ASTExpr() ast.Expr {
	var expr ast.IndexExpr
	if i.Receiver != nil {
		expr.X = i.Receiver.ASTExpr()
	}
	if i.Index != nil {
		expr.Index = i.Index.ASTExpr()
	}
	return &expr
}
