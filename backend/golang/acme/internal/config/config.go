package config

import (
	"encoding/json"
	"go-learn/acme/internal/logging"
	"os"
)

const DefaultEnvVar = "ACME_CONFIG"

var App *Config

type Config struct {
	DSN                 string
	Address             string
	BasePrice           float64
	ExchangeRateBaseURL string
	ExchangeAPIKEY      string
}

func init() {
	filename, found := os.LookupEnv(DefaultEnvVar)
	if !found {
		logging.L.Error("failed to locate file specified by %s", DefaultEnvVar)
		return
	}

	_ = load(filename)
}

func load(filename string) error {
	App := &Config{}
	bytes, err := os.ReadFile(filename)
	if err != nil {
		logging.L.Error("failed to parse config file. err: %s", err)
		return err
	}

	err = json.Unmarshal(bytes, App)
	if err != nil {
		logging.L.Error("failed to parse config file. err: %s", err)
		return err
	}

	return nil
}
