package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

// XMattersProviderModel describes the provider data model.
type XMattersProviderModel struct {
	BaseURL types.String `tfsdk:"base_url" tf:"required"`
	Auth    types.Object `tfsdk:"auth" tf:"required"`
}
