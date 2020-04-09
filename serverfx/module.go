package serverfx

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewServerConfig),
	fx.Invoke(startServer),
)
