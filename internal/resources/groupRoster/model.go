package groupRoster

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GroupRosterModel represents an xMatters GroupRoster object in the Provider.
type GroupRosterModel struct {
	ID          types.String `tfsdk:"id"`
	Group       types.Object `tfsdk:"group"`
	Members     types.Set    `tfsdk:"members"`
	LastUpdated types.String `tfsdk:"last_updated"`
}
