package internal

import (
	pb "invest-api-go-sdk/internal/investapi"

	"google.golang.org/grpc"
)

type OrderStreamInterface interface {
	// Recv listens for incoming messages and block until first one is received.
	Recv() (*pb.TradesStreamResponse, error)
}

type OrdersStreamService struct {
	client pb.OrdersStreamServiceClient
	stream pb.OrdersStreamService_TradesStreamClient
}

func NewOrdersStreamService(conn *grpc.ClientConn, config TradeBotConfig) (*OrdersStreamService, error) {
	client := pb.NewOrdersStreamServiceClient(conn)
	ctx := CreateStreamContext(config)

	stream, err := client.TradesStream(ctx, &pb.TradesStreamRequest{Accounts: config.AccountID})
	if err != nil {
		return nil, err
	}

	return &OrdersStreamService{
		client: client,
		stream: stream,
	}, nil
}

func (os *OrdersStreamService) Recv() (*pb.TradesStreamResponse, error) {
	return os.stream.Recv()
}
