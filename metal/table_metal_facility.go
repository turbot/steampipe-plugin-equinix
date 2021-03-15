package metal

import (
	"context"

	metal "github.com/packethost/packngo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableMetalFacility(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "metal_facility",
		Description: "Equinix Metal facilities.",
		List: &plugin.ListConfig{
			Hydrate: listFacility,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the facility."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the facility."},
			// Other columns
			// TODO: address is always empty?
			{Name: "address", Type: proto.ColumnType_JSON, Description: "Address of the facility."},
			{Name: "code", Type: proto.ColumnType_STRING, Description: "Code of the facility."},
			{Name: "features", Type: proto.ColumnType_JSON, Description: "Features of the facility."},
			// TODO: href is always empty?
			{Name: "href", Type: proto.ColumnType_STRING, Description: "URL of the facility."},
			// Resource columns
			// TODO: href is always empty, so this breaks too?
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromField("URL").Transform(transform.EnsureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]string{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listFacility(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("metal_facility.listFacility", "connection_error", err)
		return nil, err
	}
	maxItems := 1000
	opts := &metal.ListOptions{
		Page:    1,
		PerPage: maxItems,
	}
	for {
		items, resp, err := conn.Facilities.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("metal_facility.listFacility", "query_error", err, "opts", opts, "resp", resp)
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
