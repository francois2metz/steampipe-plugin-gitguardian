package gitguardian

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/schema"
)

type gitguardianConfig struct {
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &gitguardianConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) gitguardianConfig {
	if connection == nil || connection.GetConfig() == nil {
		return gitguardianConfig{}
	}
	config, _ := connection.GetConfig().(gitguardianConfig)
	return config
}
