package main

import (
	"invest-api-go-sdk/internal"
)

// ServicePool is a ready-to-use scope for all available non-stream services.
type ServicePool struct {
	internal.UsersInterface
	internal.InstrumentsInterface
	internal.OrderStreamInterface
	internal.OrdersInterface
}

func NewServicePool() (*ServicePool, error) {
	conn, err := internal.CreateClientConn()
	if err != nil {
		return nil, err
	}

	cfg := internal.TradeBotConfig{
		IsSandbox: true,
		Token:     "t.A8z_9kqohPFcTWmgthruzrMGEoumk-xvNqfSkcJpTREr_UjOz9G97WR3QgtCCJqjM0IycPibRuGhctFQd616uA",
		AccountID: []string{"invest-api-go-sdk-test-account"},
	}

	ordersStreamService, err := internal.NewOrdersStreamService(conn, cfg)
	if err != nil {
		return nil, err
	}

	return &ServicePool{
		UsersInterface:       internal.NewUsersService(conn, cfg),
		InstrumentsInterface: internal.NewInstrumentsService(conn, cfg),
		OrderStreamInterface: ordersStreamService,
		OrdersInterface:      internal.NewOrdersService(conn, cfg),
	}, nil
}
