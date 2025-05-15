package serviceDependency

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/xmatters-go"
)

// ServiceDependencyModel represents an xMatters ServiceDependency object in the Provider.
type ServiceDependencyModel struct {
	ID               types.String `tfsdk:"id"`
	Service          types.String `tfsdk:"service_id"`
	DependentService types.String `tfsdk:"dependent_service_id"`
	LastUpdated      types.String `tfsdk:"last_updated"`
}

// ServiceDependencyParams is a method that takes the proposed configuration changes `ServiceDependencyModel` and builds the API representation in the form of `*xmatters.PushServiceDependencyParams`.
// The reverse of this method is `ServiceDependencyToModel` which handles building a state representation using the API response.
func (in ServiceDependencyModel) ServiceDependencyParams() xmatters.PushServiceDependencyParams {
	return xmatters.PushServiceDependencyParams{
		ID:                 in.ID.ValueString(),
		ServiceID:          in.Service.ValueString(),
		DependentServiceID: in.DependentService.ValueString(),
	}
}
