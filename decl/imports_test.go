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
