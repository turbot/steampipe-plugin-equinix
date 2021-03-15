package metal

import (
	"context"

	metal "github.com/packethost/packngo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableMetalEvent(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "metal_event",
		Description: "Events the user has access to.",
		List: &plugin.ListConfig{
			Hydrate: listEvent,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			Hydrate:           getEvent,
			ShouldIgnoreError: isNotFoundError([]string{"404 Not found"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the event."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(timestampToIsoTimestamp), Description: "When the event was created."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the event."},
			// Other columns
			{Name: "body", Type: proto.ColumnType_STRING, Description: "Body of the event."},
			{Name: "href", Type: proto.ColumnType_STRING, Description: "URL for this event."},
			// TODO: interpolated seems to be identical to the body?
			// {Name: "interpolated", Type: proto.ColumnType_STRING, Description: "TODO"},
			// TODO: relationships URLs seem broken / incomplete? What type are they?
			{Name: "relationships", Type: proto.ColumnType_JSON, Description: "Relationships from this event."},
			// TODO: state is always empty?
			{Name: "state", Type: proto.ColumnType_STRING, Description: "State of the event."},
		},
	}
}

func listEvent(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("metal_event.listEvent", "connection_error", err)
		return nil, err
	}
	maxItems := 1000
	opts := &metal.ListOptions{
		Page:    1,
		PerPage: maxItems,
	}
	for {
		items, resp, err := conn.Events.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("metal_event.listEvent", "query_error", err, "opts", opts, "resp", resp)
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

func getEvent(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("metal_event.getEvent", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	event, resp, err := conn.Events.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("metal_event.getEvent", "query_error", err, "id", id, "resp", resp)
		return nil, err
	}
	return event, nil
}
