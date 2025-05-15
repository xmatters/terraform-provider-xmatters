package userQuotas

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceModel represents an xMatters Service object in the Provider.
type UserQuotasModel struct {
	StakeholderUsersEnabled types.Bool   `tfsdk:"stakeholder_users_enabled" tf:"computed"`
	StakeholderUsers        types.Object `tfsdk:"stakeholder_users" tf:"computed"`
	FullUsers               types.Object `tfsdk:"full_users" tf:"computed"`
}

// QuotaDetailsModel represents the quota details for a user license type.
type QuotaDetailsModel struct {
	Total  types.Int64 `tfsdk:"total" tf:"computed"`
	Active types.Int64 `tfsdk:"active" tf:"computed"`
	Unused types.Int64 `tfsdk:"unused" tf:"computed"`
}
