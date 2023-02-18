package gitguardian

import (
	"context"

	"github.com/Gaardsholt/go-gitguardian/auditlogs"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGitguardianAuditLog(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "gitguardian_audit_log",
		Description: "List audit logs.",
		List: &plugin.ListConfig{
			Hydrate: listAuditLog,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "date", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
				{Name: "event_name", Require: plugin.Optional},
				{Name: "member_id", Require: plugin.Optional},
				{Name: "member_name", Require: plugin.Optional},
				{Name: "member_email", Require: plugin.Optional},
			},
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
	quals := d.EqualsQuals
	eventName := quals["event_name"].GetStringValue()
	memberId := quals["member_id"]
	memberName := quals["member_name"].GetStringValue()
	memberEmail := quals["member_email"].GetStringValue()
	date := d.Quals["date"]

	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < int64(perPage) {
		perPage = int(*d.QueryContext.Limit)
	}

	opts := auditlogs.AuditLogsListOptions{
		PerPage:     &perPage,
		EventName:   eventName,
		MemberName:  memberName,
		MemberEmail: memberEmail,
	}
	if memberId != nil {
		memberId := int(memberId.GetInt64Value())
		opts.MemberId = &memberId
	}

	if date != nil {
		for _, q := range date.Quals {
			timestamp := q.Value.GetTimestampValue().AsTime()
			switch q.Operator {
			case "=":
				opts.DateBefore = &timestamp
				opts.DateAfter = &timestamp
			case ">=", ">":
				opts.DateAfter = &timestamp
			case "<", "<=":
				opts.DateBefore = &timestamp
			}
		}
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
		if d.RowsRemaining(ctx) <= 0 {
			break
		}
	}
	return nil, nil
}
