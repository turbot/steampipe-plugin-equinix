package equinix

import (
	"context"

	metal "github.com/packethost/packngo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableEquinixMetalOrganization(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "equinix_metal_organization",
		Description: "Organizations the user has access to.",
		List: &plugin.ListConfig{
			Hydrate: listOrganization,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			Hydrate:           getOrganization,
			ShouldIgnoreError: isNotFoundError([]string{"404 Not found"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the organization."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the organization."},
			// Other columns
			{Name: "address", Type: proto.ColumnType_JSON, Description: "Address of the organization."},
			{Name: "billing_phone", Type: proto.ColumnType_STRING, Description: "Billing phone number of the organization."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the organization was created."},
			{Name: "credit_amount", Type: proto.ColumnType_DOUBLE, Description: "Credit amount available to the organization."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description for the organization."},
			{Name: "href", Type: proto.ColumnType_STRING, Description: "URL for the organization."},
			{Name: "logo", Type: proto.ColumnType_STRING, Description: "Logo of the organization."},
			{Name: "logo_thumb", Type: proto.ColumnType_STRING, Description: "Logo thumbnail of the organization."},
			{Name: "main_phone", Type: proto.ColumnType_STRING, Description: "Main phone number of the organization."},
			// TODO - Does this field actually work?
			//{Name: "members", Type: proto.ColumnType_JSON, Description: "Members of the organization."},
			{Name: "member_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Users").Transform(usersToIDs), Description: "IDs of the members of this organization."},
			// TODO - Does this field actually work?
			//{Name: "owners", Type: proto.ColumnType_JSON, Description: "Owners of the organization."},
			{Name: "owner_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Owners").Transform(usersToIDs), Description: "IDs of the owners of this organization."},
			// TODO - Should this be limited to the project IDs only? Project data appears incomplete?
			//{Name: "projects", Type: proto.ColumnType_JSON, Description: "Projects for the organization."},
			{Name: "project_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Projects").Transform(projectsToIDs), Description: "Projects for the organization."},
			{Name: "tax_id", Type: proto.ColumnType_STRING, Description: "Tax ID of the organization."},
			{Name: "twitter", Type: proto.ColumnType_STRING, Description: "Twitter handle for the organization."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the organization was updated."},
			{Name: "website", Type: proto.ColumnType_STRING, Description: "Website for the organization."},
			// Resource columns
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromField("URL").Transform(transform.EnsureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]string{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_organization.listOrganization", "connection_error", err)
		return nil, err
	}
	maxItems := 1000
	opts := &metal.ListOptions{
		Page:    1,
		PerPage: maxItems,
	}
	for {
		items, resp, err := conn.Organizations.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("equinix_metal_organization.listOrganization", "query_error", err, "opts", opts, "resp", resp)
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

func getOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_organization.getOrganization", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	organization, resp, err := conn.Organizations.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_organization.getOrganization", "query_error", err, "id", id, "resp", resp)
		return nil, err
	}
	return organization, nil
}
