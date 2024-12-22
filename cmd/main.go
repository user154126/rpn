package main

import (
	"fmt"

	"github.com/user154126/rpn/internal/application"
)

func main() {
	app := application.New()
	fmt.Println("RunServer")
	// app.Run()
	app.RunServer()
}