package di

import (
	"errors"
	"github.com/victorcel/pruebasGolan/models"
	"github.com/victorcel/pruebasGolan/repositories"
	"github.com/victorcel/pruebasGolan/uc"
)

func providerMessage(phrase string) (models.Message, error) {
	if phrase == "" {
		return "", errors.New("ProviderMessage no puede ser vacio")
	}
	return models.NewMessage(phrase), nil
}

func providerGreeter(message models.Message) repositories.Greeter {
	return repositories.NewGreeter(message)
}

func providerEvent(r repositories.Greeter) (uc.Event, error) {
	return uc.NewEvent(r)
}
