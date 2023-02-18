package gitguardian

import (
	"context"

	"github.com/Gaardsholt/go-gitguardian/sources"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitguardianSource(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "gitguardian_source",
		Description: "Retrieve details on sources monitored by GitGuardian.",
		List: &plugin.ListConfig{
			Hydrate: listSource,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSource,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "Unique identifier of the source.",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Description: "URL of the source.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Source type.",
			},
			{
				Name:        "full_name",
				Type:        proto.ColumnType_STRING,
				Description: "The full name of the source.",
			},
			{
				Name:        "visibility",
				Type:        proto.ColumnType_STRING,
				Description: "The visibility of the source (public or private).",
			},
			{
				Name:        "health",
				Type:        proto.ColumnType_STRING,
				Description: "Source health: safe, unknown or at_risk.",
			},
			{
				Name:        "open_incidents_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of open secret incidents with at least one occurrence on this source.",
				Transform:   transform.FromField("OpenIncidentsCount"),
			},
			{
				Name:        "closed_incidents_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of closed secret incidents with at least one occurrence on this source.",
				Transform:   transform.FromField("ClosedIncidentsCount"),
			},
			{
				Name:        "last_scan_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Creation date of this historical scan.",
				Transform:   transform.FromField("LastScan.Date"),
			},
			{
				Name:        "last_scan_status",
				Type:        proto.ColumnType_STRING,
				Description: "Status of the last scan: pending, running, canceled, failed, too_large, timeout, finished.",
				Transform:   transform.FromField("LastScan.Status"),
			},
		},
	}
}

func listSource(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_source.listSource", "connection_error", err)
		return nil, err
	}
	c, err := sources.NewClient(client)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_source.listSource", "connection_error", err)
		return nil, err
	}
	perPage := 100
	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < int64(perPage) {
		perPage = int(*d.QueryContext.Limit)
	}

	opts := sources.ListOptions{
		PerPage: &perPage,
	}
	for {
		result, pagination, err := c.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("gitguardian_source.listSource", err)
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

func getSource(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_source.getSource", "connection_error", err)
		return nil, err
	}
	c, err := sources.NewClient(client)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_source.getSource", "connection_error", err)
		return nil, err
	}

	id := d.EqualsQuals["id"].GetInt64Value()

	result, err := c.Get(int(id))
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_sources.getSources", err)
		return nil, err
	}
	return result.Result, nil
}
