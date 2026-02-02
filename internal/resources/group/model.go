package group

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

// GroupModel represents an xMatters Group object in the Provider.
type GroupModel struct {
	ID                types.String                  `tfsdk:"id"`
	Name              customTypes.CustomStringValue `tfsdk:"name"`
	Status            types.String                  `tfsdk:"status"`
	Description       customTypes.CustomStringValue `tfsdk:"description"`
	GroupType         types.String                  `tfsdk:"group_type"`
	AllowDuplicates   types.Bool                    `tfsdk:"allow_duplicates"`
	Site              types.String                  `tfsdk:"site"`
	ObservedByAll     types.Bool                    `tfsdk:"observed_by_all"`
	Observers         types.Set                     `tfsdk:"observers"`
	UseDefaultDevices types.Bool                    `tfsdk:"use_default_devices"`
	Supervisors       types.Set                     `tfsdk:"supervisors"`
	ExternalKey       customTypes.CustomStringValue `tfsdk:"external_key"`
	ExternallyOwned   types.Bool                    `tfsdk:"externally_owned"`
	Criteria          types.Object                  `tfsdk:"criteria"`
	LastUpdated       types.String                  `tfsdk:"last_updated"`
}

// GroupParams is a method that takes the proposed configuration changes `GroupModel` and builds the API representation in the form of `*xmatters.PushGroupParams`.
// The reverse of this method is `GroupToModel` which handles building a state representation using the API response.
func (in GroupModel) GroupParams(diags *diag.Diagnostics) xmatters.PushGroupParams {
	groupParams := xmatters.PushGroupParams{
		ID:          in.ID.ValueString(),
		TargetName:  in.Name.ValueString(),
		Description: in.Description.ValueString(),
		ExternalKey: in.ExternalKey.ValueString(),
		GroupType:   in.GroupType.ValueString(),
		Observers:   utils.ExpandReferenceNameSet(diags, in.Observers),
		Site:        in.Site.ValueString(),
		Status:      in.Status.ValueString(),
		Supervisors: utils.ExpandReferenceIDSet(diags, in.Supervisors),
		Criteria:    utils.ExpandGroupCriteriaObject(diags, in.Criteria),
	}
	// Following attributes are bool or integer parameters set as computed so should only be sent to the API if configured
	if !in.ObservedByAll.IsUnknown() {
		groupParams.ObservedByAll = in.ObservedByAll.ValueBoolPointer()
	}
	if !in.UseDefaultDevices.IsUnknown() {
		groupParams.UseDefaultDevices = in.UseDefaultDevices.ValueBoolPointer()
	}
	if !in.AllowDuplicates.IsUnknown() {
		groupParams.AllowDuplicates = in.AllowDuplicates.ValueBoolPointer()
	}
	if !in.ExternallyOwned.IsUnknown() {
		groupParams.ExternallyOwned = in.ExternallyOwned.ValueBoolPointer()
	}
	return groupParams
}
