package equinix

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableEquinixMetalCapacity(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "equinix_metal_capacity",
		Description: "List of facilities and plans with their current capacity.",
		List: &plugin.ListConfig{
			Hydrate: listCapacity,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "facility", Type: proto.ColumnType_STRING, Description: "Facility code."},
			{Name: "plan", Type: proto.ColumnType_STRING, Description: "Server type."},
			{Name: "level", Type: proto.ColumnType_STRING, Description: "Current capacity level: normal, limited or unavailable."},
		},
	}
}

type capacityRow struct {
	Facility string `json:"facility"`
	Plan     string `json:"plan"`
	Level    string `json:"level"`
}

func listCapacity(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_capacity.listCapacity", "connection_error", err)
		return nil, err
	}
	report, resp, err := conn.CapacityService.List()
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_capacity.listCapacity", "query_error", err, "resp", resp)
		return nil, err
	}
	for facility, facilityData := range *report {
		for plan, planData := range facilityData {
			d.StreamListItem(ctx, capacityRow{
				Facility: facility,
				Plan:     plan,
				Level:    planData.Level,
			})
		}
	}
	return nil, nil
}
