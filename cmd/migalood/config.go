package main

import serverconfig "github.com/cosmos/cosmos-sdk/server/config"

// initAppConfig generates contents for `app.toml`.
// It takes the default template and config, and appends custom parameters.

func initAppConfig() (string, interface{}) {
	template := serverconfig.DefaultConfigTemplate

	cfg := serverconfig.DefaultConfig()
	cfg.MinGasPrices = "0uwhale"

	return template, cfg
}
