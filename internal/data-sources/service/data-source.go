package service

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/helpers"
	"github.com/xmatters/xmatters-go"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &ServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &ServiceDataSource{}
)

// NewServiceDataSource is a helper function to simplify the provider implementation.
func NewServiceDataSource() datasource.DataSource {
	return &ServiceDataSource{}
}

// ServiceDataSource defines the data-source implementation for a list of xMatters Service.
type ServiceDataSource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the data source type name.
func (d *ServiceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service"
}

// Configure adds the provider configured client to the data source.
func (d *ServiceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *ServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ServiceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the desired list of service from the xMatters API
	serviceReturn, err := d.client.GetService(state.ServiceID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("service", err))
		return
	}

	// Transform the API response into a Terraform state
	newState := ServiceToModel(&resp.Diagnostics, serviceReturn)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

// ServiceToModel is a function that takes the API payload `xmatters.Service` and builds the state representation in the form of `ServiceModel`.
func ServiceToModel(diags *diag.Diagnostics, service xmatters.Service) ServiceModel {
	model := ServiceModel{
		ID:           types.StringPointerValue(service.ID),
		Name:         types.StringPointerValue(service.TargetName),
		Description:  types.StringPointerValue(service.Description),
		Type:         types.StringPointerValue(service.ServiceType),
		Tier:         types.StringPointerValue(service.ServiceTier),
		Owner:        utils.FlattenGroupReferenceID(service.OwnedBy),
		ServiceLinks: utils.FlattenServiceLinkSet(diags, service.ServiceLinks),
	}
	return model
}
