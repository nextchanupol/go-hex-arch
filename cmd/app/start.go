package app

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	repository "github.com/nextchanupol/go-fiber-server/internal/core/repositories"
	custsrv "github.com/nextchanupol/go-fiber-server/internal/core/services/adapters"
	customerhdl "github.com/nextchanupol/go-fiber-server/internal/handlers"
	"github.com/spf13/viper"
)

func Start() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	initMiddlewares(app)

	dbClient := initDB()

	customerRepository := repository.NewCustomerRepositoryDB(dbClient)
	_ = customerRepository

	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryMock

	custService := custsrv.NewCustomerService(customerRepository)

	custHdl := customerhdl.NewCustomerHandler(custService)

	v1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		c.Set("version", "1.0.0")
		return c.Next()
	})

	v1.Get("/customers", custHdl.GetCustomers)
	v1.Get("/customers/:customerID", custHdl.GetCustomerByID)

	app.Server().MaxConnsPerIP = 1000

	// app.Get("/monitor", monitor.New())

	serve(viper.GetString("app.addr"), strconv.Itoa(viper.GetInt("app.port")), app)

}
