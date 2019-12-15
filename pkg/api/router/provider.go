package router

import "github.com/google/wire"

var RouteSet =  wire.NewSet(NewRoutes)
