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

package decl_test

import (
	"testing"

	"github.com/palantir/goastwriter/decl"
)

func TestImports(t *testing.T) {
	runTest(t, []testCase{
		{
			name: "single import",
			val: decl.NewImports(map[string]string{
				"fmt": "",
			}),
			want: `import (
	"fmt"
)`,
		},
		{
			name: "multiple imports",
			val: decl.NewImports(map[string]string{
				"fmt":     "",
				"strings": "",
			}),
			want: `import (
	"fmt"
	"strings"
)`,
		},
		{
			name: "imports with aliases",
			val: decl.NewImports(map[string]string{
				"fmt":     "alias",
				"strings": "",
			}),
			want: `import (
	alias "fmt"
	"strings"
)`,
		},
	})
}
