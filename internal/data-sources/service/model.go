package service

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceModel represents an xMatters Service object in the Provider.
type ServiceModel struct {
	ServiceID    types.String `tfsdk:"service_id" tf:"required"`
	ID           types.String `tfsdk:"id" tf:"computed"`
	Name         types.String `tfsdk:"name" tf:"computed"`
	Description  types.String `tfsdk:"description" tf:"computed"`
	Type         types.String `tfsdk:"type" tf:"computed"`
	Tier         types.String `tfsdk:"tier" tf:"computed"`
	Owner        types.String `tfsdk:"owner" tf:"computed"`
	ServiceLinks types.Set    `tfsdk:"links" tf:"computed"`
}
