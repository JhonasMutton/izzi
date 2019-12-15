package client

import "github.com/google/wire"

var ClientSet = wire.NewSet(NewMongeralAegonClient, wire.Bind(new(IMongeralAegonClient), new(MongeralAegonClient)),
	NewComplineClientt, wire.Bind(new(IComplineClient), new(ComplineClient)),
	NewBigIDClient, wire.Bind(new(IBigIDClient), new(BigIDClient)), )
