package site

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *SiteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.SiteDataSourceDescription,
		Attributes:          SiteDataSourceSchema(),
	}
}

// SiteDataSourceSchema returns the schema for the site and sites data sources.
func SiteDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"site_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: describe.SiteIDSearch,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(3),
			},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.SiteID,
		},
		"address1": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.SiteAddress1,
		},
		"address2": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.SiteAddress2,
		},
		"city": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.SiteCity,
		},
		"country": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.SiteCountry,
		},
		"language": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.SiteLanguage,
		},
		"latitude": schema.Float64Attribute{
			Computed:            true,
			MarkdownDescription: describe.SiteLatitude,
		},
		"longitude": schema.Float64Attribute{
			Computed:            true,
			MarkdownDescription: describe.SiteLongitude,
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.SiteName,
		},
		"postal_code": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.SitePostalCode,
		},
		"state": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.SiteState,
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.SiteStatus,
		},
		"timezone": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.SiteTimezone,
		},
	}
}
