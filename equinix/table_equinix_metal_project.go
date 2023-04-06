package equinix

import (
	"context"

	metal "github.com/packethost/packngo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableEquinixMetalProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "equinix_metal_project",
		Description: "Projects the user has access to.",
		List: &plugin.ListConfig{
			Hydrate: listProject,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			Hydrate:           getProject,
			ShouldIgnoreError: isNotFoundError([]string{"404 Not found"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the project."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the project."},
			// Other columns
			{Name: "backend_transfer_enabled", Type: proto.ColumnType_BOOL, Description: "True if backend transfer is enabled for the project."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the project was created."},
			// Use equinix_metal_devices instead
			// {Name: "devices", Type: proto.ColumnType_JSON, Description: "Devices in the project."},
			{Name: "href", Type: proto.ColumnType_STRING, Description: "URL for the project."},
			//{Name: "members", Type: proto.ColumnType_JSON, Description: "Users in the project."},
			{Name: "member_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Users").Transform(usersToIDs), Description: "IDs of the members of this project."},
			//{Name: "organization", Type: proto.ColumnType_JSON, Description: "Organization for the project."},
			{Name: "organization_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Organization.URL").Transform(hrefToID), Description: "Organization for the project."},
			//{Name: "payment_method", Type: proto.ColumnType_JSON, Description: "Payment method for the project."},
			{Name: "payment_method_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("PaymentMethod.URL").Transform(hrefToID), Description: "Payment method for the project."},
			// Use equinix_metal_ssh_keys instead
			// {Name: "ssh_keys", Type: proto.ColumnType_JSON, Description: "SSH Keys in the project."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the project was updated."},
			// Resource columns
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromField("URL").Transform(transform.EnsureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]string{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_project.listProject", "connection_error", err)
		return nil, err
	}
	maxItems := 1000
	opts := &metal.ListOptions{
		Page:    1,
		PerPage: maxItems,
	}
	for {
		items, resp, err := conn.Projects.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("equinix_metal_project.listProject", "query_error", err, "opts", opts, "resp", resp)
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

func getProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_project.getProject", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	project, resp, err := conn.Projects.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_project.getProject", "query_error", err, "id", id, "resp", resp)
		return nil, err
	}
	return project, nil
}
