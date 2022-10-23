package gitguardian

import (
	"context"

	"github.com/Gaardsholt/go-gitguardian/members"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableGitguardianMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "gitguardian_member",
		Description: "List members of the workspace.",
		List: &plugin.ListConfig{
			Hydrate: listMember,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "role", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getMember,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "Unique identifier of the member.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the member.",
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "Email of the member.",
			},
			{
				Name:        "role",
				Type:        proto.ColumnType_STRING,
				Description: "Role of the member (owner, manager, member, viewer, restricted).",
			},
		},
	}
}

func listMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_member.listMember", "connection_error", err)
		return nil, err
	}
	c, err := members.NewClient(client)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_member.listMember", "connection_error", err)
		return nil, err
	}
	perPage := 100
	quals := d.KeyColumnQuals
	role := quals["role"].GetStringValue()

	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < int64(perPage) {
		perPage = int(*d.QueryContext.Limit)
	}

	opts := members.ListOptions{
		PerPage: &perPage,
		Role:    members.MembersListRole(role),
	}

	for {
		result, pagination, err := c.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("gitguardian_member.listMember", err)
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

func getMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_member.getMember", "connection_error", err)
		return nil, err
	}
	c, err := members.NewClient(client)
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_member.getMember", "connection_error", err)
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetInt64Value()

	result, err := c.Get(int(id))
	if err != nil {
		plugin.Logger(ctx).Error("gitguardian_member.getMember", err)
		return nil, err
	}
	return result.Result, nil
}
