//+build wireinject

package di

import "github.com/google/wire"

var stdSet = wire.NewSet(
	providerEvent,
	providerMessage,
	providerGreeter,
	)
