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

package comment

import (
	"fmt"
	"go/ast"
	"strings"
)

// ToComment converts the provided string into an *ast.CommentGroup that represents a single-line comment that occurs on
// its own line. If the provided string is multiple lines (that is, if it contains at least one '\n'), then each line is
// turned into its own comment line.
func ToComment(content string) *ast.CommentGroup {
	content = strings.TrimRight(content, "\n")
	if content == "" {
		return nil
	}

	var comments []*ast.Comment
	for i, commentLine := range strings.Split(content, "\n") {
		fmtStr := "// %s"
		if i == 0 {
			fmtStr = fmt.Sprintf("\n%s", fmtStr)
		}
		comments = append(comments, &ast.Comment{
			Text: fmt.Sprintf(fmtStr, commentLine),
		})
	}

	return &ast.CommentGroup{
		List: comments,
	}
}
