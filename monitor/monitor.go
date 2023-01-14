package monitor

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
		Metrics: &MetricsConfigs{
			Endpoint: cfg.Metrics.Endpoint,
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
