package equinix

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type equinixConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &equinixConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) equinixConfig {
	if connection == nil || connection.GetConfig() == nil {
		return equinixConfig{}
	}
	config, _ := connection.GetConfig().(equinixConfig)
	return config
}
