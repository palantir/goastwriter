// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"fmt"
	"go/ast"
	"go/token"
)

type StringVal string

func (s StringVal) ASTExpr() ast.Expr {
	return &ast.BasicLit{
		Kind:  token.STRING,
		Value: fmt.Sprintf("%q", s),
	}
}

type IntVal int

func (i IntVal) ASTExpr() ast.Expr {
	return &ast.BasicLit{
		Kind:  token.INT,
		Value: fmt.Sprintf("%d", i),
	}
}

type VariableVal string

func (v VariableVal) ASTExpr() ast.Expr {
	return ast.NewIdent(string(v))
}

var Nil = VariableVal("nil")

var Blank = VariableVal("_")
