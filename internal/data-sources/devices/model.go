package devices

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

// DevicessModel contains the search fields and return values for the Provider's Devicess data source.
type DevicesModel struct {
	Filters *DevicesFilterModel `tfsdk:"filters" tf:"optional"`
	Devices types.List          `tfsdk:"devices" tf:"computed"`
}

// DevicessSearchModel contains the search fields for the Provider's Devicess data source.
type DevicesFilterModel struct {
	DeviceStatus types.String `tfsdk:"device_status" tf:"optional"`
	DeviceType   types.List   `tfsdk:"device_type" tf:"optional"`
	DeviceNames  types.List   `tfsdk:"device_names" tf:"optional"`
}

// APIParams returns the xmatters.GetDevicessParams object based on the DevicessModel instance.
func (in DevicesModel) APIParams(diags *diag.Diagnostics) xmatters.GetDevicesParams {
	search := xmatters.GetDevicesParams{
		Embed: "timeframes",
	}
	if in.Filters != nil {
		search.DeviceStatus = in.Filters.DeviceStatus.ValueString()
		search.DeviceType = utils.ExpandStringList(diags, in.Filters.DeviceType)
		search.DeviceNames = utils.ExpandEncodedStringList(diags, in.Filters.DeviceNames) // Due to the way the API is designed, we need to URL-encode the device names prior to sending the request
	}
	return search
}
