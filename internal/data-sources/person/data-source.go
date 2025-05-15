package person

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
	_ datasource.DataSource              = &PersonDataSource{}
	_ datasource.DataSourceWithConfigure = &PersonDataSource{}
)

// NewPersonDataSource is a helper function to simplify the provider implementation.
func NewPersonDataSource() datasource.DataSource {
	return &PersonDataSource{}
}

// PersonDataSource defines the data-source implementation for a list of xMatters Person.
type PersonDataSource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the data source type name.
func (d *PersonDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_person"
}

// Configure adds the provider configured client to the data source.
func (d *PersonDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *PersonDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state PersonModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the desired list of person from the xMatters API
	personReturn, err := d.client.GetPerson(state.PersonID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("person", err))
		return
	}

	// Transform the API response into a Terraform state
	newState := PersonToModel(&resp.Diagnostics, personReturn)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

// PersonToModel is a function that takes the API payload `xmatters.Person` and builds the state representation in the form of `PersonModel`.
func PersonToModel(diags *diag.Diagnostics, person xmatters.Person) PersonModel {
	model := PersonModel{
		ID:              types.StringPointerValue(person.ID),
		TargetName:      types.StringPointerValue(person.TargetName),
		FirstName:       types.StringPointerValue(person.FirstName),
		LastName:        types.StringPointerValue(person.LastName),
		Roles:           utils.FlattenRoleSet(diags, person.Roles),
		Status:          types.StringPointerValue(person.Status),
		WebLogin:        types.StringPointerValue(person.WebLogin),
		Site:            utils.FlattenReferenceID(person.Site),
		Timezone:        types.StringPointerValue(person.Timezone),
		Language:        types.StringPointerValue(person.Language),
		Supervisors:     utils.FlattenSupervisorSet(diags, person.Supervisors),
		PhoneLogin:      types.StringPointerValue(person.PhoneLogin),
		LicenseType:     types.StringPointerValue(person.LicenseType),
		ExternalKey:     types.StringPointerValue(person.ExternalKey),
		ExternallyOwned: types.BoolPointerValue(person.ExternallyOwned),
		LastLogin:       types.StringPointerValue(person.LastLogin),
	}
	if diags.HasError() {
		return PersonModel{}
	}
	return model
}
