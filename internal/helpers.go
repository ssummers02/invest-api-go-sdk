package internal

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpcMetadata "google.golang.org/grpc/metadata"
)

const (
	ApiURL = "invest-public-api.tinkoff.ru:443"

	DefaultRequestTimeout = 30 * time.Second
)

type TradeBotConfig struct {
	IsSandbox bool   `default:"true" split_words:"true"`
	Token     string `required:"true"`
	AccountID string `split_words:"true"` // required in non-sandbox mode
}

func CreateClientConn() (*grpc.ClientConn, error) {
	tlsConfig := tls.Config{}

	return grpc.Dial(ApiURL, grpc.WithTransportCredentials(credentials.NewTLS(&tlsConfig)))
}

// CreateRequestContext returns context for API calls with timeout and auth headers attached.
func CreateRequestContext(cfg TradeBotConfig) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultRequestTimeout)

	authHeader := fmt.Sprintf("Bearer %s", cfg.Token)
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", authHeader)
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "x-tracking-id", uuid.New().String())
	//ctx = grpcMetadata.AppendToOutgoingContext(ctx, "x-app-name", AppName)

	return ctx, cancel
}
