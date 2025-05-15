package services

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

// ServicesModel contains the search fields and return values for the Provider's Services data source.
type ServicesModel struct {
	Owner    types.String         `tfsdk:"owner" tf:"optional"`
	Search   *ServicesSearchModel `tfsdk:"search" tf:"optional"`
	Services types.List           `tfsdk:"services" tf:"computed"`
}

// ServicesSearchModel contains the search fields for the Provider's Services data source.
type ServicesSearchModel struct {
	Terms   types.String `tfsdk:"terms" tf:"optional"`
	Operand types.String `tfsdk:"operand" tf:"optional"`
	Fields  types.List   `tfsdk:"fields" tf:"optional"`
}

// APIParams returns the xmatters.GetServicesParams object based on the ServicesModel instance.
func (in ServicesModel) APIParams(diags *diag.Diagnostics) xmatters.GetServicesParams {
	params := xmatters.GetServicesParams{
		OwnedBy: in.Owner.ValueString(),
	}
	// Check for user provided Search fields
	if in.Search != nil {
		params.Search = in.Search.Terms.ValueString()
		params.Operand = in.Search.Operand.ValueString()
		params.Fields = utils.ExpandStringList(diags, in.Search.Fields)
	}
	return params
}
