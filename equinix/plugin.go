package equinix

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-equinix",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"equinix_metal_device":           tableEquinixMetalDevice(ctx),
			"equinix_metal_capacity":         tableEquinixMetalCapacity(ctx),
			"equinix_metal_facility":         tableEquinixMetalFacility(ctx),
			"equinix_metal_event":            tableEquinixMetalEvent(ctx),
			"equinix_metal_operating_system": tableEquinixMetalOperatingSystem(ctx),
			"equinix_metal_organization":     tableEquinixMetalOrganization(ctx),
			"equinix_metal_plan":             tableEquinixMetalPlan(ctx),
			"equinix_metal_project":          tableEquinixMetalProject(ctx),
			"equinix_metal_project_ssh_key":  tableEquinixMetalProjectSSHKey(ctx),
			"equinix_metal_ssh_key":          tableEquinixMetalSSHKey(ctx),
			"equinix_metal_spot_price":       tableEquinixMetalSpotPrice(ctx),
		},
	}
	return p
}
