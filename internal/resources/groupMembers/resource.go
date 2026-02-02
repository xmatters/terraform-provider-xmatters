package groupMembers

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
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/helpers"
	"github.com/xmatters/xmatters-go"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &GroupMembersResource{}
	_ resource.ResourceWithConfigure   = &GroupMembersResource{}
	_ resource.ResourceWithImportState = &GroupMembersResource{}
)

// NewGroupMembersResource is a helper function to simplify the provider implementation.
func NewGroupMembersResource() resource.Resource {
	return &GroupMembersResource{}
}

// GroupMembersResource defines the resource implementation.
type GroupMembersResource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the resource type name.
func (r *GroupMembersResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group_members"
}

// Configure adds the provider configured client to the resource.
func (r *GroupMembersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *GroupMembersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan GroupMembersModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new xMatters GroupMembers
	groupMembersReturn, err := r.client.PushGroupMembers(plan.ID.ValueString(), utils.ExpandGroupMemberSet(&resp.Diagnostics, plan.Members))
	if resp.Diagnostics.HasError() {
		return
	}
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("group_members", err))
		return
	}

	// Map response body to schema and populate Computed attribute values
	state := GroupMembersToModel(&resp.Diagnostics, groupMembersReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Created xMatters Group Members", map[string]any{"success": true})
}

// Read refreshes the Terraform state with the latest data.
func (r *GroupMembersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GroupMembersModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed GroupMembers from xMatters
	groupMembersReturn, err := r.client.GetGroupMembers(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("group_members", err))
		return
	}

	// Capture state `lastUpdated` timestamp
	lastUpdated := state.LastUpdated
	// Overwrite GroupMembers with refreshed state
	state = GroupMembersToModel(&resp.Diagnostics, groupMembersReturn)
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

	tflog.Info(ctx, "Retrieved xMatters Group Members data", map[string]any{"success": true})
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *GroupMembersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan GroupMembersModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing xMatters GroupMembers and fetch updated GroupMembers
	groupMembersReturn, err := r.client.PushGroupMembers(plan.ID.ValueString(), utils.ExpandGroupMemberSet(&resp.Diagnostics, plan.Members))
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("group_members", err))
		return
	}

	// Updated resource state with updated data and timestamp
	state := GroupMembersToModel(&resp.Diagnostics, groupMembersReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Save state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Updated xMatters Group Members", map[string]any{"success": true})
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *GroupMembersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state GroupMembersModel

	// Retrieve values from state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing GroupMembers
	err := r.client.DeleteGroupMembers(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DeleteErrorDiagnostic("group_members", err))
		return
	}
}

// ImportState imports a resource from an existing xMatters resource.
func (r *GroupMembersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// GroupMembersToModel is a function that takes the API payload `xmatters.GroupMembers` and builds the state representation in the form of `GroupMembersModel`.
// The reverse of this method is `GroupMembersParams` which handles building an API representation using the proposed config.
func GroupMembersToModel(diags *diag.Diagnostics, groupMembers xmatters.GroupMembers) GroupMembersModel {
	model := GroupMembersModel{
		ID:      types.StringPointerValue(groupMembers.ID),
		Group:   utils.FlattenGroupReferenceObject(diags, groupMembers.Group),
		Members: utils.FlattenGroupMemberSet(diags, groupMembers.Members),
	}
	if diags.HasError() {
		return GroupMembersModel{}
	}
	return model
}
