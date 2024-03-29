package main

import (
	"github.com/francois2metz/steampipe-plugin-gitguardian/gitguardian"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: gitguardian.Plugin})
}
