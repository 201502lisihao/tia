//+build wireinject

package inject

import (
	"github.com/google/wire"

	"my-project/internal"
	"my-project/internal/app"
	"my-project/internal/config"
)

func InitApp(env string) (*internal.App, func(), error) {
	wire.Build(
		internal.NewApp,
		app.WireSet,
		config.WireSet,
	)

	return &internal.App{}, nil, nil
}
