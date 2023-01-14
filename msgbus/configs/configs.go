package configs

type Configs struct {
	Dsn    string `json:"dsn"`
	Region string `json:"region"`
	Name   string `json:"name"`

	MaxRetry int `json:"max_retry"`
}
