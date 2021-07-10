package models

type Message string

func NewMessage() Message {
	return Message("Hi there!")
}