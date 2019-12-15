package config

import (
	"github.com/JhonasMutton/izzi/pkg/client"
	"github.com/hellofresh/health-go"
)

type HealthCheck struct {
	bigIdClient client.BigIDClient
	compline client.ComplineClient
	mongeralAegon client.MongeralAegonClient
}

func NewHealthCheck(bigIdClient client.BigIDClient, compline client.ComplineClient, mongeralAegon client.MongeralAegonClient) HealthCheck {
	return HealthCheck{bigIdClient: bigIdClient, compline: compline, mongeralAegon: mongeralAegon}
}

func (hc *HealthCheck) SetupHealthCheck() (err error) {
	err = health.Register(health.Config{
		Name: "BigId-client",
		Check: func() error {
			e := hc.bigIdClient.Ping()
			return e
		},
	})
	if err!=nil {
		return
	}

	err = health.Register(health.Config{
		Name: "Compline-client",
		Check: func() error {
			e := hc.compline.Ping()
			return e
		},
	})
	if err!=nil {
		return
	}

	err = health.Register(health.Config{
		Name: "MongeralAegon-client",
		Check: func() error {
			e := hc.mongeralAegon.Ping()
			return e
		},
	})

	return
}
