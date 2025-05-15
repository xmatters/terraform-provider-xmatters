package service

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

// ServiceModel represents an xMatters Service object in the Provider.
type ServiceModel struct {
	ID           types.String                  `tfsdk:"id"`
	Name         customTypes.CustomStringValue `tfsdk:"name"`
	Description  customTypes.CustomStringValue `tfsdk:"description"`
	Type         customTypes.CustomStringValue `tfsdk:"type"`
	Tier         types.String                  `tfsdk:"tier"`
	Owner        types.String                  `tfsdk:"owner"`
	ServiceLinks types.Set                     `tfsdk:"links"`
	LastUpdated  types.String                  `tfsdk:"last_updated"`
}

// ServiceParams is a method that takes the proposed configuration changes `ServiceModel` and builds the API representation in the form of `*xmatters.PushServiceParams`.
// The reverse of this method is `ServiceToModel` which handles building a state representation using the API response.
func (in ServiceModel) ServiceParams(diags *diag.Diagnostics) xmatters.PushServiceParams {
	return xmatters.PushServiceParams{
		ID:           in.ID.ValueString(),
		TargetName:   in.Name.ValueString(),
		Description:  in.Description.ValueStringPointer(),
		ServiceType:  in.Type.ValueString(),
		ServiceTier:  in.Tier.ValueStringPointer(),
		OwnedBy:      utils.ExpandGroupReferenceId(in.Owner),
		ServiceLinks: utils.ExpandServiceLinkSet(diags, in.ServiceLinks),
	}
}
