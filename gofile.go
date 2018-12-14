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

package goastwriter

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/printer"
	"go/token"

	"github.com/palantir/goastwriter/astgen"
)

func Write(pkgName string, components ...astgen.ASTDecl) ([]byte, error) {
	if pkgName == "" {
		return nil, fmt.Errorf("PackageName must be non-empty")
	}

	var decls []ast.Decl
	for _, currDecler := range components {
		decls = append(decls, currDecler.ASTDecl())
	}

	var buf bytes.Buffer
	if err := printer.Fprint(&buf, token.NewFileSet(), &ast.File{
		Name:  ast.NewIdent(pkgName),
		Decls: decls,
	}); err != nil {
		return nil, fmt.Errorf("failed to convert AST to source code: %v", err)
	}
	return format.Source(buf.Bytes())
}
