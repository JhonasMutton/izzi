package config

import "github.com/google/wire"

var AppConfigSet = wire.NewSet(NewHealthCheck)
