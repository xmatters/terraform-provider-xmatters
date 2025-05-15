package person

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
	_ resource.Resource                = &PersonResource{}
	_ resource.ResourceWithConfigure   = &PersonResource{}
	_ resource.ResourceWithImportState = &PersonResource{}
)

// NewPersonResource is a helper function to simplify the provider implementation.
func NewPersonResource() resource.Resource {
	return &PersonResource{}
}

// PersonResource defines the resource implementation.
type PersonResource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the resource type name.
func (r *PersonResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_person"
}

// Configure adds the provider configured client to the resource.
func (r *PersonResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *PersonResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PersonModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	createParams := plan.PersonParams(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new xMatters Person
	personReturn, err := r.client.PushPerson(createParams)
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("person", err))
		return
	}

	// Map response body to schema and populate Computed attribute values
	state := PersonToModel(&resp.Diagnostics, personReturn, plan)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Created xMatters Person", map[string]any{"success": true})
}

// Read refreshes the Terraform state with the latest data.
func (r *PersonResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PersonModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed Person from xMatters
	personReturn, err := r.client.GetPerson(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("person", err))
		return
	}

	// Capture state `lastUpdated` timestamp
	lastUpdated := state.LastUpdated
	// Overwrite Person with refreshed state
	state = PersonToModel(&resp.Diagnostics, personReturn, state)
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

	tflog.Info(ctx, "Retrieved xMatters Person data", map[string]any{"success": true})
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *PersonResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PersonModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	updateParams := plan.PersonParams(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing xMatters Person and fetch updated Person
	personReturn, err := r.client.PushPerson(updateParams)
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("person", err))
		return
	}

	// Updated resource state with updated data and timestamp
	state := PersonToModel(&resp.Diagnostics, personReturn, plan)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Save state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Updated xMatters Person", map[string]any{"success": true})
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *PersonResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PersonModel

	// Retrieve values from state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing Person
	err := r.client.DeletePerson(state.ID.ValueStringPointer())
	if err != nil {
		resp.Diagnostics.Append(helpers.DeleteErrorDiagnostic("person", err))
		return
	}
}

// ImportState imports a resource from an existing xMatters resource.
func (r *PersonResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// PersonToModel is a function that takes the API payload `xmatters.Person` and builds the state representation in the form of `PersonModel`.
// The reverse of this method is `PersonParams` which handles building an API representation using the proposed config.
func PersonToModel(diags *diag.Diagnostics, person xmatters.Person, plan PersonModel) PersonModel {
	model := PersonModel{
		ID:              types.StringPointerValue(person.ID),
		TargetName:      customTypes.StringPointerValue(person.TargetName),
		FirstName:       customTypes.StringPointerValue(person.FirstName),
		LastName:        customTypes.StringPointerValue(person.LastName),
		Roles:           plan.Roles,
		Status:          types.StringPointerValue(person.Status),
		WebLogin:        customTypes.StringPointerValue(person.WebLogin),
		Site:            utils.FlattenReferenceID(person.Site),
		Timezone:        types.StringPointerValue(person.Timezone),
		Language:        types.StringPointerValue(person.Language),
		Supervisors:     utils.FlattenSupervisorSet(diags, person.Supervisors),
		PhoneLogin:      types.StringPointerValue(person.PhoneLogin),
		PhonePin:        plan.PhonePin,
		LicenseType:     customTypes.StringPointerValue(person.LicenseType),
		ExternalKey:     customTypes.StringPointerValue(person.ExternalKey),
		ExternallyOwned: types.BoolPointerValue(person.ExternallyOwned),
	}
	if person.Roles != nil {
		model.Roles = utils.FlattenRoleSet(diags, person.Roles)
	}
	if diags.HasError() {
		return PersonModel{}
	}
	return model
}
