package equinix

import (
	"context"

	metal "github.com/packethost/packngo"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableEquinixMetalProjectSSHKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "equinix_metal_project_ssh_key",
		Description: "List all SSH keys in all projects.",
		List: &plugin.ListConfig{
			ParentHydrate: listProject,
			Hydrate:       listProjectSSHKey,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "project_id", Type: proto.ColumnType_STRING, Hydrate: getParentProjectID, Transform: transform.FromValue(), Description: "ID of the Project."},
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

func listProjectSSHKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_project_ssh_key.listProjectSSHKey", "connection_error", err)
		return nil, err
	}
	maxItems := 1000
	opts := &metal.ListOptions{
		Page:    1,
		PerPage: maxItems,
	}
	project := h.Item.(metal.Project)
	for {
		items, resp, err := conn.Projects.ListSSHKeys(project.ID, opts)
		if err != nil {
			plugin.Logger(ctx).Error("equinix_metal_project_ssh_key.listProjectSSHKey", "query_error", err, "opts", opts, "resp", resp)
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
