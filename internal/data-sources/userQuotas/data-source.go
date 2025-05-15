package userQuotas

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
	_ datasource.DataSource              = &UserQuotasDataSource{}
	_ datasource.DataSourceWithConfigure = &UserQuotasDataSource{}
)

// NewUserQuotasDataSource is a helper function to simplify the provider implementation.
func NewUserQuotasDataSource() datasource.DataSource {
	return &UserQuotasDataSource{}
}

// UserQuotasDataSource defines the data-source implementation for a list of xMatters User License Quota.
type UserQuotasDataSource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the data source type name.
func (d *UserQuotasDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_license_quotas"
}

// Configure adds the provider configured client to the data source.
func (d *UserQuotasDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *UserQuotasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state UserQuotasModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the desired list of user_license_quotas from the xMatters API
	quotasReturn, err := d.client.GetUserQuotas()
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("user license quotas", err))
		return
	}

	// Transform the API response into a Terraform state
	newState := UserQuotasToModel(&resp.Diagnostics, quotasReturn)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

// UserQuotasToModel is a function that takes the API payload `xmatters.UserQuota` and builds the state representation in the form of `UserQuotasModel`.
func UserQuotasToModel(diags *diag.Diagnostics, quotas xmatters.UserQuotas) UserQuotasModel {
	model := UserQuotasModel{
		StakeholderUsersEnabled: types.BoolPointerValue(quotas.StakeholderUsersEnabled),
		StakeholderUsers:        utils.FlattenQuotaDetailsObject(diags, quotas.StakeholderUsers),
		FullUsers:               utils.FlattenQuotaDetailsObject(diags, quotas.FullUsers),
	}
	return model
}
