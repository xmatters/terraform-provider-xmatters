package site

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/xmatters/terraform-provider-xmatters/internal/utils/helpers"
	"github.com/xmatters/xmatters-go"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &SiteDataSource{}
	_ datasource.DataSourceWithConfigure = &SiteDataSource{}
)

// NewSiteDataSource is a helper function to simplify the provider implementation.
func NewSiteDataSource() datasource.DataSource {
	return &SiteDataSource{}
}

// SiteDataSource defines the data-source implementation for a list of xMatters Site.
type SiteDataSource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the data source type name.
func (d *SiteDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site"
}

// Configure adds the provider configured client to the data source.
func (d *SiteDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*xmatters.XMattersAPI)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *xmatters.XMattersAPI, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

// Read refreshes the Terraform state with the latest data.
func (d *SiteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SiteModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the desired list of site from the xMatters API
	siteReturn, err := d.client.GetSite(state.SiteID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("site", err))
		return
	}

	// Transform the API response into a Terraform state
	newState := SiteToModel(siteReturn)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

// SiteToModel is a function that takes the API payload `xmatters.Site` and builds the state representation in the form of `SiteModel`.
func SiteToModel(site xmatters.Site) SiteModel {
	model := SiteModel{
		ID:         types.StringPointerValue(site.ID),
		Address1:   types.StringPointerValue(site.Address1),
		Address2:   types.StringPointerValue(site.Address2),
		City:       types.StringPointerValue(site.City),
		Country:    types.StringPointerValue(site.Country),
		Language:   types.StringPointerValue(site.Language),
		Latitude:   types.Float64PointerValue(site.Latitude),
		Longitude:  types.Float64PointerValue(site.Longitude),
		Name:       types.StringPointerValue(site.Name),
		PostalCode: types.StringPointerValue(site.PostalCode),
		State:      types.StringPointerValue(site.State),
		Status:     types.StringPointerValue(site.Status),
		Timezone:   types.StringPointerValue(site.Timezone),
	}
	return model
}
