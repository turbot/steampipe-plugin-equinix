package metal

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-metal",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"metal_device":           tableMetalDevice(ctx),
			"metal_facility":         tableMetalFacility(ctx),
			"metal_event":            tableMetalEvent(ctx),
			"metal_operating_system": tableMetalOperatingSystem(ctx),
			"metal_organization":     tableMetalOrganization(ctx),
			"metal_plan":             tableMetalPlan(ctx),
			"metal_project":          tableMetalProject(ctx),
			"metal_project_ssh_key":  tableMetalProjectSSHKey(ctx),
			"metal_ssh_key":          tableMetalSSHKey(ctx),
		},
	}
	return p
}
