// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type KeyValue struct {
	Key   astgen.ASTExpr
	Value astgen.ASTExpr
}

func NewKeyValue(key string, value astgen.ASTExpr) *KeyValue {
	return &KeyValue{
		Key:   VariableVal(key),
		Value: value,
	}
}

func (k *KeyValue) ASTExpr() ast.Expr {
	return &ast.KeyValueExpr{
		Key:   k.Key.ASTExpr(),
		Value: k.Value.ASTExpr(),
	}
}
