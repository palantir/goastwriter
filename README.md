<p align="right">
<a href="https://autorelease.general.dmz.palantir.tech/palantir/goastwriter"><img src="https://img.shields.io/badge/Perform%20an-Autorelease-success.svg" alt="Autorelease"></a>
</p>

goastwriter 👻✍️
================
[![](https://godoc.org/github.com/palantir/goastwriter?status.svg)](http://godoc.org/github.com/palantir/goastwriter)

`goastwriter` is a library that offers abstractions for defining and writing Go source files
programmatically. It is effectively a convenience wrapper for the structs and functions defined in
the `go/ast` package. However, by providing higher-level abstractions, convenience functions and
tests, `goastwriter` makes writing Go code programatically much simpler for simple use cases.

`goastwriter` processes its generated code using `gofmt`, so its output is `gofmt`-compliant.

Usage
-----
`goastwriter.Write` is the primary function exported by the package and generates the code for a
single Go file. It is provided with the package name that should be used for the file and the
components that make up the file.

Example
-------
Code for generating a Go source file:

```go
out, _ := goastwriter.Write("testpkg",
    decl.NewImports(map[string]string{
        "fmt":       "",
        "go/format": "gofmt",
    }),
    &decl.Struct{
        Name:    "Foo",
        Comment: "Foo is a struct",
        Fields: []decl.StructField{
            {
                Name:    "Bar",
                Type:    expression.StringType,
                Comment: "Bar is a field",
            },
            {
                Name:    "baz",
                Type:    expression.BoolType.Pointer(),
                Comment: "Baz is a field",
            },
        },
    },
    &decl.Function{
        Name: "Bar",
        Params: []*decl.FuncParam{
            decl.NewFuncParam("input", expression.Type("Foo").Pointer()),
        },
        ReturnTypes: []expression.Type{
            expression.Type("Foo").Pointer(),
            expression.ErrorType,
        },
        Body: []astgen.ASTStmt{
            &statement.Expression{
                Expr: expression.NewCallFunction("fmt", "Println"),
            },
            &statement.Expression{
                Expr: expression.NewCallFunction("gofmt", "Source", expression.Nil),
            },
            &statement.Return{
                Values: []astgen.ASTExpr{
                    expression.VariableVal("input"),
                    expression.Nil,
                },
            },
        },
    },
)
fmt.Println(string(out))
```

Output:

```go
package testpkg

import (
	"fmt"
	gofmt "go/format"
)

// Foo is a struct
type Foo struct {
	// Bar is a field
	Bar string
	// Baz is a field
	baz *bool
}

func Bar(input *Foo) (*Foo, error) {
	fmt.Println()
	gofmt.Source(nil)
	return input, nil
}
```
