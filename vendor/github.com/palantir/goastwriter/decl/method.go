// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decl

import (
	"go/ast"

	"github.com/palantir/goastwriter/expression"
)

type Method struct {
	Function
	ReceiverName string
	ReceiverType expression.Type
}

func (m *Method) ASTDecl() ast.Decl {
	return m.Function.funcDecl(&ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(m.ReceiverName),
		},
		Type: m.ReceiverType.ToIdent(),
	})
}
