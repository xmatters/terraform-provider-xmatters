package group

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
	_ resource.Resource                = &GroupResource{}
	_ resource.ResourceWithConfigure   = &GroupResource{}
	_ resource.ResourceWithImportState = &GroupResource{}
)

// NewGroupResource is a helper function to simplify the provider implementation.
func NewGroupResource() resource.Resource {
	return &GroupResource{}
}

// GroupResource defines the resource implementation.
type GroupResource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the resource type name.
func (r *GroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

// Configure adds the provider configured client to the resource.
func (r *GroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *GroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan GroupModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	createParams := plan.GroupParams(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new xMatters Group
	groupReturn, err := r.client.PushGroup(createParams)
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("group", err))
		return
	}

	// Map response body to schema and populate Computed attribute values
	state := GroupToModel(&resp.Diagnostics, groupReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	utils.PreserveExplicitEmptyString(plan.GroupType, &state.GroupType)
	utils.PreserveExplicitEmptyString(plan.ExternalKey.StringValue, &state.ExternalKey.StringValue)
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Created xMatters Group", map[string]any{"success": true})
}

// Read refreshes the Terraform state with the latest data.
func (r *GroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GroupModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed Group from xMatters
	groupReturn, err := r.client.GetGroup(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("group", err))
		return
	}

	// Capture state `lastUpdated` timestamp
	priorState := state
	lastUpdated := state.LastUpdated
	// Overwrite Group with refreshed state
	state = GroupToModel(&resp.Diagnostics, groupReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	utils.PreserveExplicitEmptyString(priorState.GroupType, &state.GroupType)
	utils.PreserveExplicitEmptyString(priorState.ExternalKey.StringValue, &state.ExternalKey.StringValue)
	// Ensure `last_updated` is preserved and set to the new state
	state.LastUpdated = lastUpdated

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Retrieved xMatters Group data", map[string]any{"success": true})
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *GroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan GroupModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	updateParams := plan.GroupParams(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing xMatters Group and fetch updated Group
	groupReturn, err := r.client.PushGroup(updateParams)
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("group", err))
		return
	}

	// Updated resource state with updated data and timestamp
	state := GroupToModel(&resp.Diagnostics, groupReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	utils.PreserveExplicitEmptyString(plan.GroupType, &state.GroupType)
	utils.PreserveExplicitEmptyString(plan.ExternalKey.StringValue, &state.ExternalKey.StringValue)
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Save state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Updated xMatters Group", map[string]any{"success": true})
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *GroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state GroupModel

	// Retrieve values from state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing Group
	err := r.client.DeleteGroup(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DeleteErrorDiagnostic("group", err))
		return
	}
}

// ImportState imports a resource from an existing xMatters resource.
func (r *GroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// GroupToModel is a function that takes the API payload `xmatters.Group` and builds the state representation in the form of `GroupModel`.
// The reverse of this method is `GroupParams` which handles building an API representation using the proposed config.
func GroupToModel(diags *diag.Diagnostics, group xmatters.Group) GroupModel {
	model := GroupModel{
		ID:                types.StringPointerValue(group.ID),
		Name:              customTypes.StringPointerValue(group.TargetName),
		Status:            types.StringPointerValue(group.Status),
		Description:       customTypes.StringPointerValue(group.Description),
		GroupType:         types.StringPointerValue(group.GroupType),
		AllowDuplicates:   types.BoolPointerValue(group.AllowDuplicates),
		Site:              utils.FlattenReferenceID(group.Site),
		ObservedByAll:     types.BoolPointerValue(group.ObservedByAll),
		Observers:         utils.FlattenReferenceNameSet(diags, group.Observers),
		UseDefaultDevices: types.BoolPointerValue(group.UseDefaultDevices),
		Supervisors:       utils.FlattenReferenceIDSet(diags, group.Supervisors),
		ExternalKey:       customTypes.StringPointerValue(group.ExternalKey),
		ExternallyOwned:   types.BoolPointerValue(group.ExternallyOwned),
		Criteria:          utils.FlattenGroupCriteriaObject(diags, group.Criteria),
	}
	if diags.HasError() {
		return GroupModel{}
	}
	return model
}
