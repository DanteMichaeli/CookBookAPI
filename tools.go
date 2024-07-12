//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
)

// run the following in the terminal to generate the models_gen.go file:
// go run github.com/99designs/gqlgen generate
