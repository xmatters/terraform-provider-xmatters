package group

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
	_ datasource.DataSource              = &GroupDataSource{}
	_ datasource.DataSourceWithConfigure = &GroupDataSource{}
)

// NewGroupDataSource is a helper function to simplify the provider implementation.
func NewGroupDataSource() datasource.DataSource {
	return &GroupDataSource{}
}

// GroupDataSource defines the data-source implementation for a list of xMatters Group.
type GroupDataSource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the data source type name.
func (d *GroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

// Configure adds the provider configured client to the data source.
func (d *GroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *GroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state GroupModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the desired list of group from the xMatters API
	groupReturn, err := d.client.GetGroup(state.GroupID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("group", err))
		return
	}

	// Transform the API response into a Terraform state
	newState := GroupToModel(&resp.Diagnostics, groupReturn)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

// GroupToModel is a function that takes the API payload `xmatters.Group` and builds the state representation in the form of `GroupModel`.
func GroupToModel(diags *diag.Diagnostics, group xmatters.Group) GroupModel {
	model := GroupModel{
		ID:              types.StringPointerValue(group.ID),
		TargetName:      types.StringPointerValue(group.TargetName),
		Description:     types.StringPointerValue(group.Description),
		Status:          types.StringPointerValue(group.Status),
		ExternalKey:     types.StringPointerValue(group.ExternalKey),
		ExternallyOwned: types.BoolPointerValue(group.ExternallyOwned),
		AllowDuplicates: types.BoolPointerValue(group.AllowDuplicates),
		Timezone:        types.StringPointerValue(group.Timezone),
		Site:            utils.FlattenReferenceID(group.Site),
		ObservedByAll:   types.BoolPointerValue(group.ObservedByAll),
		Observers:       utils.FlattenReferenceNameSet(diags, group.Observers),
		Supervisors:     utils.FlattenReferenceIDSet(diags, group.Supervisors),
		Services:        utils.FlattenServiceSet(diags, group.Services),
		GroupType:       types.StringPointerValue(group.GroupType),
	}
	return model
}
