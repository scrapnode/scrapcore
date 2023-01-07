package monitor

import "context"

type Configs struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Version   string `json:"version"`

	Tracer *TracerConfigs
}

type TracerConfigs struct {
	Endpoint string  `json:"endpoint"`
	Ratio    float64 `json:"ratio"`
}

type Monitor interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
}
