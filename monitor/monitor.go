package monitor

import "context"

type Configs struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Version   string `json:"version"`

	Tracer  *TracerConfigs  `json:"tracer"`
	Metrics *MetricsConfigs `json:"metrics"`
}

func (cfg *Configs) Clone() *Configs {
	return &Configs{
		Namespace: cfg.Namespace,
		Name:      cfg.Name,
		Version:   cfg.Version,
		Tracer: &TracerConfigs{
			Endpoint: cfg.Tracer.Endpoint,
			Ratio:    cfg.Tracer.Ratio,
		},
	}
}

type TracerConfigs struct {
	Endpoint string  `json:"endpoint"`
	Ratio    float64 `json:"ratio"`
}

type MetricsConfigs struct {
	Endpoint string `json:"endpoint"`
}

type Monitor interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
}
