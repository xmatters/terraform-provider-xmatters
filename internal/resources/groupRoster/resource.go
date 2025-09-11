package groupRoster

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
	_ resource.Resource                = &GroupRosterResource{}
	_ resource.ResourceWithConfigure   = &GroupRosterResource{}
	_ resource.ResourceWithImportState = &GroupRosterResource{}
)

// NewGroupRosterResource is a helper function to simplify the provider implementation.
func NewGroupRosterResource() resource.Resource {
	return &GroupRosterResource{}
}

// GroupRosterResource defines the resource implementation.
type GroupRosterResource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the resource type name.
func (r *GroupRosterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group_roster"
}

// Configure adds the provider configured client to the resource.
func (r *GroupRosterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *GroupRosterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan GroupRosterModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new xMatters GroupRoster
	groupRosterReturn, err := r.client.PushGroupRoster(plan.ID.ValueString(), utils.ExpandGroupMemberSet(&resp.Diagnostics, plan.Members))
	if resp.Diagnostics.HasError() {
		return
	}
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("group_roster", err))
		return
	}

	// Map response body to schema and populate Computed attribute values
	state := GroupRosterToModel(&resp.Diagnostics, groupRosterReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Created xMatters Group Roster", map[string]any{"success": true})
}

// Read refreshes the Terraform state with the latest data.
func (r *GroupRosterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GroupRosterModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed GroupRoster from xMatters
	groupRosterReturn, err := r.client.GetGroupRoster(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("group_roster", err))
		return
	}

	// Capture state `lastUpdated` timestamp
	lastUpdated := state.LastUpdated
	// Overwrite GroupRoster with refreshed state
	state = GroupRosterToModel(&resp.Diagnostics, groupRosterReturn)
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

	tflog.Info(ctx, "Retrieved xMatters Group Roster data", map[string]any{"success": true})
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *GroupRosterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan GroupRosterModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing xMatters GroupRoster and fetch updated GroupRoster
	groupRosterReturn, err := r.client.PushGroupRoster(plan.ID.ValueString(), utils.ExpandGroupMemberSet(&resp.Diagnostics, plan.Members))
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("group_roster", err))
		return
	}

	// Updated resource state with updated data and timestamp
	state := GroupRosterToModel(&resp.Diagnostics, groupRosterReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Save state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Updated xMatters Group Roster", map[string]any{"success": true})
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *GroupRosterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state GroupRosterModel

	// Retrieve values from state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing GroupRoster
	err := r.client.DeleteGroupRoster(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DeleteErrorDiagnostic("group_roster", err))
		return
	}
}

// ImportState imports a resource from an existing xMatters resource.
func (r *GroupRosterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// GroupRosterToModel is a function that takes the API payload `xmatters.GroupRoster` and builds the state representation in the form of `GroupRosterModel`.
// The reverse of this method is `GroupRosterParams` which handles building an API representation using the proposed config.
func GroupRosterToModel(diags *diag.Diagnostics, groupRoster xmatters.GroupRoster) GroupRosterModel {
	model := GroupRosterModel{
		ID:      types.StringPointerValue(groupRoster.ID),
		Group:   utils.FlattenGroupReferenceObject(diags, groupRoster.Group),
		Members: utils.FlattenGroupMemberSet(diags, groupRoster.Members),
	}
	if diags.HasError() {
		return GroupRosterModel{}
	}
	return model
}
