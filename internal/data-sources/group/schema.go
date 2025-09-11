package group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *GroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Returns a group in your xMatters instance that matches the provided criteria.",
		Attributes:          GroupDataSourceSchema(),
	}
}

// GroupDataSourceSchema returns the schema for the group and groups data sources.
func GroupDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"group_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: describe.GroupIDSearch,
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupID,
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupTargetName,
		},
		"description": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupDescription,
		},
		// ...existing code...
	}
}
