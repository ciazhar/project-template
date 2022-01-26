package sentry

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

func Init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("sentry.dsn"),
	})
	if err != nil {
		panic(err)
	}
}
