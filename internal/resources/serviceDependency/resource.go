package serviceDependency

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/helpers"
	"github.com/xmatters/xmatters-go"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &ServiceDependencyResource{}
	_ resource.ResourceWithConfigure   = &ServiceDependencyResource{}
	_ resource.ResourceWithImportState = &ServiceDependencyResource{}
)

// NewServiceDependencyResource is a helper function to simplify the provider implementation.
func NewServiceDependencyResource() resource.Resource {
	return &ServiceDependencyResource{}
}

// ServiceDependencyResource defines the resource implementation.
type ServiceDependencyResource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the resource type name.
func (r *ServiceDependencyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_dependency"
}

// Configure adds the provider configured client to the resource.
func (r *ServiceDependencyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*xmatters.XMattersAPI)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Create creates the resource and sets the initial Terraform state.
func (r *ServiceDependencyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceDependencyModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	createParams := plan.ServiceDependencyParams()
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new xMatters ServiceDependency
	serviceDepReturn, err := r.client.PushServiceDependency(createParams)
	if err != nil {
		var xmErr xmatters.XMattersError
		if errors.As(err, &xmErr) && xmErr.Code == 409 {
			resp.Diagnostics.AddError(
				"xMatters ServiceDependency already exists.",
				fmt.Sprintf("Remove the serviceDependancy from the Terraform State and Import.\nxMatters API Error: %d - %s.\n  %s", xmErr.Code, xmErr.Reason, xmErr.Message),
			)
			return
		}
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("serviceDependency", err))
		return
	}

	// Map response body to schema and populate Computed attribute values
	state := ServiceDependencyToModel(serviceDepReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Created xMatters Service Dependency", map[string]any{"success": true})
}

// Read refreshes the Terraform state with the latest data.
func (r *ServiceDependencyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceDependencyModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed ServiceDependency from xMatters
	serviceDepReturn, err := r.client.GetServiceDependency(state.ID.ValueString())
	if err != nil {
		var xMerr xmatters.XMattersError
		if errors.As(err, &xMerr) && xMerr.Code == 404 {
			// If the ServiceDependency is not found, remove it from the Terraform state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("ServiceDependency", err))
		return
	}

	// Capture state `last_updated` timestamp
	last_updated := state.LastUpdated
	// Overwrite ServiceDependency with refreshed state
	state = ServiceDependencyToModel(serviceDepReturn)

	// Ensure `last_updated` is preserved and set to the new state
	state.LastUpdated = last_updated

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Retrieved xMatters ServiceDependency data", map[string]any{"success": true})
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *ServiceDependencyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ServiceDependencyModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	updateParams := plan.ServiceDependencyParams()
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing xMatters ServiceDependency and fetch updated ServiceDependency
	serviceDepReturn, err := r.client.PushServiceDependency(updateParams)
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("ServiceDependency", err))
		return
	}

	// Updated resource state with updated data and timestamp
	state := ServiceDependencyToModel(serviceDepReturn)

	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Save state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Updated xMatters ServiceDependency", map[string]any{"success": true})
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *ServiceDependencyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceDependencyModel

	// Retrieve values from state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing ServiceDependency
	err := r.client.DeleteServiceDependency(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DeleteErrorDiagnostic("ServiceDependency", err))
		return
	}
}

// ImportState imports a resource from an existing xMatters resource.
func (r *ServiceDependencyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// ServiceDependencyToModel is a function that takes the API payload `xmatters.ServiceDependency` and builds the state representation in the form of `ServiceDependencyModel`.
// The reverse of this method is `ServiceDependencyParams` which handles building an API representation using the proposed config.
func ServiceDependencyToModel(serviceDep xmatters.ServiceDependency) ServiceDependencyModel {
	model := ServiceDependencyModel{
		ID:               types.StringPointerValue(serviceDep.ID),
		Service:          utils.FlattenServiceReferenceID(serviceDep.Service),
		DependentService: utils.FlattenServiceReferenceID(serviceDep.DependentService),
	}
	return model
}
