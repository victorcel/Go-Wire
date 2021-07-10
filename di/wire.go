//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/victorcel/pruebasGolan/uc"
)

func Initialize()  (uc.Event, error) {
	wire.Build(
	providerEvent,
	providerMessage,
	providerGreeter,
	)
	return uc.Event{}, nil
}
