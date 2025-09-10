package device

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DeviceModel represents an xMatters Device object in the Provider.
type DeviceModel struct {
	DeviceID          types.String `tfsdk:"device_id" tf:"required"`
	ID                types.String `tfsdk:"id"`
	TargetName        types.String `tfsdk:"target_name"`
	Country           types.String `tfsdk:"country"`
	DefaultDevice     types.Bool   `tfsdk:"default_device"`
	Delay             types.Int32  `tfsdk:"delay"`
	DeviceType        types.String `tfsdk:"device_type"`
	EmailAddress      types.String `tfsdk:"email_address"`
	ExternalKey       types.String `tfsdk:"external_key"`
	ExternallyOwned   types.Bool   `tfsdk:"externally_owned"`
	Name              types.String `tfsdk:"name"`
	Owner             types.String `tfsdk:"owner"`
	PhoneNumber       types.String `tfsdk:"phone_number"`
	PIN               types.String `tfsdk:"pin"`
	PriorityThreshold types.String `tfsdk:"priority_threshold"`
	Sequence          types.Int32  `tfsdk:"sequence"`
	Status            types.String `tfsdk:"status"`
	TestStatus        types.String `tfsdk:"test_status"`
	Timeframes        types.Set    `tfsdk:"timeframes"`
	TwoWayDevice      types.Bool   `tfsdk:"two_way_device"`
}
