package equinix

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableEquinixMetalSpotPrice(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "equinix_metal_spot_price",
		Description: "List of facilities and plans with their current spot price.",
		List: &plugin.ListConfig{
			Hydrate: listSpotPrice,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "facility", Type: proto.ColumnType_STRING, Description: "Facility code."},
			{Name: "plan", Type: proto.ColumnType_STRING, Description: "Server type."},
			{Name: "current_price", Type: proto.ColumnType_DOUBLE, Description: "Current spot price."},
		},
	}
}

type spotPriceRow struct {
	Facility     string  `json:"facility"`
	Plan         string  `json:"plan"`
	CurrentPrice float64 `json:"current_price"`
}

func listSpotPrice(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_spot_price.listSpotPrice", "connection_error", err)
		return nil, err
	}
	report, resp, err := conn.SpotMarket.Prices()
	if err != nil {
		plugin.Logger(ctx).Error("equinix_metal_spot_price.listSpotPrice", "query_error", err, "resp", resp)
		return nil, err
	}
	for facility, facilityData := range report {
		for plan, price := range facilityData {
			d.StreamListItem(ctx, spotPriceRow{
				Facility:     facility,
				Plan:         plan,
				CurrentPrice: price,
			})
		}
	}
	return nil, nil
}
