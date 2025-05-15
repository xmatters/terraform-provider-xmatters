package site

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SiteModel represents an xMatters Site object in the Provider.
type SiteModel struct {
	SiteID     types.String  `tfsdk:"site_id" tf:"required"`
	Address1   types.String  `tfsdk:"address1" tf:"computed"`
	Address2   types.String  `tfsdk:"address2" tf:"computed"`
	City       types.String  `tfsdk:"city" tf:"computed"`
	Country    types.String  `tfsdk:"country" tf:"computed"`
	ID         types.String  `tfsdk:"id" tf:"computed"`
	Language   types.String  `tfsdk:"language" tf:"computed"`
	Latitude   types.Float64 `tfsdk:"latitude" tf:"computed"`
	Longitude  types.Float64 `tfsdk:"longitude" tf:"computed"`
	Name       types.String  `tfsdk:"name" tf:"computed"`
	PostalCode types.String  `tfsdk:"postal_code" tf:"computed"`
	State      types.String  `tfsdk:"state" tf:"computed"`
	Status     types.String  `tfsdk:"status" tf:"computed"`
	Timezone   types.String  `tfsdk:"timezone" tf:"computed"`
}
