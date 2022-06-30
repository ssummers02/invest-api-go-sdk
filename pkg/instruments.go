package pkg

import (
	pb "github.com/ssummers02/invest-api-go-sdk/pkg/investapi"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

type InstrumentsInterface interface {
	// TradingSchedules The method of obtaining the trading schedule of trading platforms.
	TradingSchedules(exchange string, from, to *timestamp.Timestamp) ([]*pb.TradingSchedule, error)
	// BondBy The method of obtaining a bond by its identifier.
	BondBy(filters *pb.InstrumentRequest) (*pb.Bond, error)
	// Bonds Method of obtaining a list of bonds.
	Bonds(status pb.InstrumentStatus) ([]*pb.Bond, error)
	// GetBondCoupons Method of obtaining a coupon payment schedule for a bond.
	GetBondCoupons(figi string, from, to *timestamp.Timestamp) ([]*pb.Coupon, error)
	// CurrencyBy The method of obtaining a currency by its identifier.
	CurrencyBy(filters *pb.InstrumentRequest) (*pb.Currency, error)
	// Currencies Method for getting a list of currencies.
	Currencies(status pb.InstrumentStatus) ([]*pb.Currency, error)
	// EtfBy The method of obtaining an investment fund by its identifier.
	EtfBy(filters *pb.InstrumentRequest) (*pb.Etf, error)
	// Etfs Method of obtaining a list of investment funds.
	Etfs(status pb.InstrumentStatus) ([]*pb.Etf, error)
	// FutureBy The method of obtaining futures by its identifier.
	FutureBy(filters *pb.InstrumentRequest) (*pb.Future, error)
	// Futures Method for getting a list of futures.
	Futures(status pb.InstrumentStatus) ([]*pb.Future, error)
	// ShareBy The method of obtaining a stock by its identifier.
	ShareBy(filters *pb.InstrumentRequest) (*pb.Share, error)
	// Shares Method of getting a list of shares.
	Shares(status pb.InstrumentStatus) ([]*pb.Share, error)
	// GetAccruedInterests The method of obtaining the accumulated coupon income on the bond.
	GetAccruedInterests(figi string, from, to *timestamp.Timestamp) ([]*pb.AccruedInterest, error)
	// GetFuturesMargin The method of obtaining the amount of the guarantee for futures.
	GetFuturesMargin(figi string) (*pb.GetFuturesMarginResponse, error)
	// GetInstrumentBy The method of obtaining basic information about the tool.
	GetInstrumentBy(filters *pb.InstrumentRequest) (*pb.Instrument, error)
	// GetDividends A method for obtaining dividend payment events for an instrument.
	GetDividends(figi string, from, to *timestamp.Timestamp) ([]*pb.Dividend, error)
	// GetAssetBy The method of obtaining an asset by its identifier.
	GetAssetBy(assetID string) (*pb.AssetFull, error)
	// GetAssets Method for getting a list of assets.
	GetAssets() ([]*pb.Asset, error)
	// GetFavorites The method of getting the favourite instruments.
	GetFavorites() ([]*pb.FavoriteInstrument, error)
	// EditFavorites The method of editing the selected instruments.
	EditFavorites(newFavourites *pb.EditFavoritesRequest) ([]*pb.FavoriteInstrument, error)
	// GetSharesBase Get a list of stocks available for trading via API
	GetSharesBase() ([]*pb.Share, error)
	// GetETFsBase Get a list of investment funds available for trading via API
	GetETFsBase() ([]*pb.Etf, error)
	// GetBondsBase Get a list of bonds available for trading via API
	GetBondsBase() ([]*pb.Bond, error)
	// GetFuturesBase Get a list of futures available for trading via API
	GetFuturesBase() ([]*pb.Future, error)
}

type InstrumentsService struct {
	client pb.InstrumentsServiceClient
	config Config
}

func NewInstrumentsService(conn *grpc.ClientConn, cfg Config) *InstrumentsService {
	client := pb.NewInstrumentsServiceClient(conn)

	return &InstrumentsService{client: client, config: cfg}
}

func (is *InstrumentsService) TradingSchedules(exchange string, from, to *timestamp.Timestamp) ([]*pb.TradingSchedule, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.TradingSchedules(ctx, &pb.TradingSchedulesRequest{
		Exchange: exchange,
		From:     from,
		To:       to,
	})
	if err != nil {
		return nil, err
	}

	return res.Exchanges, nil
}

func (is *InstrumentsService) BondBy(filters *pb.InstrumentRequest) (*pb.Bond, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.BondBy(ctx, filters)
	if err != nil {
		return nil, err
	}

	return res.Instrument, nil
}

func (is *InstrumentsService) Bonds(status pb.InstrumentStatus) ([]*pb.Bond, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.Bonds(ctx, &pb.InstrumentsRequest{
		InstrumentStatus: status,
	})
	if err != nil {
		return nil, err
	}

	return res.Instruments, nil
}

func (is *InstrumentsService) GetBondsBase() ([]*pb.Bond, error) {
	return is.Bonds(pb.InstrumentStatus_INSTRUMENT_STATUS_BASE)
}

func (is *InstrumentsService) GetBondCoupons(figi string, from, to *timestamp.Timestamp) ([]*pb.Coupon, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.GetBondCoupons(ctx, &pb.GetBondCouponsRequest{
		Figi: figi,
		From: from,
		To:   to,
	})
	if err != nil {
		return nil, err
	}

	return res.Events, nil
}

func (is *InstrumentsService) CurrencyBy(filters *pb.InstrumentRequest) (*pb.Currency, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.CurrencyBy(ctx, filters)
	if err != nil {
		return nil, err
	}

	return res.Instrument, nil
}

func (is *InstrumentsService) Currencies(status pb.InstrumentStatus) ([]*pb.Currency, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.Currencies(ctx, &pb.InstrumentsRequest{
		InstrumentStatus: status,
	})
	if err != nil {
		return nil, err
	}

	return res.Instruments, nil
}

func (is *InstrumentsService) EtfBy(filters *pb.InstrumentRequest) (*pb.Etf, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.EtfBy(ctx, filters)
	if err != nil {
		return nil, err
	}

	return res.Instrument, nil
}

func (is *InstrumentsService) Etfs(status pb.InstrumentStatus) ([]*pb.Etf, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.Etfs(ctx, &pb.InstrumentsRequest{
		InstrumentStatus: status,
	})
	if err != nil {
		return nil, err
	}

	return res.Instruments, nil
}

func (is *InstrumentsService) GetETFsBase() ([]*pb.Etf, error) {
	return is.Etfs(pb.InstrumentStatus_INSTRUMENT_STATUS_BASE)
}

func (is *InstrumentsService) FutureBy(filters *pb.InstrumentRequest) (*pb.Future, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.FutureBy(ctx, filters)
	if err != nil {
		return nil, err
	}

	return res.Instrument, nil
}

func (is *InstrumentsService) Futures(status pb.InstrumentStatus) ([]*pb.Future, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.Futures(ctx, &pb.InstrumentsRequest{
		InstrumentStatus: status,
	})
	if err != nil {
		return nil, err
	}

	return res.Instruments, nil
}

func (is *InstrumentsService) GetFuturesBase() ([]*pb.Future, error) {
	return is.Futures(pb.InstrumentStatus_INSTRUMENT_STATUS_BASE)
}

func (is *InstrumentsService) ShareBy(filters *pb.InstrumentRequest) (*pb.Share, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.ShareBy(ctx, filters)
	if err != nil {
		return nil, err
	}

	return res.Instrument, nil
}

func (is *InstrumentsService) Shares(status pb.InstrumentStatus) ([]*pb.Share, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.Shares(ctx, &pb.InstrumentsRequest{
		InstrumentStatus: status,
	})
	if err != nil {
		return nil, err
	}

	return res.Instruments, nil
}

func (is *InstrumentsService) GetSharesBase() ([]*pb.Share, error) {
	return is.Shares(pb.InstrumentStatus_INSTRUMENT_STATUS_BASE)
}

func (is *InstrumentsService) GetAccruedInterests(figi string, from, to *timestamp.Timestamp) ([]*pb.AccruedInterest, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.GetAccruedInterests(ctx, &pb.GetAccruedInterestsRequest{
		Figi: figi,
		From: from,
		To:   to,
	})
	if err != nil {
		return nil, err
	}

	return res.AccruedInterests, nil
}

func (is *InstrumentsService) GetFuturesMargin(figi string) (*pb.GetFuturesMarginResponse, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.GetFuturesMargin(ctx, &pb.GetFuturesMarginRequest{
		Figi: figi,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (is *InstrumentsService) GetInstrumentBy(filters *pb.InstrumentRequest) (*pb.Instrument, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.GetInstrumentBy(ctx, filters)
	if err != nil {
		return nil, err
	}

	return res.Instrument, nil
}

func (is *InstrumentsService) GetDividends(figi string, from, to *timestamp.Timestamp) ([]*pb.Dividend, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.GetDividends(ctx, &pb.GetDividendsRequest{
		Figi: figi,
		From: from,
		To:   to,
	})
	if err != nil {
		return nil, err
	}

	return res.Dividends, nil
}

func (is *InstrumentsService) GetAssetBy(assetID string) (*pb.AssetFull, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.GetAssetBy(ctx, &pb.AssetRequest{
		Id: assetID,
	})
	if err != nil {
		return nil, err
	}

	return res.Asset, nil
}

func (is *InstrumentsService) GetAssets() ([]*pb.Asset, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.GetAssets(ctx, &pb.AssetsRequest{})
	if err != nil {
		return nil, err
	}

	return res.Assets, nil
}

func (is *InstrumentsService) GetFavorites() ([]*pb.FavoriteInstrument, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.GetFavorites(ctx, &pb.GetFavoritesRequest{})
	if err != nil {
		return nil, err
	}

	return res.FavoriteInstruments, nil
}

func (is *InstrumentsService) EditFavorites(newFavourites *pb.EditFavoritesRequest) ([]*pb.FavoriteInstrument, error) {
	ctx, cancel := CreateRequestContext(is.config)
	defer cancel()

	res, err := is.client.EditFavorites(ctx, newFavourites)
	if err != nil {
		return nil, err
	}

	return res.FavoriteInstruments, nil
}
