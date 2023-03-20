package equinix

import (
	"context"

	metal "github.com/packethost/packngo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableEquinixMetalMetro(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "equinix_metal_metro",
		Description: "Equinix Metal metros.",
		List: &plugin.ListConfig{
			Hydrate: listMetro,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the metro."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the metro."},
			// Other columns
			{Name: "code", Type: proto.ColumnType_STRING, Description: "Code of the metro."},
			{Name: "country", Type: proto.ColumnType_STRING, Description: "Country of the metro."},
			{Name: "href", Type: proto.ColumnType_STRING, Description: "URL of the metro."},
		},
	}
}

func listMetro(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_metro.listMetro", "connection_error", err)
		return nil, err
	}
	maxItems := 1000
	opts := &metal.ListOptions{
		Page:    1,
		PerPage: maxItems,
	}
	for {
		items, resp, err := conn.Metros.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("equinix_metal_metro.listMetro", "query_error", err, "opts", opts, "resp", resp)
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
