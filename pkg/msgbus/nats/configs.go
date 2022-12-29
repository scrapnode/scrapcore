package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/scrapnode/scrapcore/pkg/msgbus/configs"
	"time"
)

func ParseJetStreamConfigs(ctx context.Context, configs *configs.Configs) *nats.StreamConfig {
	// @TODO: allow configure JetStream configs from context

	// default
	cfg := &nats.StreamConfig{
		Name:     NewStreamName(configs),
		Replicas: 3,
		// 8kb/msg -> 4Gb
		MaxMsgs:  524288,
		MaxBytes: 8388608,
		MaxAge:   time.Hour,
		Subjects: []string{NewSubject(configs, nil)},
	}

	return cfg
}
