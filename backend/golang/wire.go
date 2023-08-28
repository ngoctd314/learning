//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
)

func initializeBaz(ctx context.Context) (Baz, error) {
	wire.Build(MegaSet)
	return Baz{}, nil
}
