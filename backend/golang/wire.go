//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

// func initializeBaz(ctx context.Context) (Baz, error) {
// 	wire.Build(MegaSet)
// 	return Baz{}, nil
// }

// Injectors from wire.go:
func initializeBar() string {
	wire.Build(Set)
	return ""
}
