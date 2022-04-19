package equinix

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableEquinixMetalOperatingSystem(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "equinix_metal_operating_system",
		Description: "Equinix Metal operating systems.",
		List: &plugin.ListConfig{
			Hydrate: listOperatingSystem,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the operating system."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "Slug of the operating system."},
			{Name: "distro", Type: proto.ColumnType_STRING, Description: "Distro of the operating system."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "Version of the operating system."},
			// Other columns
			{Name: "provisionable_on", Type: proto.ColumnType_JSON, Description: "Device types that the operating system can be provisioned on."},
		},
	}
}

func listOperatingSystem(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_operating_system.listOperatingSystem", "connection_error", err)
		return nil, err
	}
	items, resp, err := conn.OperatingSystems.List()
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_operating_system.listOperatingSystem", "query_error", err, "resp", resp)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
