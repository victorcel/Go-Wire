// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"github.com/victorcel/pruebasGolan/uc"
)

// Injectors from wire.go:

func Initialize() (uc.Event, error) {
	message := providerMessage()
	greeter := providerGreeter(message)
	event, err := providerEvent(greeter)
	if err != nil {
		return uc.Event{}, err
	}
	return event, nil
}
