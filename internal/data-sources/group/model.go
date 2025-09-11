package group

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GroupModel represents an xMatters Group object in the Provider.
type GroupModel struct {
	GroupID         types.String `tfsdk:"group_id" tf:"required"`
	ID              types.String `tfsdk:"id" tf:"computed"`
	TargetName      types.String `tfsdk:"name" tf:"computed"`
	Description     types.String `tfsdk:"description" tf:"computed"`
	Status          types.String `tfsdk:"status" tf:"computed"`
	ExternalKey     types.String `tfsdk:"external_key" tf:"computed"`
	ExternallyOwned types.Bool   `tfsdk:"externally_owned" tf:"computed"`
	AllowDuplicates types.Bool   `tfsdk:"allow_duplicates" tf:"computed"`
	Timezone        types.String `tfsdk:"timezone" tf:"computed"`
	Site            types.String `tfsdk:"site" tf:"computed"`
	ObservedByAll   types.Bool   `tfsdk:"observed_by_all" tf:"computed"`
	Observers       types.Set    `tfsdk:"observers" tf:"computed"`
	Supervisors     types.Set    `tfsdk:"supervisors" tf:"computed"`
	Services        types.Set    `tfsdk:"services" tf:"computed"`
	GroupType       types.String `tfsdk:"group_type" tf:"computed"`
}
