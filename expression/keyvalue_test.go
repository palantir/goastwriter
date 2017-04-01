// Copyright 2017 Palantir Technologies. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestKeyValueExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "simple key value expression",
			val: &expression.KeyValue{
				Key: &expression.Selector{
					Receiver: expression.VariableVal("http"),
					Selector: "Client",
				},
				Value: expression.StringVal("foo"),
			},
			want: `http.Client: "foo"`,
		},
	})
}
