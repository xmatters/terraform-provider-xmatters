package device

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
	_ datasource.DataSource              = &DeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceDataSource{}
)

// NewDeviceDataSource is a helper function to simplify the provider implementation.
func NewDeviceDataSource() datasource.DataSource {
	return &DeviceDataSource{}
}

// DeviceDataSource defines the data-source implementation for a list of xMatters Device.
type DeviceDataSource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the data source type name.
func (d *DeviceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

// Configure adds the provider configured client to the data source.
func (d *DeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *DeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state DeviceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the desired list of device from the xMatters API
	deviceReturn, err := d.client.GetDevice(state.DeviceID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("device", err))
		return
	}

	// Transform the API response into a Terraform state
	newState := DeviceToModel(&resp.Diagnostics, deviceReturn)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

// DeviceToModel is a function that takes the API payload `xmatters.Device` and builds the state representation in the form of `DeviceModel`.
func DeviceToModel(diags *diag.Diagnostics, device xmatters.Device) DeviceModel {
	model := DeviceModel{
		ID:                types.StringPointerValue(device.ID),
		TargetName:        types.StringPointerValue(device.TargetName),
		Country:           types.StringPointerValue(device.Country),
		DefaultDevice:     types.BoolPointerValue(device.DefaultDevice),
		Delay:             types.Int32PointerValue(device.Delay),
		DeviceType:        types.StringPointerValue(device.DeviceType),
		EmailAddress:      types.StringPointerValue(device.EmailAddress),
		ExternalKey:       types.StringPointerValue(device.ExternalKey),
		ExternallyOwned:   types.BoolPointerValue(device.ExternallyOwned),
		Name:              types.StringPointerValue(device.Name),
		Owner:             utils.FlattenPersonReferenceID(device.Owner),
		PhoneNumber:       types.StringPointerValue(device.PhoneNumber),
		PIN:               types.StringPointerValue(device.PIN),
		PriorityThreshold: types.StringPointerValue(device.PriorityThreshold),
		Sequence:          types.Int32PointerValue(device.Sequence),
		Status:            types.StringPointerValue(device.Status),
		TestStatus:        types.StringPointerValue(device.TestStatus),
		Timeframes:        utils.FlattenTimeframeSet(diags, device.Timeframes),
		TwoWayDevice:      types.BoolPointerValue(device.TwoWayDevice),
	}
	return model
}
