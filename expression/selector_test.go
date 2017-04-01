// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestSelectorExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple selector expression",
			val: &expression.Selector{
				Receiver: expression.VariableVal("foo"),
				Selector: "Bar",
			},
			want: `foo.Bar`,
		},
	})
}
