package main

import (
	"github.com/equinix/steampipe-plugin-metal/metal"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: metal.Plugin})
}
