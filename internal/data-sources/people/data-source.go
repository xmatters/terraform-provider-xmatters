package people

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
	_ datasource.DataSource              = &PeopleDataSource{}
	_ datasource.DataSourceWithConfigure = &PeopleDataSource{}
)

// NewPeopleDataSource is a helper function to simplify the provider implementation.
func NewPeopleDataSource() datasource.DataSource {
	return &PeopleDataSource{}
}

// PeopleDataSource defines the data-source implementation for a list of xMatters People.
type PeopleDataSource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the data source type name.
func (d *PeopleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_people"
}

// Configure adds the provider configured client to the data source.
func (d *PeopleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *PeopleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state PeopleModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the desired list of people from the xMatters API
	peopleReturn, err := d.client.GetPersonList(state.APIParams(&resp.Diagnostics))
	if resp.Diagnostics.HasError() {
		return
	}
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("people", err))
		return
	}

	// Transform the API response into a Terraform state
	newState := PeopleModel{
		People: utils.FlattenPersonList(&resp.Diagnostics, peopleReturn),
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}
