package gitguardian

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-gitguardian",
		DefaultTransform: transform.FromGo().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"gitguardian_audit_log":       tableGitguardianAuditLog(ctx),
			"gitguardian_member":          tableGitguardianMember(ctx),
			"gitguardian_secret_incident": tableGitguardianSecretIncident(ctx),
			"gitguardian_source":          tableGitguardianSource(ctx),
		},
	}
	return p
}
