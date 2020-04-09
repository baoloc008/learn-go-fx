package serverfx

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(newServerConfig),
	fx.Invoke(startServer),
)
