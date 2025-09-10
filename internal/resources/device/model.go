package device

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

// DeviceModel represents an xMatters Device object in the Provider.
type DeviceModel struct {
	ID                types.String                  `tfsdk:"id"`
	TargetName        types.String                  `tfsdk:"target_name"`
	DefaultDevice     types.Bool                    `tfsdk:"default_device"`
	Delay             types.Int32                   `tfsdk:"delay"`
	DeviceType        types.String                  `tfsdk:"device_type"`
	EmailAddress      customTypes.CustomStringValue `tfsdk:"email_address"`
	ExternalKey       customTypes.CustomStringValue `tfsdk:"external_key"`
	ExternallyOwned   types.Bool                    `tfsdk:"externally_owned"`
	Name              customTypes.CustomStringValue `tfsdk:"name"`
	Owner             types.String                  `tfsdk:"owner"`
	PhoneNumber       types.String                  `tfsdk:"phone_number"`
	PIN               customTypes.CustomStringValue `tfsdk:"pin"`
	PriorityThreshold customTypes.CustomStringValue `tfsdk:"priority_threshold"`
	Status            types.String                  `tfsdk:"status"`
	TestStatus        customTypes.CustomStringValue `tfsdk:"test_status"`
	Timeframes        types.Set                     `tfsdk:"timeframes"`
	TwoWayDevice      types.Bool                    `tfsdk:"two_way_device"`
	LastUpdated       types.String                  `tfsdk:"last_updated"`
}

// DeviceParams is a method that takes the proposed configuration changes `DeviceModel` and builds the API representation in the form of `*xmatters.PushDeviceParams`.
// The reverse of this method is `DeviceToModel` which handles building a state representation using the API response.
func (in DeviceModel) DeviceParams(diags *diag.Diagnostics) xmatters.PushDeviceParams {
	return xmatters.PushDeviceParams{
		ID:                in.ID.ValueString(),
		Delay:             in.Delay.ValueInt32Pointer(),
		DeviceType:        in.DeviceType.ValueString(),
		EmailAddress:      in.EmailAddress.ValueString(),
		ExternalKey:       in.ExternalKey.ValueStringPointer(),
		ExternallyOwned:   in.ExternallyOwned.ValueBoolPointer(),
		Name:              in.Name.ValueString(),
		Owner:             in.Owner.ValueString(),
		PhoneNumber:       in.PhoneNumber.ValueString(),
		PIN:               in.PIN.ValueString(),
		PriorityThreshold: in.PriorityThreshold.ValueString(),
		Status:            in.Status.ValueString(),
		TestStatus:        in.TestStatus.ValueString(),
		Timeframes:        utils.ExpandTimeframeSet(diags, in.Timeframes),
		DefaultDevice:     in.DefaultDevice.ValueBoolPointer(),
		TwoWayDevice:      in.TwoWayDevice.ValueBoolPointer(),
	}
}
