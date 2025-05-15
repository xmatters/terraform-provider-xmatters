package sites

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
func (r *SitesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.SitesDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"search": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: describe.SitesSearch,
				Attributes:          SitesDataSourceSearchSchema(),
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: describe.SitesFilters,
				Attributes:          SiteDataSourceFiltersSchema(),
			},
			"sites": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: describe.SitesList,
				NestedObject: schema.NestedAttributeObject{
					Attributes: SiteDataSourceSchema(),
				},
			},
		},
	}
}

// SiteDataSourceSchema is a helper function to simplify the SitesDataSource Schema implementation.
// It represents the Provider's implmentation of the xMatters Site object.
func SiteDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
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

// SitesDataSourceSearchSchema is a helper function to simplify the SitesDataSource Schema implementation.
// It represents the Provider's optional search parameters to filter the list of sites returned.
func SitesDataSourceSearchSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"terms": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.SitesSearchTerms,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(2),
			},
		},
		"operand": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.SitesSearchOperand,
			Validators: []validator.String{
				utils.OperandValidator{},
			},
		},
		"fields": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.SitesSearchFields,
			ElementType:         types.StringType,
		},
	}
}

// SiteDataSourceFiltersSchema is a helper function to simplify the SitesDataSource Schema implementation.
// It represents the Provider's optional filter parameters to filter the list of sites returned.
func SiteDataSourceFiltersSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"country": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.SiteFilterCountry,
		},
		"geocoded": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: describe.SiteFilterGeocoded,
		},
		"status": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.SiteFilterStatus,
			Validators: []validator.String{
				utils.StatusValidator{},
			},
		},
	}
}
