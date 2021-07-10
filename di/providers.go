package di

import (
	"github.com/victorcel/pruebasGolan/models"
	"github.com/victorcel/pruebasGolan/repositories"
	"github.com/victorcel/pruebasGolan/uc"
)

func providerMessage(phrase string) models.Message {
	return models.NewMessage(phrase)
}

func providerGreeter(message models.Message) repositories.Greeter {
	return repositories.NewGreeter(message)
}

func providerEvent(r repositories.Greeter) (uc.Event, error) {
	return uc.NewEvent(r)
}