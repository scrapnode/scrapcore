package transport

import "context"

type Transport interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Run(ctx context.Context) error
}

type Configs struct {
	ListenAddress string `json:"listen_address"`
}

type H map[string]interface{}
