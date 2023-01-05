package xsender

import (
	"context"
	"errors"
	"net/url"
)

type Configs struct {
	TimeoutInSeconds int `json:"timeout_in_seconds"`
	RetryMax         int `json:"retry_max"`
}

type Send func(req *Request) (*Response, error)

func New(ctx context.Context, cfg *Configs) Send {
	factory := map[string]Send{
		"http":  NewHttp(ctx, cfg),
		"https": NewHttp(ctx, cfg),
	}

	return func(req *Request) (*Response, error) {
		uri, err := url.Parse(req.Uri)
		if err != nil {
			return nil, err
		}

		send, ok := factory[uri.Scheme]
		if !ok {
			return nil, errors.New("xsender: unsupported scheme")
		}

		return send(req)
	}
}
