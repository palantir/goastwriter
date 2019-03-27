// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type Selector struct {
	Receiver astgen.ASTExpr
	Selector string
}

func NewSelector(receiver astgen.ASTExpr, selector string) *Selector {
	return &Selector{
		Receiver: receiver,
		Selector: selector,
	}
}

func (s *Selector) ASTExpr() ast.Expr {
	var receiver ast.Expr
	if s.Receiver != nil {
		receiver = s.Receiver.ASTExpr()
	}
	return &ast.SelectorExpr{
		X:   receiver,
		Sel: ast.NewIdent(s.Selector),
	}
}
