package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/userQuotas"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

// UserQuotasToModel
func TestUserQuotasToModel(t *testing.T) {
	testUserQuotas := xmatters.UserQuotas{
		StakeholderUsersEnabled: utils.RandBoolPointer(),
		StakeholderUsers: &xmatters.QuotaDetails{
			Total:  utils.RandInt64Pointer(),
			Active: utils.RandInt64Pointer(),
			Unused: utils.RandInt64Pointer(),
		},
		FullUsers: &xmatters.QuotaDetails{
			Total:  utils.RandInt64Pointer(),
			Active: utils.RandInt64Pointer(),
			Unused: utils.RandInt64Pointer(),
		},
	}
	type args struct {
		diags *diag.Diagnostics
		in    xmatters.UserQuotas
	}
	tests := []struct {
		name     string
		args     args
		expected userQuotas.UserQuotasModel
	}{
		{
			name: "empty service",
			args: args{
				diags: &diag.Diagnostics{},
				in:    xmatters.UserQuotas{},
			},
			expected: userQuotas.UserQuotasModel{
				StakeholderUsers: types.ObjectNull(utils.UserQuotaDetailsObjectType.AttrTypes),
				FullUsers:        types.ObjectNull(utils.UserQuotaDetailsObjectType.AttrTypes),
			},
		},
		{
			name: "valid service",
			args: args{
				diags: &diag.Diagnostics{},
				in:    testUserQuotas,
			},
			expected: userQuotas.UserQuotasModel{
				StakeholderUsersEnabled: types.BoolPointerValue(testUserQuotas.StakeholderUsersEnabled),
				StakeholderUsers: types.ObjectValueMust(
					utils.UserQuotaDetailsObjectType.AttrTypes,
					map[string]attr.Value{
						"total":  types.Int64PointerValue(testUserQuotas.StakeholderUsers.Total),
						"active": types.Int64PointerValue(testUserQuotas.StakeholderUsers.Active),
						"unused": types.Int64PointerValue(testUserQuotas.StakeholderUsers.Unused),
					},
				),
				FullUsers: types.ObjectValueMust(
					utils.UserQuotaDetailsObjectType.AttrTypes,
					map[string]attr.Value{
						"total":  types.Int64PointerValue(testUserQuotas.FullUsers.Total),
						"active": types.Int64PointerValue(testUserQuotas.FullUsers.Active),
						"unused": types.Int64PointerValue(testUserQuotas.FullUsers.Unused),
					},
				),
			},
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := userQuotas.UserQuotasToModel(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.expected, got)
		})
	}
}
