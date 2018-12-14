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
