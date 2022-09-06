package equinix

import (
	"context"

	metal "github.com/packethost/packngo"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableEquinixMetalDevice(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "equinix_metal_device",
		Description: "List all devices in all projects.",
		List: &plugin.ListConfig{
			ParentHydrate: listProject,
			Hydrate:       listDevice,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			Hydrate:           getDevice,
			ShouldIgnoreError: isNotFoundError([]string{"404 Not found"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "project_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Project.URL").Transform(hrefToID), Description: "ID of the Project."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the device."},
			{Name: "hostname", Type: proto.ColumnType_STRING, Description: "Hostname of the device."},
			// Other columns
			{Name: "always_pxe", Type: proto.ColumnType_STRING, Description: "True if PXE is always enabled for the device."},
			{Name: "billing_cycle", Type: proto.ColumnType_STRING, Description: "Billing cycle for the device."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the device was created."},
			{Name: "customdata", Type: proto.ColumnType_JSON, Description: "Custom data associated with the device."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the device."},
			//{Name: "facility", Type: proto.ColumnType_JSON, Description: "Facility for the device."},
			{Name: "facility_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Facility.ID"), Description: "Facility for the device."},
			{Name: "hardware_reservation", Type: proto.ColumnType_JSON, Description: "Hardware reservation for the device."},
			{Name: "href", Type: proto.ColumnType_STRING, Description: "URL of the device."},
			// Use the full IP info here, easier than a separate table per device
			{Name: "ip_addresses", Type: proto.ColumnType_JSON, Description: "Network configuration for the device."},
			{Name: "ipxe_script_url", Type: proto.ColumnType_STRING, Description: "IPXE script URL for the device."},
			{Name: "locked", Type: proto.ColumnType_BOOL, Description: "True if the device is locked."},
			//{Name: "metro", Type: proto.ColumnType_JSON, Description: "Metro for the device."},
			{Name: "metro_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Metro.ID"), Description: "Metro for the device."},
			{Name: "network_ports", Type: proto.ColumnType_JSON, Description: "List of network ports for the device."},
			//{Name: "operating_system", Type: proto.ColumnType_JSON, Description: "OS for the device."},
			{Name: "operating_system_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("OS.Slug"), Description: "OS for the device."},
			//{Name: "plan", Type: proto.ColumnType_JSON, Description: "Plan for the device."},
			{Name: "plan_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Plan.ID"), Description: "Plan for the device."},
			//{Name: "project", Type: proto.ColumnType_JSON, Description: "Project for the device."},
			// TODO - do the provisioning fields work? use events table instead?
			{Name: "provisioning_events", Type: proto.ColumnType_JSON, Description: "Provisioning events for the device."},
			{Name: "provisioning_percentage", Type: proto.ColumnType_DOUBLE, Description: "Provisioning percentage complete for the device."},
			{Name: "root_password", Type: proto.ColumnType_STRING, Description: "Root password for the device. Only available for 24hr after launch."},
			{Name: "short_id", Type: proto.ColumnType_STRING, Description: "Short ID for the device."},
			{Name: "spot_instance", Type: proto.ColumnType_BOOL, Description: "True if the device is a spot instance."},
			{Name: "spot_price_max", Type: proto.ColumnType_DOUBLE, Description: "Maximum spot price allowed for the device."},
			//{Name: "ssh_keys", Type: proto.ColumnType_JSON, Description: "SSH Keys deployed to the device."},
			{Name: "ssh_key_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("SSHKeys").Transform(sshKeysToIDs), Description: "SSH Keys deployed to the device."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "State of the device."},
			// TODO - does this work?
			{Name: "storage", Type: proto.ColumnType_JSON, Description: "Storage details for the device."},
			{Name: "switch_uuid", Type: proto.ColumnType_STRING, Description: "Switch UUID for the device."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags"), Description: "Tags for the device in source list form."},
			{Name: "termination_time", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the device was terminated."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the device was updated."},
			{Name: "user", Type: proto.ColumnType_STRING, Description: "User for the device."},
			{Name: "userdata", Type: proto.ColumnType_STRING, Description: "User data for the device."},
			//{Name: "volumes", Type: proto.ColumnType_JSON, Description: "Volumes for the device."},
			{Name: "volume_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Volumes").Transform(volumesToIDs), Description: "Volumes for the device."},
			// Resource columns
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromField("Href").Transform(transform.EnsureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags").Transform(transform.StringArrayToMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Hostname"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listDevice(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_device.listDevice", "connection_error", err)
		return nil, err
	}
	maxItems := 1000
	opts := &metal.ListOptions{
		Page:    1,
		PerPage: maxItems,
	}
	project := h.Item.(metal.Project)
	for {
		items, resp, err := conn.Devices.List(project.ID, opts)
		if err != nil {
			plugin.Logger(ctx).Error("equinix_metal_device.listDevice", "query_error", err, "opts", opts, "resp", resp)
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

func getDevice(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_device.getDevice", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	item, resp, err := conn.Devices.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_device.getDevice", "query_error", err, "id", id, "resp", resp)
		return nil, err
	}
	return item, nil
}
