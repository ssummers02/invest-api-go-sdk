package pkg

import (
	pb "github.com/ssummers02/invest-api-go-sdk/pkg/investapi"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

type OperationsInterface interface {
	// GetOperations Method for getting a list of account transactions.
	GetOperations(accountID string, from, to *timestamp.Timestamp, state pb.OperationState, figi string) ([]*pb.Operation, error)
	// GetPortfolio The method of obtaining a portfolio by account.
	GetPortfolio(accountID string) (*pb.PortfolioResponse, error)
	// GetPositions Method for getting a list of account positions.
	GetPositions(accountID string) (*pb.PositionsResponse, error)
	// GetWithdrawLimits The method of obtaining the available balance for withdrawal of funds.
	GetWithdrawLimits(accountID string) (*pb.WithdrawLimitsResponse, error)
}

type OperationsService struct {
	client pb.OperationsServiceClient
	config Config
}

func NewOperationsService(conn *grpc.ClientConn, config Config) *OperationsService {
	client := pb.NewOperationsServiceClient(conn)

	return &OperationsService{
		client: client,
		config: config,
	}
}

func (os *OperationsService) GetOperations(accountID string, from, to *timestamp.Timestamp, state pb.OperationState, figi string) ([]*pb.Operation, error) {
	ctx, cancel := CreateRequestContext(os.config)
	defer cancel()

	res, err := os.client.GetOperations(ctx, &pb.OperationsRequest{
		AccountId: accountID,
		From:      from,
		To:        to,
		State:     state,
		Figi:      figi,
	})
	if err != nil {
		return nil, err
	}

	return res.Operations, nil
}

func (os *OperationsService) GetPortfolio(accountID string) (*pb.PortfolioResponse, error) {
	ctx, cancel := CreateRequestContext(os.config)
	defer cancel()

	res, err := os.client.GetPortfolio(ctx, &pb.PortfolioRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (os *OperationsService) GetPositions(accountID string) (*pb.PositionsResponse, error) {
	ctx, cancel := CreateRequestContext(os.config)
	defer cancel()

	res, err := os.client.GetPositions(ctx, &pb.PositionsRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (os *OperationsService) GetWithdrawLimits(accountID string) (*pb.WithdrawLimitsResponse, error) {
	ctx, cancel := CreateRequestContext(os.config)
	defer cancel()

	res, err := os.client.GetWithdrawLimits(ctx, &pb.WithdrawLimitsRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
