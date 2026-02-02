package groupMembers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GroupMembersModel represents an xMatters GroupMembers object in the Provider.
type GroupMembersModel struct {
	ID          types.String `tfsdk:"id"`
	Group       types.Object `tfsdk:"group"`
	Members     types.Set    `tfsdk:"members"`
	LastUpdated types.String `tfsdk:"last_updated"`
}
