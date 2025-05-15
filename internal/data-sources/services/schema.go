package services

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *ServicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.ServicesDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"owner": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: describe.ServicesOwner,
			},
			"search": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: describe.ServicesSearch,
				Attributes:          ServicesDataSourceSearchSchema(),
			},
			"services": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: describe.ServicesList,
				NestedObject: schema.NestedAttributeObject{
					Attributes: ServiceDataSourceSchema(),
				},
			},
		},
	}
}

// ServiceDataSourceSchema is a helper function to simplify the ServicesDataSource Schema implementation.
// It represents the Provider's implmentation of the xMatters Service object.
func ServiceDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
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

// ServicesDataSourceSearchSchema is a helper function to simplify the ServicesDataSource Schema implementation.
// It represents the Provider's optional search parameters to filter the list of services returned.
func ServicesDataSourceSearchSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"terms": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.ServicesSearchTerms,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(2),
			},
		},
		"operand": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.ServicesSearchOperand,
			Validators: []validator.String{
				utils.OperandValidator{},
			},
		},
		"fields": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.ServicesSearchFields,
			ElementType:         types.StringType,
		},
	}
}
