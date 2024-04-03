package grpc

import (
	"github.com/jbakhtin/goph-keeper-client/internal/infrastructure/persistance/grpc/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config interface {
	GetGRPCServerAddress() string
	GetGRPCServerNetwork() string
}

type Client struct {
	cfg  Config
	conn *grpc.ClientConn
}

func NewClient(cfg Config) (*Client, error) {
	jwtCredentials, err := credentials.NewJWTCredentials()
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(cfg.GetGRPCServerAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(jwtCredentials),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		cfg:  cfg,
		conn: conn,
	}, nil
}

func (c *Client) GetConn() *grpc.ClientConn {
	return c.conn
}
