package pkg

import (
	pb "github.com/ssummers02/invest-api-go-sdk/pkg/investapi"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

type SandboxInterface interface {
	// OpenSandboxAccount The method of registering an account in the sandbox.
	OpenSandboxAccount() (string, error)
	// GetSandboxAccounts The method of getting accounts in the sandbox.
	GetSandboxAccounts() ([]*pb.Account, error)
	// CloseSandboxAccount The method of closing an account in the sandbox.
	CloseSandboxAccount(accountID string) error
	// PostSandboxOrder The method of placing a trade order in the sandbox.
	PostSandboxOrder(order *pb.PostOrderRequest) (*pb.PostOrderResponse, error)
	// GetSandboxOrders Method for getting a list of active applications for an account in the sandbox.
	GetSandboxOrders(accountID string) ([]*pb.OrderState, error)
	// CancelSandboxOrder Method for getting a list of active orders for an account in the sandbox.
	CancelSandboxOrder(accountID string, orderID string) (*timestamp.Timestamp, error)
	// GetSandboxOrderState The method of obtaining the order status in the sandbox.
	GetSandboxOrderState(accountID string, orderID string) (*pb.OrderState, error)
	// GetSandboxPositions The method of obtaining positions on the virtual sandbox account.
	GetSandboxPositions(accountID string) (*pb.PositionsResponse, error)
	// GetSandboxOperations The method of receiving operations in the sandbox by account number.
	GetSandboxOperations(filter *pb.OperationsRequest) ([]*pb.Operation, error)
	// GetSandboxPortfolio The method of getting a portfolio in the sandbox.
	GetSandboxPortfolio(accountID string) (*pb.PortfolioResponse, error)
	// SandboxPayIn The method of depositing funds in the sandbox.
	SandboxPayIn(accountID string, amount *pb.MoneyValue) (*pb.MoneyValue, error)
}

type SandboxService struct {
	client pb.SandboxServiceClient
	config Config
}

func NewSandboxService(conn *grpc.ClientConn, config Config) *SandboxService {
	client := pb.NewSandboxServiceClient(conn)

	return &SandboxService{client: client, config: config}
}

func (ss *SandboxService) OpenSandboxAccount() (string, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.OpenSandboxAccount(ctx, &pb.OpenSandboxAccountRequest{})
	if err != nil {
		return "", err
	}

	return res.AccountId, nil
}

func (ss *SandboxService) GetSandboxAccounts() ([]*pb.Account, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.GetSandboxAccounts(ctx, &pb.GetAccountsRequest{})
	if err != nil {
		return nil, err
	}

	return res.Accounts, nil
}

func (ss *SandboxService) CloseSandboxAccount(accountID string) error {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	_, err := ss.client.CloseSandboxAccount(ctx, &pb.CloseSandboxAccountRequest{
		AccountId: accountID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (ss *SandboxService) PostSandboxOrder(order *pb.PostOrderRequest) (*pb.PostOrderResponse, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.PostSandboxOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ss *SandboxService) GetSandboxOrders(accountID string) ([]*pb.OrderState, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.GetSandboxOrders(ctx, &pb.GetOrdersRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}

	return res.Orders, nil
}

func (ss *SandboxService) CancelSandboxOrder(accountID string, orderID string) (*timestamp.Timestamp, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.CancelSandboxOrder(ctx, &pb.CancelOrderRequest{
		AccountId: accountID,
		OrderId:   orderID,
	})
	if err != nil {
		return nil, err
	}

	return res.Time, nil
}

func (ss *SandboxService) GetSandboxOrderState(accountID string, orderID string) (*pb.OrderState, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.GetSandboxOrderState(ctx, &pb.GetOrderStateRequest{
		AccountId: accountID,
		OrderId:   orderID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ss *SandboxService) GetSandboxPositions(accountID string) (*pb.PositionsResponse, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.GetSandboxPositions(ctx, &pb.PositionsRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ss *SandboxService) GetSandboxOperations(filter *pb.OperationsRequest) ([]*pb.Operation, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.GetSandboxOperations(ctx, filter)
	if err != nil {
		return nil, err
	}

	return res.Operations, nil
}

func (ss *SandboxService) GetSandboxPortfolio(accountID string) (*pb.PortfolioResponse, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.GetSandboxPortfolio(ctx, &pb.PortfolioRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ss *SandboxService) SandboxPayIn(accountID string, amount *pb.MoneyValue) (*pb.MoneyValue, error) {
	ctx, cancel := CreateRequestContext(ss.config)
	defer cancel()

	res, err := ss.client.SandboxPayIn(ctx, &pb.SandboxPayInRequest{
		AccountId: accountID,
		Amount:    amount,
	})
	if err != nil {
		return nil, err
	}

	return res.Balance, nil
}
