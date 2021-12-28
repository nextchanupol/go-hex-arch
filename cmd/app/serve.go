package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	hlps "github.com/nextchanupol/go-fiber-server/utils"
)

func serve(addr string, port string, app *fiber.App) {
	go func() {

		hlps.Info(fmt.Sprintf("Starting server on %s:%s ...", addr, port))
		err := app.Listen(fmt.Sprintf("%s:%s", addr, port))
		hlps.Fatal(err.Error())
		log.Fatalf("error opening file: %v", err)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
