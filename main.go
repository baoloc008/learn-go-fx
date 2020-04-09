package main

import (
	"go.uber.org/fx"
	"learn-go-fx/serverfx"
)

func main() {
	app := fx.New(
		serverfx.Module,
	)
	app.Run()
}
