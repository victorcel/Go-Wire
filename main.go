package main

import (
	"fmt"
	"github.com/victorcel/pruebasGolan/di"
	"os"
)

func main() {
	e, err := di.Initialize("Todo Ok")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
	e.Stop()
}
