//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/victorcel/pruebasGolan/uc"
)

func Initialize(phrase string) (uc.Event, error) {
	wire.Build( stdSet)
	return uc.Event{}, nil
}
