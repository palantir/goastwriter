// Copyright (c) 2018 Palantir Technologies. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
