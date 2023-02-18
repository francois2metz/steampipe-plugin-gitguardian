package gitguardian

import (
	"context"
	"errors"
	"os"

	"github.com/Gaardsholt/go-gitguardian/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (client.ClientOption, error) {
	// get gitguardian client from cache
	cacheKey := "gitguardian"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(client.ClientOption), nil
	}

	token := os.Getenv("GITGUARDIAN_TOKEN")

	gitguardianConfig := GetConfig(d.Connection)

	if gitguardianConfig.Token != nil {
		token = *gitguardianConfig.Token
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file or set the GITGUARDIAN_TOKEN environment variable and then restart Steampipe")
	}

	client := client.WithApiKey(token)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}
