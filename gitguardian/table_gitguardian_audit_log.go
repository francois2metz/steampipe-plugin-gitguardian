package gitguardian

import (
	"context"

	"github.com/Gaardsholt/go-gitguardian/auditlogs"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableGitguardianAuditLog(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "gitguardian_audit_log",
		Description: "List audit logs.",
		List: &plugin.ListConfig{
			Hydrate:    listAuditLog,
			KeyColumns: []*plugin.KeyColumn{},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "Unique identifier of the log.",
			},
			{
				Name:        "date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Date the event occurred.",
			},
			{
				Name:        "member_email",
				Type:        proto.ColumnType_STRING,
				Description: "Email of the member at the time he/she did the event.",
			},
			{
				Name:        "member_name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the member at the time he/she did the event.",
			},
			{
				Name:        "member_id",
				Type:        proto.ColumnType_INT,
				Description: "ID of the member that did the event. Can be null if the member has been deleted since then.",
			},
			{
				Name:        "api_token_id",
				Type:        proto.ColumnType_INT,
				Description: "ID of the API token associated to the event if it was done through the API.",
			},
			{
				Name:        "ip_address",
				Type:        proto.ColumnType_INET,
				Description: "The ip address.",
			},
			{
				Name:        "action_type",
				Type:        proto.ColumnType_STRING,
				Description: "Type of the action: READ, CREATE, UPDATE, DELETE, OTHER.",
			},
			{
				Name:        "event_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the event.",
			},
			{
				Name:        "data",
				Type:        proto.ColumnType_JSON,
				Description: "Additional data associated to the event.",
			},
		},
	}
}

func listAuditLog(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_audit_log.listAuditLog", "connection_error", err)
		return nil, err
	}
	c, err := auditlogs.NewClient(client)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_audit_log.listAuditLog", "connection_error", err)
		return nil, err
	}
	perPage := 100

	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < int64(perPage) {
		perPage = int(*d.QueryContext.Limit)
	}

	opts := auditlogs.AuditLogsListOptions{
		PerPage: &perPage,
	}
	for {
		result, pagination, err := c.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("gitguardian_audit_log.listAuditLog", err)
			return nil, err
		}
		for _, r := range result.Result {
			d.StreamListItem(ctx, r)
		}
		if pagination.NextCursor == "" {
			break
		}
		opts.Cursor = pagination.NextCursor
		if d.QueryStatus.RowsRemaining(ctx) <= 0 {
			break
		}
	}
	return nil, nil
}
