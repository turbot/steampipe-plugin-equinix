package equinix

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type equinixConfig struct {
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
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
