package xconfig

import (
	"github.com/spf13/viper"
	"os"
)

var ENV_DEV = "development"
var ENV_PROD = "production"

type Configs struct {
	Env     string `json:"env" mapstructure:"SCRAP_ENV"`
	Version string `json:"version" mapstructure:"SCRAP_VERSION"`
}

func (cfg *Configs) Debug() bool {
	return cfg.Env == ENV_DEV
}

func (cfg *Configs) Unmarshal(provider *viper.Viper) error {
	provider.SetDefault("SCRAP_ENV", ENV_PROD)
	provider.Set("SCRAP_VERSION", version())

	return provider.Unmarshal(cfg)
}

func New(dirs ...string) (*viper.Viper, error) {
	provider := viper.New()
	provider.AutomaticEnv()
	provider.SetConfigName("configs")
	provider.SetConfigType("props")

	for _, dir := range dirs {
		provider.AddConfigPath(dir)
		if err := provider.MergeInConfig(); err != nil {
			// ignore not found files, otherwise return error
			if _, notfound := err.(viper.ConfigFileNotFoundError); !notfound {
				return nil, err
			}
		}
	}
	return provider, nil
}

func version() string {
	if body, err := os.ReadFile(".version"); err == nil {
		return string(body)
	}

	return "22.2.22"
}
