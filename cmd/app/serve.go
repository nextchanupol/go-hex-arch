package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func serve(addr string, port string, app *fiber.App) {
	go func() {

		// logging.Info(fmt.Sprintf("Starting server on %s:%s ...", addr, port))
		err := app.Listen(fmt.Sprintf("%s:%s", addr, port))
		// logging.Fatal(err.Error())
		log.Fatalf("error opening file: %v", err)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
