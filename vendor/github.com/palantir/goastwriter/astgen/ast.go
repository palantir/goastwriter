// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astgen

import (
	"go/ast"
)

type ASTDecl interface {
	ASTDecl() ast.Decl
}

type ASTExpr interface {
	ASTExpr() ast.Expr
}

type ASTSpec interface {
	ASTSpec() ast.Spec
}

type ASTStmt interface {
	ASTStmt() ast.Stmt
}
