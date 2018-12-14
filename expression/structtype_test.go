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

package expression_test

import (
	"testing"

	"github.com/palantir/goastwriter/expression"
)

func TestStructTypeExpression(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "struct type",
			val: &expression.StructType{
				Fields: []*expression.StructField{
					{
						Name:    "Bar",
						Type:    "string",
						Comment: "Bar is a field",
					},
					{
						Name:    "baz",
						Type:    expression.Type("bool").Pointer(),
						Comment: "Baz is a field",
					},
				},
			},
			want: `struct {
// Bar is a field
	Bar	string
// Baz is a field
	baz	*bool
}`,
		},
	})
}
