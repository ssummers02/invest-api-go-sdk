package pkg

import (
	pb "github.com/ssummers02/invest-api-go-sdk/pkg/investapi"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

type OrdersInterface interface {
	// PostOrder The method of submitting the order.
	PostOrder(order *pb.PostOrderRequest) (*pb.PostOrderResponse, error)
	// CancelOrder The method of cancellation of the trade order.
	CancelOrder(accountID string, orderID string) (*timestamp.Timestamp, error)
	// GetOrderState The method of obtaining the status of a trade order.
	GetOrderState(accountID string, orderID string) (*pb.OrderState, error)
	// GetOrders The method of getting a list of active orders for the account.
	GetOrders(accountID string) ([]*pb.OrderState, error)
}

type OrdersService struct {
	client pb.OrdersServiceClient
	config Config
}

func NewOrdersService(conn *grpc.ClientConn, config Config) *OrdersService {
	client := pb.NewOrdersServiceClient(conn)

	return &OrdersService{
		client: client,
		config: config,
	}
}

func (os *OrdersService) PostOrder(order *pb.PostOrderRequest) (*pb.PostOrderResponse, error) {
	ctx, cancel := CreateRequestContext(os.config)

	defer cancel()

	res, err := os.client.PostOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (os *OrdersService) CancelOrder(accountID string, orderID string) (*timestamp.Timestamp, error) {
	ctx, cancel := CreateRequestContext(os.config)

	defer cancel()

	res, err := os.client.CancelOrder(ctx, &pb.CancelOrderRequest{
		AccountId: accountID,
		OrderId:   orderID,
	})
	if err != nil {
		return nil, err
	}

	return res.Time, nil
}

func (os *OrdersService) GetOrderState(accountID string, orderID string) (*pb.OrderState, error) {
	ctx, cancel := CreateRequestContext(os.config)

	defer cancel()

	res, err := os.client.GetOrderState(ctx, &pb.GetOrderStateRequest{
		AccountId: accountID,
		OrderId:   orderID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (os *OrdersService) GetOrders(accountID string) ([]*pb.OrderState, error) {
	ctx, cancel := CreateRequestContext(os.config)

	defer cancel()

	res, err := os.client.GetOrders(ctx, &pb.GetOrdersRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}

	return res.Orders, nil
}
