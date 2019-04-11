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
