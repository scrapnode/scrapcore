package noop

import (
	"context"
	"github.com/scrapnode/scrapcore/xmonitor/configs"
)

func New(ctx context.Context, cfg *configs.Configs) (*Monitor, error) {
	return &Monitor{propergator: &Propergator{}}, nil
}
