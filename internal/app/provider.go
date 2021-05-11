package app

import (
	"github.com/google/wire"

	"my-project/internal/app/server"
)

var WireSet = wire.NewSet(
	NewServerRunner,
	server.NewApiServer,
)
