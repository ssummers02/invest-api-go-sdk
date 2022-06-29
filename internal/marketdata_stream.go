package internal

import (
	pb "invest-api-go-sdk/internal/investapi"

	"google.golang.org/grpc"
)

type MarketDataStreamInterface interface {
	// Recv listens for incoming messages and block until first one is received.
	Recv() (*pb.MarketDataResponse, error)
	// Send puts pb.MarketDataRequest into a stream.
	Send(request *pb.MarketDataRequest) error
}

type MarketDataStream struct {
	client pb.MarketDataStreamServiceClient
	stream pb.MarketDataStreamService_MarketDataStreamClient
}

func NewMarketDataStream(conn *grpc.ClientConn, config TradeBotConfig) (*MarketDataStream, error) {
	client := pb.NewMarketDataStreamServiceClient(conn)
	ctx := CreateStreamContext(config)

	stream, err := client.MarketDataStream(ctx)
	if err != nil {
		return nil, err
	}

	return &MarketDataStream{client: client, stream: stream}, nil
}

func (ms *MarketDataStream) Recv() (*pb.MarketDataResponse, error) {
	return ms.stream.Recv()
}

func (ms *MarketDataStream) Send(request *pb.MarketDataRequest) error {
	return ms.stream.Send(request)
}
