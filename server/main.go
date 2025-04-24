package main

import (
	"os"

	"github.com/kylep342/mendel/src/app"
)

func main() {
	if os.Getenv("SERVER_PORT") == "" {
		panic("env variable 'SERVER_PORT' must be set")
	}

	a := app.App{}
	a.Initialize()
	a.Run()
}
