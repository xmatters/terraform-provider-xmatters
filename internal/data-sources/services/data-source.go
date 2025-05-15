package services

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
	_ datasource.DataSource              = &ServicesDataSource{}
	_ datasource.DataSourceWithConfigure = &ServicesDataSource{}
)

// NewServicesDataSource is a helper function to simplify the provider implementation.
func NewServicesDataSource() datasource.DataSource {
	return &ServicesDataSource{}
}

// ServicesDataSource defines the data-source implementation for a list of xMatters Services.
type ServicesDataSource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the data source type name.
func (d *ServicesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_services"
}

// Configure adds the provider configured client to the data source.
func (d *ServicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *ServicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ServicesModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the desired list of services from the xMatters API
	serviceReturn, err := d.client.GetServiceList(state.APIParams(&resp.Diagnostics))
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("services", err))
		return
	}

	// Transform the API response into a Terraform state
	newState := ServicesModel{
		Services: utils.FlattenServiceList(&resp.Diagnostics, serviceReturn),
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}
