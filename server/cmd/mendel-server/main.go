// main.go
package main

import (
	"github.com/kylep342/mendel/internal/app"
	"github.com/kylep342/mendel/internal/constants"
)

func main() {
	env := constants.LoadEnv()

	a := app.App{}
	a.Initialize(env)

	a.Run(env)
}
