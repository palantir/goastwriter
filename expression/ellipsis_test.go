// Copyright 2020 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestEllipsisExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "ellipsis",
			val:  expression.Ellipsis{},
			want: `...`,
		},
	})
}
