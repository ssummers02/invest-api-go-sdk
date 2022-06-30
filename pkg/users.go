package pkg

import (
	pb "github.com/ssummers02/invest-api-go-sdk/pkg/investapi"

	"google.golang.org/grpc"
)

type UsersInterface interface {
	// GetAccounts The method of receiving user accounts.
	GetAccounts() ([]*pb.Account, error)
	// GetMarginAttributes Calculation of margin indicators on the account.
	GetMarginAttributes(accountID string) (*pb.GetMarginAttributesResponse, error)
	// GetUserTariff Request for the user's tariff.
	GetUserTariff() (*pb.GetUserTariffResponse, error)
	// GetInfo The method of obtaining information about the user.
	GetInfo() (*pb.GetInfoResponse, error)
}

type UsersService struct {
	client pb.UsersServiceClient
	config Config
}

func NewUsersService(conn *grpc.ClientConn, config Config) *UsersService {
	client := pb.NewUsersServiceClient(conn)

	return &UsersService{
		client: client,
		config: config,
	}
}

func (us *UsersService) GetAccounts() ([]*pb.Account, error) {
	ctx, cancel := CreateRequestContext(us.config)
	defer cancel()

	response, err := us.client.GetAccounts(ctx, &pb.GetAccountsRequest{})

	if err != nil {
		return nil, err
	}

	return response.Accounts, nil
}

func (us *UsersService) GetMarginAttributes(accountID string) (*pb.GetMarginAttributesResponse, error) {
	ctx, cancel := CreateRequestContext(us.config)
	defer cancel()

	response, err := us.client.GetMarginAttributes(ctx, &pb.GetMarginAttributesRequest{AccountId: accountID})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (us *UsersService) GetUserTariff() (*pb.GetUserTariffResponse, error) {
	ctx, cancel := CreateRequestContext(us.config)
	defer cancel()

	res, err := us.client.GetUserTariff(ctx, &pb.GetUserTariffRequest{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *UsersService) GetInfo() (*pb.GetInfoResponse, error) {
	ctx, cancel := CreateRequestContext(us.config)
	defer cancel()

	res, err := us.client.GetInfo(ctx, &pb.GetInfoRequest{})
	if err != nil {
		return nil, err
	}

	return res, nil
}
