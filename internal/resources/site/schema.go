package site

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *SiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.SiteResourceDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: describe.SiteID,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"address1": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Optional:            true,
				MarkdownDescription: describe.SiteAddress1,
			},
			"address2": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Optional:            true,
				MarkdownDescription: describe.SiteAddress2,
			},
			"city": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Optional:            true,
				MarkdownDescription: describe.SiteCity,
			},
			"country": schema.StringAttribute{
				CustomType:          customTypes.CustomCountryType{},
				Required:            true,
				MarkdownDescription: describe.SiteCountry,
			},
			"language": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: describe.SiteLanguage,
				Validators: []validator.String{
					utils.LanguageValidator{},
				},
			},
			"latitude": schema.Float64Attribute{
				Optional:            true,
				MarkdownDescription: describe.SiteLatitude,
				Validators: []validator.Float64{
					utils.LatitudeValidator{},
				},
			},
			"longitude": schema.Float64Attribute{
				Optional:            true,
				MarkdownDescription: describe.SiteLongitude,
				Validators: []validator.Float64{
					utils.LongitudeValidator{},
				},
			},
			"name": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.SiteName,
			},
			"postal_code": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Optional:            true,
				MarkdownDescription: describe.SitePostalCode,
			},
			"state": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Optional:            true,
				MarkdownDescription: describe.SiteState,
			},
			"status": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: describe.SiteStatus,
				Validators: []validator.String{
					utils.StatusValidator{},
				},
			},
			"timezone": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: describe.SiteTimezone,
				Validators: []validator.String{
					utils.TimezoneValidator{},
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
