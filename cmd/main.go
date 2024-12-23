package main

import (
	"fmt"

	"github.com/ItzB1ack/CalculatorYL-master/internal/application"
)

func main() {
	app := application.New()
	fmt.Println("Сервер запущен на порту 8080")
	app.RunServer()
}
