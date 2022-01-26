package main

import (
	"github.com/ciazhar/emobi-service/cmd/emobi-service/app"
	"github.com/ciazhar/emobi-service/cmd/emobi-service/router"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {

	ca := &cli.App{
		Name:    viper.GetString("name"),
		Version: viper.GetString("version"),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "env",
				Aliases: []string{"e"},
				Usage:   "Load configuration from `FILE`",
			},
		},
		Action: func(c *cli.Context) error {
			env := c.String("env")
			if env == "" {
				env = "default"
			}
			viper.SetConfigName(env)
			a, err := app.Init(env)
			if err != nil {
				panic(err)
			}
			return router.Init(a)
		},
	}

	if err := ca.Run(os.Args); err != nil {
		panic(err.Error())
	}
}
