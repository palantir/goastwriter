// Copyright 2018 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"
)

type BuiltIn string

const (
	AppendBuiltIn BuiltIn = "append"
	CopyBuiltIn   BuiltIn = "copy"
	LenBuiltIn    BuiltIn = "len"
	MakeBuiltIn   BuiltIn = "make"
	NewBuiltIn    BuiltIn = "new"
	PanicBuiltIn  BuiltIn = "panic"
)

func (c BuiltIn) ASTExpr() ast.Expr {
	return ast.NewIdent(string(c))
}
