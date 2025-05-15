package site

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/helpers"
	"github.com/xmatters/xmatters-go"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &SiteResource{}
	_ resource.ResourceWithConfigure   = &SiteResource{}
	_ resource.ResourceWithImportState = &SiteResource{}
)

// NewSiteResource is a helper function to simplify the provider implementation.
func NewSiteResource() resource.Resource {
	return &SiteResource{}
}

// SiteResource defines the resource implementation.
type SiteResource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the resource type name.
func (r *SiteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site"
}

// Configure adds the provider configured client to the resource.
func (r *SiteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *SiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SiteModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	createParams := plan.SiteParams()

	// Create new xMatters Site
	siteReturn, err := r.client.PushSite(createParams)
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("site", err))
		return
	}

	// Map response body to schema and populate Computed attribute values
	state := SiteToModel(siteReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Created xMatters Site", map[string]any{"success": true})
}

// Read refreshes the Terraform state with the latest data.
func (r *SiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SiteModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed Site from xMatters
	siteReturn, err := r.client.GetSite(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("site", err))
		return
	}

	// Capture state `lastUpdated` timestamp
	lastUpdated := state.LastUpdated
	// Overwrite Site with refreshed state
	state = SiteToModel(siteReturn)
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

	tflog.Info(ctx, "Retrieved xMatters Site data", map[string]any{"success": true})
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *SiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SiteModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	updateParams := plan.SiteParams()

	// Update existing xMatters Site and fetch updated Site
	siteReturn, err := r.client.PushSite(updateParams)
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("site", err))
		return
	}

	// Updated resource state with updated data and timestamp
	state := SiteToModel(siteReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Save state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Updated xMatters Site", map[string]any{"success": true})
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *SiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SiteModel

	// Retrieve values from state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing Site
	err := r.client.DeleteSite(state.ID.ValueStringPointer())
	if err != nil {
		resp.Diagnostics.Append(helpers.DeleteErrorDiagnostic("site", err))
		return
	}
}

// ImportState imports a resource from an existing xMatters resource.
func (r *SiteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// SiteToModel is a function that takes the API payload `xmatters.Site` and builds the state representation in the form of `SiteModel`.
// The reverse of this method is `SiteParams` which handles building an API representation using the proposed config.
func SiteToModel(site xmatters.Site) SiteModel {
	return SiteModel{
		Address1:   customTypes.StringPointerValue(site.Address1),
		Address2:   customTypes.StringPointerValue(site.Address2),
		City:       customTypes.StringPointerValue(site.City),
		Country:    customTypes.CountryPointerValue(site.Country),
		ID:         types.StringPointerValue(site.ID),
		Language:   types.StringPointerValue(site.Language),
		Latitude:   types.Float64PointerValue(site.Latitude),
		Longitude:  types.Float64PointerValue(site.Longitude),
		Name:       customTypes.StringPointerValue(site.Name),
		PostalCode: customTypes.StringPointerValue(site.PostalCode),
		State:      customTypes.StringPointerValue(site.State),
		Status:     types.StringPointerValue(site.Status),
		Timezone:   types.StringPointerValue(site.Timezone),
	}
}
