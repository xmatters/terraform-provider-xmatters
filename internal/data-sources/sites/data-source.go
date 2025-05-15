package sites

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/helpers"
	"github.com/xmatters/xmatters-go"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &SitesDataSource{}
	_ datasource.DataSourceWithConfigure = &SitesDataSource{}
)

// NewSitesDataSource is a helper function to simplify the provider implementation.
func NewSitesDataSource() datasource.DataSource {
	return &SitesDataSource{}
}

// SitesDataSource defines the data-source implementation for a list of xMatters Sites.
type SitesDataSource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the data source type name.
func (d *SitesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sites"
}

// Configure adds the provider configured client to the data source.
func (d *SitesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *SitesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SitesModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the desired list of sites from the xMatters API
	serviceReturn, err := d.client.GetSiteList(state.APIParams(&resp.Diagnostics))
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("sites", err))
		return
	}

	// Transform the API response into a Terraform state
	newState := SitesModel{
		Sites: utils.FlattenSiteList(&resp.Diagnostics, serviceReturn),
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}
