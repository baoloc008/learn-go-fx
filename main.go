package main

import "go.uber.org/fx"

func main() {
	app := fx.New()
	app.Run()
}