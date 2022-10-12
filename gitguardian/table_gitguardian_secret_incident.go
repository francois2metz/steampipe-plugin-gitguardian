package gitguardian

import (
	"context"

	"github.com/Gaardsholt/go-gitguardian/incidents"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableGitguardianSecretIncident(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "gitguardian_secret_incident",
		Description: "List secret incidents detected by the GitGuardian dashboard.",
		List: &plugin.ListConfig{
			Hydrate: listSecretIncident,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "assignee_email", Require: plugin.Optional},
				{Name: "date", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
				{Name: "severity", Require: plugin.Optional},
				{Name: "status", Require: plugin.Optional},
				{Name: "validity", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSecretIncident,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "Unique identifier of the incident.",
			},
			{
				Name:        "date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Last trigger date.",
			},
			{
				Name:        "secret_hash",
				Type:        proto.ColumnType_STRING,
				Description: "Unique hash.",
			},
			{
				Name:        "gitguardian_url",
				Type:        proto.ColumnType_STRING,
				Description: "The URL to gitguardian.",
			},
			{
				Name:        "regression",
				Type:        proto.ColumnType_BOOL,
				Description: "True if it's a regression.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "Status of the incident: IGNORED, TRIGGERED, ASSIGNED, RESOLVED).",
			},
			{
				Name:        "assignee_email",
				Type:        proto.ColumnType_STRING,
				Description: "Assignee email.",
			},
			{
				Name:        "occurrences_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of occurrences.",
				Transform:   transform.FromField("OccurrencesCount"),
			},
			{
				Name:        "ignore_reason",
				Type:        proto.ColumnType_STRING,
				Description: "The reason of the ignore status: test_credential, false_positive, low_risk",
			},
			{
				Name:        "ignored_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date where it has been ignored.",
			},
			{
				Name:        "secret_revoked",
				Type:        proto.ColumnType_BOOL,
				Description: "True if the secret has been revoked.",
			},
			{
				Name:        "severity",
				Type:        proto.ColumnType_STRING,
				Description: "Severity of the incident: critical, high, medium, low, info, unknown.",
			},
			{
				Name:        "validity",
				Type:        proto.ColumnType_STRING,
				Description: "The validity state: valid, invalid, failed_to_check, no_checker, unknown",
			},
			{
				Name:        "resolved_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date where it has been resolved.",
			},
			{
				Name:        "share_url",
				Type:        proto.ColumnType_STRING,
				Description: "The public URL of the incident (if any).",
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: "Can be FROM_HISTORICAL_SCAN, IGNORED_IN_CHECK_RUN, PUBLIC, REGRESSION, SENSITIVE_FILE or TEST_FILE",
			},
		},
	}
}

func listSecretIncident(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_secret_incident.listSecretIncident", "connection_error", err)
		return nil, err
	}
	c, err := incidents.NewClient(client)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_secret_incident.listSecretIncident", "connection_error", err)
		return nil, err
	}
	perPage := 100
	quals := d.KeyColumnQuals
	assigneeEmail := quals["assignee_email"].GetStringValue()
	status := incidents.IncidentsListStatus(quals["status"].GetStringValue())
	severity := incidents.IncidentsListSeverity(quals["severity"].GetStringValue())
	validity := incidents.IncidentsListValidity(quals["validity"].GetStringValue())
	date := d.Quals["date"]

	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < int64(perPage) {
		perPage = int(*d.QueryContext.Limit)
	}

	opts := incidents.ListOptions{
		AssigneeEmail: assigneeEmail,
		PerPage:       &perPage,
		Severity:      &severity,
		Status:        &status,
		Validity:      &validity,
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
			plugin.Logger(ctx).Error("gitguardian_secret_incident.listSecretIncident", err)
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

func getSecretIncident(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_secret_incident.getSecretIncident", "connection_error", err)
		return nil, err
	}
	c, err := incidents.NewClient(client)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_secret_incident.getSecretIncident", "connection_error", err)
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetInt64Value()

	result, err := c.Get(int(id), incidents.GetOptions{})
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_secret_incident.getSecretIncident", err)
		return nil, err
	}
	return result.Result, nil
}
