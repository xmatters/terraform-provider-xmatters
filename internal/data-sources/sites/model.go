package sites

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

// SitesModel contains the search fields and return values for the Provider's Sites data source.
type SitesModel struct {
	Search  *SitesSearchModel `tfsdk:"search" tf:"optional"`
	Filters *SitesFilterModel `tfsdk:"filters" tf:"optional"`
	Sites   types.List        `tfsdk:"sites" tf:"computed"`
}

// SitesSearchModel contains the search fields for the Provider's Sites data source.
type SitesSearchModel struct {
	Terms   types.String `tfsdk:"terms" tf:"optional"`
	Operand types.String `tfsdk:"operand" tf:"optional"`
	Fields  types.List   `tfsdk:"fields" tf:"optional"`
}

// SitesFilterModel contains the filter fields for the Provider's Sites data source.
type SitesFilterModel struct {
	Country  types.String `tfsdk:"country" tf:"optional"`
	Geocoded types.Bool   `tfsdk:"geocoded" tf:"optional"`
	Status   types.String `tfsdk:"status" tf:"optional"`
}

// APIParams returns the xmatters.GetSitesParams object based on the SitesModel instance.
func (in SitesModel) APIParams(diags *diag.Diagnostics) xmatters.GetSitesParams {
	params := xmatters.GetSitesParams{}
	// Check for user provided Search fields
	if in.Search != nil {
		params.Search = in.Search.Terms.ValueString()
		params.Operand = in.Search.Operand.ValueString()
		params.Fields = utils.ExpandStringList(diags, in.Search.Fields)
	}
	if in.Filters != nil {
		params.Country = in.Filters.Country.ValueString()
		params.Geocoded = in.Filters.Geocoded.ValueBoolPointer()
		params.Status = in.Filters.Status.ValueString()
	}
	return params
}
