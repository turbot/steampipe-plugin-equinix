package main

import (
	"github.com/turbot/steampipe-plugin-equinix/equinix"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: equinix.Plugin})
}
