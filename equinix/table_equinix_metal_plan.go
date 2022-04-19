package equinix

import (
	"context"

	metal "github.com/packethost/packngo"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableEquinixMetalPlan(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "equinix_metal_plan",
		Description: "Equinix Metal plans.",
		List: &plugin.ListConfig{
			Hydrate: listPlan,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the plan."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "Slug of the plan."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the plan."},
			// Other columns
			// NOTE: returns IDs only
			{Name: "available_in", Type: proto.ColumnType_JSON, Transform: transform.FromField("AvailableIn").Transform(facilitiesToIDs), Description: "Facility IDs where the plan is available."},
			// NOTE: returns IDs only
			{Name: "available_in_metros", Type: proto.ColumnType_JSON, Transform: transform.FromField("AvailableInMetros").Transform(metrosToIDs), Description: "Metro IDs where the plan is available."},
			{Name: "class", Type: proto.ColumnType_STRING, Description: "Class of the plan."},
			{Name: "deployment_types", Type: proto.ColumnType_JSON, Description: "Deployment types for the plan."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the plan."},
			{Name: "href", Type: proto.ColumnType_STRING, Description: "URL of the plan."},
			{Name: "legacy", Type: proto.ColumnType_BOOL, Description: "True if this is a legacy plan."},
			{Name: "line", Type: proto.ColumnType_STRING, Description: "Line of the plan."},
			{Name: "pricing", Type: proto.ColumnType_JSON, Description: "Pricing of the plan."},
			{Name: "specs", Type: proto.ColumnType_JSON, Description: "Specifications of the plan."},
		},
	}
}

func listPlan(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_plan.listPlan", "connection_error", err)
		return nil, err
	}
	maxItems := 1000
	opts := &metal.ListOptions{
		Page:    1,
		PerPage: maxItems,
	}
	for {
		items, resp, err := conn.Plans.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("equinix_metal_plan.listPlan", "query_error", err, "opts", opts, "resp", resp)
			return nil, err
		}
		for _, i := range items {
			d.StreamListItem(ctx, i)
		}
		// ugh ... the Go SDK offers no way to check if we've reached the end, so we use this ugly hack
		if len(items) < maxItems {
			break
		}
		opts.Page++
	}
	return nil, nil
}
