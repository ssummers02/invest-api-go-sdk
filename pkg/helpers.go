package pkg

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
	APIURL                = "invest-public-api.tinkoff.ru:443"
	DefaultRequestTimeout = 30 * time.Second
	AppName               = "ssummers02/invest-api-go-sdk"
)

type Config struct {
	Token     string   `required:"true"`
	AccountID []string `split_words:"true"` // required in non-sandbox mode
}

func CreateClientConn() (*grpc.ClientConn, error) {
	tlsConfig := tls.Config{MinVersion: tls.VersionTLS12}

	return grpc.Dial(APIURL, grpc.WithTransportCredentials(credentials.NewTLS(&tlsConfig)))
}

// CreateRequestContext returns context for API calls with timeout and auth headers attached.
func CreateRequestContext(cfg Config) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultRequestTimeout)

	authHeader := fmt.Sprintf("Bearer %s", cfg.Token)
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", authHeader)
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "x-tracking-id", uuid.New().String())
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "x-app-name", AppName)

	return ctx, cancel
}

// CreateStreamContext returns context for streams with auth headers attached.
func CreateStreamContext(cfg Config) context.Context {
	ctx := context.TODO()

	authHeader := fmt.Sprintf("Bearer %s", cfg.Token)
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", authHeader)
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "x-tracking-id", uuid.New().String())
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "x-app-name", AppName)

	return ctx
}
