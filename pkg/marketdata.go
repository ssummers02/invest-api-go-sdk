package pkg

import (
	pb "github.com/ssummers02/invest-api-go-sdk/pkg/investapi"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

type MarketDataInterface interface {
	// GetCandles The method of requesting historical candlesticks by instrument.
	GetCandles(figi string, from, to *timestamp.Timestamp, interval pb.CandleInterval) ([]*pb.HistoricCandle, error)
	// GetLastPrices The method of requesting the latest prices for instruments.
	GetLastPrices(figi []string) ([]*pb.LastPrice, error)
	// GetOrderBook The method of obtaining a glass by instrument.
	GetOrderBook(figi string, depth int) (*pb.GetOrderBookResponse, error)
	// GetTradingStatus The method of requesting the status of trading on instruments.
	GetTradingStatus(figi string) (*pb.GetTradingStatusResponse, error)
	// GetLastTrades The method of requesting the latest depersonalized transactions on the instrument.
	GetLastTrades(figi string, from, to *timestamp.Timestamp) ([]*pb.Trade, error)
}

type MarketDataService struct {
	client pb.MarketDataServiceClient
	config Config
}

func NewMarketDataService(conn *grpc.ClientConn, config Config) *MarketDataService {
	client := pb.NewMarketDataServiceClient(conn)

	return &MarketDataService{
		client: client,
		config: config,
	}
}

func (ms *MarketDataService) GetCandles(figi string, from, to *timestamp.Timestamp, interval pb.CandleInterval) ([]*pb.HistoricCandle, error) {
	ctx, cancel := CreateRequestContext(ms.config)
	defer cancel()

	res, err := ms.client.GetCandles(ctx, &pb.GetCandlesRequest{
		Figi:     figi,
		From:     from,
		To:       to,
		Interval: interval,
	})
	if err != nil {
		return nil, err
	}

	return res.Candles, nil
}

func (ms *MarketDataService) GetLastPrices(figi []string) ([]*pb.LastPrice, error) {
	ctx, cancel := CreateRequestContext(ms.config)
	defer cancel()

	res, err := ms.client.GetLastPrices(ctx, &pb.GetLastPricesRequest{
		Figi: figi,
	})
	if err != nil {
		return nil, err
	}

	return res.LastPrices, nil
}

func (ms *MarketDataService) GetOrderBook(figi string, depth int) (*pb.GetOrderBookResponse, error) {
	ctx, cancel := CreateRequestContext(ms.config)
	defer cancel()

	res, err := ms.client.GetOrderBook(ctx, &pb.GetOrderBookRequest{
		Figi:  figi,
		Depth: int32(depth),
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ms *MarketDataService) GetTradingStatus(figi string) (*pb.GetTradingStatusResponse, error) {
	ctx, cancel := CreateRequestContext(ms.config)
	defer cancel()

	res, err := ms.client.GetTradingStatus(ctx, &pb.GetTradingStatusRequest{
		Figi: figi,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ms *MarketDataService) GetLastTrades(figi string, from, to *timestamp.Timestamp) ([]*pb.Trade, error) {
	ctx, cancel := CreateRequestContext(ms.config)
	defer cancel()

	res, err := ms.client.GetLastTrades(ctx, &pb.GetLastTradesRequest{
		Figi: figi,
		From: from,
		To:   to,
	})
	if err != nil {
		return nil, err
	}

	return res.Trades, nil
}
