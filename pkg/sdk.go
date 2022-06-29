package pkg

// ServicePool is a ready-to-use scope for all available non-stream services.
type ServicePool struct {
	UsersInterface
	InstrumentsInterface
	OrderStreamInterface
	OrdersInterface
	OperationsInterface
	MarketDataInterface
	MarketDataStreamInterface
	StopOrdersInterface
	SandboxInterface
}

func NewServicePool(cfg Config) (*ServicePool, error) {
	conn, err := CreateClientConn()
	if err != nil {
		return nil, err
	}

	ordersStreamService, err := NewOrdersStreamService(conn, cfg)
	if err != nil {
		return nil, err
	}

	marketDataStreamService, err := NewMarketDataStream(conn, cfg)
	if err != nil {
		return nil, err
	}

	return &ServicePool{
		UsersInterface:            NewUsersService(conn, cfg),
		InstrumentsInterface:      NewInstrumentsService(conn, cfg),
		OrderStreamInterface:      ordersStreamService,
		OrdersInterface:           NewOrdersService(conn, cfg),
		OperationsInterface:       NewOperationsService(conn, cfg),
		MarketDataInterface:       NewMarketDataService(conn, cfg),
		MarketDataStreamInterface: marketDataStreamService,
		StopOrdersInterface:       NewStopOrdersService(conn, cfg),
		SandboxInterface:          NewSandboxService(conn, cfg),
	}, nil
}
