package main

import (
	"invest-api-go-sdk/internal"
)

// ServicePool is a ready-to-use scope for all available non-stream services.
type ServicePool struct {
	internal.UsersService
	internal.InstrumentsService
}

func NewServicePool() (*ServicePool, error) {
	conn, err := internal.CreateClientConn()
	if err != nil {
		return nil, err
	}

	cfg := internal.TradeBotConfig{
		IsSandbox: true,
		Token:     "t.A8z_9kqohPFcTWmgthruzrMGEoumk-xvNqfSkcJpTREr_UjOz9G97WR3QgtCCJqjM0IycPibRuGhctFQd616uA",
		AccountID: "invest-api-go-sdk-test-account",
	}

	return &ServicePool{
		UsersService:       *internal.NewUsersService(conn, cfg),
		InstrumentsService: *internal.NewInstrumentsService(conn, cfg),
	}, nil
}
