package service

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/helpers"
	"github.com/xmatters/xmatters-go"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &ServiceResource{}
	_ resource.ResourceWithConfigure   = &ServiceResource{}
	_ resource.ResourceWithImportState = &ServiceResource{}
)

// NewServiceResource is a helper function to simplify the provider implementation.
func NewServiceResource() resource.Resource {
	return &ServiceResource{}
}

// ServiceResource defines the resource implementation.
type ServiceResource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the resource type name.
func (r *ServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service"
}

// Configure adds the provider configured client to the resource.
func (r *ServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *ServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	createParams := plan.ServiceParams(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new xMatters Service
	serviceReturn, err := r.client.PushService(createParams)
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("service", err))
		return
	}

	// Map response body to schema and populate Computed attribute values
	state := ServiceToModel(&resp.Diagnostics, serviceReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Created xMatters Service", map[string]any{"success": true})
}

// Read refreshes the Terraform state with the latest data.
func (r *ServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed Service from xMatters
	serviceReturn, err := r.client.GetService(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("service", err))
		return
	}

	// Capture state `lastUpdated` timestamp
	lastUpdated := state.LastUpdated
	// Overwrite Service with refreshed state
	state = ServiceToModel(&resp.Diagnostics, serviceReturn)
	if resp.Diagnostics.HasError() {
		return
	}

	// Ensure `last_updated` is preserved and set to the new state
	state.LastUpdated = lastUpdated

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Retrieved xMatters Service data", map[string]any{"success": true})
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *ServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ServiceModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	updateParams := plan.ServiceParams(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing xMatters Service and fetch updated Service
	serviceReturn, err := r.client.PushService(updateParams)
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("service", err))
		return
	}

	// Updated resource state with updated data and timestamp
	state := ServiceToModel(&resp.Diagnostics, serviceReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Save state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Updated xMatters Service", map[string]any{"success": true})
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *ServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceModel

	// Retrieve values from state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing Service
	err := r.client.DeleteService(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DeleteErrorDiagnostic("service", err))
		return
	}
}

// ImportState imports a resource from an existing xMatters resource.
func (r *ServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// ServiceToModel is a function that takes the API payload `xmatters.Service` and builds the state representation in the form of `ServiceModel`.
// The reverse of this method is `ServiceParams` which handles building an API representation using the proposed config.
func ServiceToModel(diags *diag.Diagnostics, service xmatters.Service) ServiceModel {
	model := ServiceModel{
		ID:           types.StringPointerValue(service.ID),
		Name:         customTypes.StringPointerValue(service.TargetName),
		Description:  customTypes.StringPointerValue(service.Description),
		Type:         customTypes.StringPointerValue(service.ServiceType),
		Tier:         types.StringPointerValue(service.ServiceTier),
		Owner:        utils.FlattenGroupReferenceID(service.OwnedBy),
		ServiceLinks: utils.FlattenServiceLinkSet(diags, service.ServiceLinks),
	}
	if diags.HasError() {
		return ServiceModel{}
	}
	return model
}
