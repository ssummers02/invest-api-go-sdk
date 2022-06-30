package pkg

import (
	"errors"

	pb "github.com/ssummers02/invest-api-go-sdk/pkg/investapi"
)

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

func (sp *ServicePool) GetLastPricesForAll() ([]*pb.LastPrice, error) {
	base, err := sp.GetSharesBase()
	if err != nil {
		return nil, err
	}
	if len(base) == 0 {
		return nil, errors.New("no shares found")
	}

	figis := make([]string, len(base))
	for i, instrument := range base {
		figis[i] = instrument.Figi
	}
	return sp.GetLastPrices(figis)
}
