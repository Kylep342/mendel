package main

import (
	"os"

	"github.com/Kylep342/mendel/app"
)

func main() {
	if os.Getenv("SERVER_PORT") == "" {
		panic("env variable 'SERVER_PORT' must be set")
	}

	a := app.App{}
	a.Initialize()
	a.Run()
}
