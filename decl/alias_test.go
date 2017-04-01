// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decl_test

import (
	"testing"

	"github.com/palantir/goastwriter/decl"
	"github.com/palantir/goastwriter/expression"
)

func TestAlias(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple alias",
			val: &decl.Alias{
				Name:    "StatusCode",
				Comment: "An HTTP response code",
				Type:    expression.IntType,
			},
			want: `type
// An HTTP response code
StatusCode int`,
		},
	})
}
