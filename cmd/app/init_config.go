package app

import (
	_ "database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	initTimeLocation()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func initTimeLocation() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	time.Local = ict
}

func initDB() *sqlx.DB {
	datasource := fmt.Sprintf("%v", viper.GetString("db.host"))

	dbClient, err := sqlx.Open(viper.GetString("db.driver"), datasource)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbClient.SetConnMaxLifetime(time.Minute * 5)
	dbClient.SetMaxOpenConns(viper.GetInt("db.max_open_conn"))
	dbClient.SetMaxIdleConns(viper.GetInt("db.max_idle_conn"))

	return dbClient
}

func initMiddlewares(app *fiber.App) {
	app.Use(logger.New(
		logger.Config{
			Format:     "[${time}] | ${pid} | ${locals:requestid} | ${status} - ${method} | ${path}\n",
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "Asia/Bangkok",
		},
	))
	app.Use(etag.New())
	app.Use(requestid.New(requestid.Config{
		Header: fiber.HeaderXRequestID,
		Generator: func() string {
			return utils.UUID()
		},
		ContextKey: "requestid",
	}))
}
