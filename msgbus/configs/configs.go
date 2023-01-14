package configs

type Configs struct {
	Provider string `json:"provider"`

	Uri    string `json:"uri"`
	Region string `json:"region"`
	Name   string `json:"name"`

	MaxRetry int `json:"max_retry"`
}
