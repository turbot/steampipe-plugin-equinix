package equinix

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type equinixConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &equinixConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) equinixConfig {
	if connection == nil || connection.Config == nil {
		return equinixConfig{}
	}
	config, _ := connection.Config.(equinixConfig)
	return config
}
