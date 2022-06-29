package pkg

import (
	pb "github.com/ssummers02/invest-api-go-sdk/pkg/investapi"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

type StopOrdersInterface interface {
	// PostStopOrder The method of placing a stop order.
	PostStopOrder(stopOrder *pb.PostStopOrderRequest) (string, error)
	// GetStopOrders Method for getting a list of active stop orders on the account.
	GetStopOrders(accountID string) ([]*pb.StopOrder, error)
	// CancelStopOrder The method of canceling the stop order.
	CancelStopOrder(accountID string, stopOrderID string) (*timestamp.Timestamp, error)
}

type StopOrdersService struct {
	client pb.StopOrdersServiceClient
	config Config
}

func NewStopOrdersService(conn *grpc.ClientConn, config Config) *StopOrdersService {
	client := pb.NewStopOrdersServiceClient(conn)

	return &StopOrdersService{client: client, config: config}
}

func (ss *StopOrdersService) PostStopOrder(stopOrder *pb.PostStopOrderRequest) (string, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.PostStopOrder(ctx, stopOrder)
	if err != nil {
		return "", err
	}

	return res.StopOrderId, nil
}

func (ss *StopOrdersService) GetStopOrders(accountID string) ([]*pb.StopOrder, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.GetStopOrders(ctx, &pb.GetStopOrdersRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}

	return res.StopOrders, nil
}

func (ss *StopOrdersService) CancelStopOrder(accountID string, stopOrderID string) (*timestamp.Timestamp, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.CancelStopOrder(ctx, &pb.CancelStopOrderRequest{
		AccountId:   accountID,
		StopOrderId: stopOrderID,
	})
	if err != nil {
		return nil, err
	}

	return res.Time, nil
}
