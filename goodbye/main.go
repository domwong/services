package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/services/goodbye/handler"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("goodbye"),
	)

	// Register Handler
	srv.Handle(new(handler.Goodbye))

	// Run the service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
