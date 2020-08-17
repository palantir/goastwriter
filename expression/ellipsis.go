// Copyright 2020 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"
)

type Ellipsis struct{}

func (Ellipsis) ASTExpr() ast.Expr {
	return &ast.Ellipsis{}
}
