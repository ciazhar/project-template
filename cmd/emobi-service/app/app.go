package app

import (
	db2 "github.com/ciazhar/emobi-service/internal/db"
	"github.com/ciazhar/emobi-service/third_party/db"
	"github.com/ciazhar/emobi-service/third_party/env"
	logger "github.com/ciazhar/emobi-service/third_party/log"
	"github.com/ciazhar/emobi-service/third_party/sentry"
	"github.com/gofiber/fiber/v2"
	"os"
)

type Application struct {
	DB     *db2.Queries
	Router fiber.Router
}

func Init(e string) (Application, error) {

	//init
	env.Init(e)
	logger.Init()
	sentry.Init()
	dbx := db.Init()

	//set default timezone
	if err := os.Setenv("TZ", "Asia/Jakarta"); err != nil {
		panic(err.Error())
	}

	return Application{
		DB: dbx,
	}, nil
}
