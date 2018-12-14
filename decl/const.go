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

package decl

import (
	"go/ast"
	"go/token"

	"github.com/palantir/goastwriter/astgen"
	"github.com/palantir/goastwriter/expression"
	"github.com/palantir/goastwriter/spec"
)

type Const struct {
	Values []*spec.Value
}

func NewConst(values ...*spec.Value) *Const {
	return &Const{
		Values: values,
	}
}

func NewConstSingleValue(name string, typ expression.Type, value astgen.ASTExpr) *Const {
	return &Const{
		Values: []*spec.Value{
			spec.NewValue(name, typ, value),
		},
	}
}

func (c *Const) ASTDecl() ast.Decl {
	var specs []ast.Spec
	for _, val := range c.Values {
		specs = append(specs, val.ASTSpec())
	}
	constDecl := &ast.GenDecl{
		Tok:   token.CONST,
		Specs: specs,
	}
	if len(specs) > 1 {
		// set Lparen to non-0 value to ensure that parenthesis are rendered
		constDecl.Lparen = token.Pos(1)
	}
	return constDecl
}
