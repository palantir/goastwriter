// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comment

import (
	"go/ast"
)

// ToComment converts the provided string into an *ast.CommentGroup that represents a single-line comment that occurs on
// its own line.
func ToComment(content string) *ast.CommentGroup {
	if content == "" {
		return nil
	}
	return &ast.CommentGroup{
		List: []*ast.Comment{
			{
				Text: "\n// " + content,
			},
		},
	}
}
