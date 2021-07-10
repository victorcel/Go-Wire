package di

import (
	"github.com/victorcel/pruebasGolan/models"
	"github.com/victorcel/pruebasGolan/repositories"
	"github.com/victorcel/pruebasGolan/uc"
)

func providerMessage() models.Message {
	return models.NewMessage()
}

func providerGreeter(message models.Message) repositories.Greeter {
	return repositories.NewGreeter(message)
}

func providerEvent(r repositories.Greeter) (uc.Event, error) {
	return uc.NewEvent(r)
}