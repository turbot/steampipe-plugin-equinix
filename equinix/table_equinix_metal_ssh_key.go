package equinix

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableEquinixMetalSSHKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "equinix_metal_ssh_key",
		Description: "SSH keys for the user.",
		List: &plugin.ListConfig{
			Hydrate: listSSHKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			Hydrate:           getSSHKey,
			ShouldIgnoreError: isNotFoundError([]string{"404 Not found"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the SSH key."},
			{Name: "label", Type: proto.ColumnType_STRING, Description: "Label for the SSH key."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the SSH key was created."},
			{Name: "fingerprint", Type: proto.ColumnType_STRING, Description: "Fingerprint of the SSH key."},
			{Name: "href", Type: proto.ColumnType_STRING, Description: "URL of the SSH key."},
			// NOTE: renamed to public_key for clarity
			{Name: "public_key", Type: proto.ColumnType_STRING, Transform: transform.FromField("Key"), Description: "Public key."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the SSH key was updated."},
			// Resource columns
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromField("URL").Transform(transform.EnsureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]string{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Label"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listSSHKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_ssh_key.listSSHKey", "connection_error", err)
		return nil, err
	}
	items, resp, err := conn.SSHKeys.List()
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_ssh_key.listSSHKey", "query_error", err, "resp", resp)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getSSHKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_ssh_key.getSSHKey", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	item, resp, err := conn.SSHKeys.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_ssh_key.getSSHKey", "query_error", err, "id", id, "resp", resp)
		return nil, err
	}
	return item, nil
}
