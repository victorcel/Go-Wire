package repositories

import (
	"github.com/victorcel/pruebasGolan/models"
	"time"
)

type Greeter struct {
	Message models.Message
	Grumpy  bool
}

func NewGreeter(m models.Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() models.Message {
	if g.Grumpy {
		return models.Message("Go away!")
	}
	return g.Message
}

func (g Greeter) Sms(text string) models.Message {
	if len(text)<1 {
		return models.Message("Vacio!")
	}
	 return models.Message(text)
}