package common

import (
	"oak-compiler/internal/pkg/ast"
)

type Error struct {
	Location ast.Location
	Extra    []ast.Location
	Message  string
}

type SystemError struct {
	Message string
}
