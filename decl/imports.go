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
	"fmt"
	"go/ast"
	"go/token"
	"sort"
)

type Imports []*Import

type Import struct {
	Alias string
	Path  string
}

func NewImports(importsToAliases map[string]string) Imports {
	var imports []*Import

	// sort by import name to make order deterministic
	var sortedKeys []string
	for k := range importsToAliases {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, currImport := range sortedKeys {
		currAlias := importsToAliases[currImport]
		imports = append(imports, &Import{
			Alias: currAlias,
			Path:  currImport,
		})
	}
	return Imports(imports)
}

func (s Imports) ASTDecl() ast.Decl {
	var specs []ast.Spec
	for _, currImport := range s {
		var alias *ast.Ident
		if currImport.Alias != "" {
			alias = ast.NewIdent(currImport.Alias)
		}
		specs = append(specs, &ast.ImportSpec{
			Name: alias,
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("%q", currImport.Path),
			},
		})
	}

	return &ast.GenDecl{
		Tok:    token.IMPORT,
		Lparen: token.Pos(1),
		Specs:  specs,
	}
}
