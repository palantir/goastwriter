// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spec

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/comment"
	"github.com/palantir/goastwriter/expression"
)

type Value struct {
	Comment string
	Names   []string
	Type    expression.Type
	Values  []astgen.ASTExpr
}

func NewValue(name string, typ expression.Type, values ...astgen.ASTExpr) *Value {
	return &Value{
		Names:  []string{name},
		Type:   typ,
		Values: values,
	}
}

func (v *Value) ASTSpec() ast.Spec {
	var names []*ast.Ident
	for _, name := range v.Names {
		names = append(names, ast.NewIdent(name))
	}
	spec := &ast.ValueSpec{
		Doc:   comment.ToComment(v.Comment),
		Names: names,
	}
	if len(v.Type) > 0 {
		spec.Type = v.Type.ToIdent()
	}
	if len(v.Values) > 0 {
		var values []ast.Expr
		for _, value := range v.Values {
			values = append(values, value.ASTExpr())
		}
		spec.Values = values
	}
	return spec
}
