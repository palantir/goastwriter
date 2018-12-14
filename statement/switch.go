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

package statement

import (
	"go/ast"

	"github.com/palantir/goastwriter/astgen"
)

type Switch struct {
	Init       astgen.ASTStmt
	Expression astgen.ASTExpr
	Cases      []CaseClause
}

func (s *Switch) ASTStmt() ast.Stmt {
	ss := &ast.SwitchStmt{}

	if s.Init != nil {
		ss.Init = s.Init.ASTStmt()
	}

	if s.Expression != nil {
		ss.Tag = s.Expression.ASTExpr()
	}

	if len(s.Cases) > 0 {
		var cases []ast.Stmt
		for _, v := range s.Cases {
			cases = append(cases, v.ASTStmt())
		}
		ss.Body = &ast.BlockStmt{
			List: cases,
		}
	}

	return ss
}
