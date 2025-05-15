package service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *ServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.ServiceDataSourceDescription,
		Attributes:          ServiceDataSourceSchema(),
	}
}

// ServiceDataSourceSchema returns the schema for the service and services data sources.
func ServiceDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"service_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: describe.ServiceIDSearch,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(3),
			},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.ServiceID,
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.ServiceName,
		},
		"description": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.ServiceDescription,
		},
		"type": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.ServiceType,
		},
		"tier": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.ServiceTier,
		},
		"owner": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.ServicesOwner,
		},
		"links": schema.SetNestedAttribute{
			Computed:            true,
			MarkdownDescription: describe.ServiceLinks,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"link_text": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: describe.ServiceLinkLabel,
					},
					"url": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: describe.ServiceLinkURL,
					},
				},
			},
		},
	}
}
