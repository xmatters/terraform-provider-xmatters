package group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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
func (r *GroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Returns a group in your xMatters instance that matches the provided criteria.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: describe.GroupResourceID,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.GroupResourceTargetName,
			},
			"status": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: describe.GroupResourceStatus,
				Validators: []validator.String{
					utils.StatusValidator{},
				},
			},
			"description": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Computed:            true,
				Optional:            true,
				MarkdownDescription: describe.GroupResourceDescription,
			},
			"group_type": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: describe.GroupResourceType,
			},
			"allow_duplicates": schema.BoolAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: describe.GroupResourceAllowDuplicates,
			},
			"site": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: describe.GroupResourceSite,
				Validators: []validator.String{
					utils.UUIDValidator{},
				},
			},
			"observed_by_all": schema.BoolAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: describe.GroupResourceObservedByAll,
			},
			// Add validation that prevents observers from being set if observed_by_all is true
			"observers": schema.SetAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: describe.GroupResourceObservers,
				ElementType:         customTypes.CustomStringType{},
			},
			"use_default_devices": schema.BoolAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: describe.GroupResourceUseDefaultDevices,
			},
			"supervisors": schema.SetAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: describe.GroupResourceSupervisors,
				ElementType:         types.StringType,
				Validators: []validator.Set{
					// Validator for each string in the list
					setvalidator.ValueStringsAre(utils.UUIDValidator{}),
				},
			},
			"external_key": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Computed:            true,
				Optional:            true,
				MarkdownDescription: describe.GroupResourceExternalKey,
			},
			"externally_owned": schema.BoolAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: describe.GroupResourceExternallyOwned,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
