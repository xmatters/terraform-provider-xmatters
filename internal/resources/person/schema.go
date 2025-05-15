package person

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *PersonResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.PersonResourceDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: describe.PersonResourceID,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"target_name": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.PersonResourceTargetName,
			},
			"first_name": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.PersonResourceFirstName,
			},
			"last_name": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.PersonResourceLastName,
			},
			"roles": schema.SetAttribute{
				Required:            true,
				MarkdownDescription: describe.PersonResourceRoles,
				ElementType:         customTypes.CustomStringType{},
			},
			"status": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: describe.PersonResourceStatus,
				Validators: []validator.String{
					utils.StatusValidator{},
				},
			},
			"web_login": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.PersonResourceWebLogin,
			},
			"site": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: describe.PersonResourceSite,
				Validators: []validator.String{
					utils.UUIDValidator{},
				},
			},
			"timezone": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: describe.PersonResourceTimezone,
				Validators: []validator.String{
					utils.TimezoneValidator{},
				},
			},
			"language": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: describe.PersonResourceLanguage,
				Validators: []validator.String{
					utils.LanguageValidator{},
				},
			},
			"supervisors": schema.SetAttribute{
				Required:            true,
				MarkdownDescription: describe.PersonResourceSupervisors,
				ElementType:         types.StringType,
				Validators: []validator.Set{
					// Validator for each string in the list
					setvalidator.ValueStringsAre(utils.UUIDValidator{}),
				},
			},
			"phone_login": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: describe.PersonResourcePhoneLogin,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^\d{1,30}$`), "Phone login must be a number with a max of 30 characters."),
				},
			},
			"phone_pin": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				MarkdownDescription: describe.PersonResourcePhonePin,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^\d{1,30}$`), "Phone pin must be a number with a max of 30 characters."),
				},
			},
			"license_type": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.PersonResourceLicenseType,
			},
			"external_key": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Optional:            true,
				MarkdownDescription: describe.PersonResourceExternalKey,
			},
			"externally_owned": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: describe.PersonResourceExternallyOwned,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
