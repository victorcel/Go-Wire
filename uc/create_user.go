package uc

import (
	"errors"
	"fmt"
	"github.com/victorcel/pruebasGolan/repositories"
)

var (
	ErrorAgeIsInvalid = errors.New("age_is_invalid")
)

type Event struct {
	Greeter repositories.Greeter // <- adding a Greeter field
}

func NewEvent(g repositories.Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func (e Event) Stop()  {
	msg := e.Greeter.Sms("Sms Enviado")
	fmt.Println(msg)
}