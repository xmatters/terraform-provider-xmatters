package groups

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

// GroupsModel represents an xMatters Groups object in the Provider.
type GroupsModel struct {
	Search  *GroupsSearchModel  `tfsdk:"search" tf:"optional"`
	Filters *GroupsFiltersModel `tfsdk:"filters" tf:"optional"`
	Options *GroupsOptionsModel `tfsdk:"options" tf:"optional"`
	Groups  types.List          `tfsdk:"groups" tf:"computed"`
}

// GroupsSearchModel contains the search fields for the Provider's Groups data source.
type GroupsSearchModel struct {
	Terms   types.String `tfsdk:"terms" tf:"optional"`
	Operand types.String `tfsdk:"operand" tf:"optional"`
	Fields  types.List   `tfsdk:"fields" tf:"optional"`
}

type GroupsFiltersModel struct {
	GroupType    types.String `tfsdk:"group_type" tf:"optional"`
	MemberExists types.String `tfsdk:"member_exists" tf:"optional"`
	Members      types.List   `tfsdk:"members" tf:"optional"`
	Sites        types.List   `tfsdk:"sites" tf:"optional"`
	Status       types.String `tfsdk:"status" tf:"optional"`
	Supervisors  types.List   `tfsdk:"supervisors" tf:"optional"`
}

// GroupsOptionsModel contains the options fields for the Provider's Groups data source.
type GroupsOptionsModel struct {
	SortBy    types.String `tfsdk:"sort_by" tf:"optional"`
	SortOrder types.String `tfsdk:"sort_order" tf:"optional"`
}

// APIParams returns the xmatters.GetGroupsParams object based on the GroupsModel instance.
func (in GroupsModel) APIParams(diags *diag.Diagnostics) xmatters.GetGroupsParams {
	groupsParams := xmatters.GetGroupsParams{
		Embed: "supervisors,observers",
	}
	// Check for user provider Search fields
	if in.Search != nil {
		groupsParams.Terms = in.Search.Terms.ValueString()
		groupsParams.Operand = in.Search.Operand.ValueString()
		groupsParams.Fields = utils.ExpandStringList(diags, in.Search.Fields)
	}
	// Check for user provider Filter fields
	if in.Filters != nil {
		groupsParams.GroupType = in.Filters.GroupType.ValueString()
		groupsParams.MemberExists = in.Filters.MemberExists.ValueString()
		groupsParams.Members = utils.ExpandStringList(diags, in.Filters.Members)
		groupsParams.Sites = utils.ExpandStringList(diags, in.Filters.Sites)
		groupsParams.Status = in.Filters.Status.ValueString()
		groupsParams.Supervisors = utils.ExpandStringList(diags, in.Filters.Supervisors)
	}
	// Check for user provider Options fields
	if in.Options != nil {
		groupsParams.SortBy = in.Options.SortBy.ValueString()
		groupsParams.SortOrder = in.Options.SortOrder.ValueString()
	}
	return groupsParams
}
