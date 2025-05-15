package service

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
func (r *ServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.ServiceResourceDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: describe.ServiceID,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.ServiceName,
			},
			"description": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Optional:            true,
				MarkdownDescription: describe.ServiceDescription,
			},
			"type": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.ServiceResourceType,
			},
			"tier": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: describe.ServiceResourceTier,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"owner": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: describe.ServiceResourceOwner,
				Validators: []validator.String{
					utils.UUIDValidator{},
				},
			},
			"links": schema.SetNestedAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: describe.ServiceResourceLinks,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"link_text": schema.StringAttribute{
							CustomType:          customTypes.CustomStringType{},
							Required:            true,
							MarkdownDescription: describe.ServiceLinkLabel,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
						},
						"url": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: describe.ServiceLinkURL,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`), "URL must be a valid URL."),
							},
						},
					},
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
