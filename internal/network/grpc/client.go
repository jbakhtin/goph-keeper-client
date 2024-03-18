package grpc

type Config interface {
	GetGRPCServerAddress() string
	GetGRPCServerNetwork() string
}

type Client struct {
	cfg Config
}
