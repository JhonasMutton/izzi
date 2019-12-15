package pkg

import (
	"github.com/JhonasMutton/izzi/pkg/api/config"
	"github.com/JhonasMutton/izzi/pkg/api/handlers"
	"github.com/JhonasMutton/izzi/pkg/api/router"
	"github.com/JhonasMutton/izzi/pkg/client"
	"github.com/JhonasMutton/izzi/pkg/usecase"
	"github.com/google/wire"
)

var (
	Container = wire.NewSet(handlers.HandlerSet, usecase.UseCaseSet, client.ClientSet, router.RouteSet, config.AppConfigSet, ApplicationSet)
)
