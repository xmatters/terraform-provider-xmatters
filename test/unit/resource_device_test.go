package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/resources/device"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

func TestResourceDeviceToModel(t *testing.T) {
	testDevice := xmatters.Device{
		ID:              utils.RandUUIDPointer(),
		TargetName:      utils.RandStringPointer(5),
		DefaultDevice:   utils.RandBoolPointer(),
		Delay:           utils.RandInt32Pointer(),
		DeviceType:      utils.RandStringPointer(5),
		EmailAddress:    utils.RandStringPointer(5),
		ExternalKey:     utils.RandStringPointer(5),
		ExternallyOwned: utils.RandBoolPointer(),
		Name:            utils.RandStringPointer(5),
		Owner: &xmatters.PersonReference{
			ID: utils.RandUUIDPointer(),
		},
		PhoneNumber:       utils.RandStringPointer(5),
		PIN:               utils.RandStringPointer(5),
		PriorityThreshold: utils.RandStringPointer(5),
		Status:            utils.RandStringPointer(5),
		TestStatus:        utils.RandStringPointer(5),
		Timeframes: []*xmatters.DeviceTimeframe{
			{
				Name:              utils.RandStringPointer(5),
				StartTime:         utils.RandStringPointer(5),
				DurationInMinutes: utils.RandInt32Pointer(),
				Days:              utils.RandStringPointerList(2, 5),
				ExcludeHolidays:   utils.RandBoolPointer(),
			},
		},
		TwoWayDevice: utils.RandBoolPointer(),
	}
	type args struct {
		diags *diag.Diagnostics
		in    xmatters.Device
	}
	tests := []struct {
		name     string
		args     args
		expected device.DeviceModel
	}{
		{
			name: "empty device",
			args: args{
				diags: &diag.Diagnostics{},
				in:    xmatters.Device{},
			},
			expected: device.DeviceModel{
				Timeframes: types.SetValueMust(
					utils.TimeframeObjectType,
					[]attr.Value{},
				),
			},
		},
		{
			name: "valid device",
			args: args{
				diags: &diag.Diagnostics{},
				in:    testDevice,
			},
			expected: device.DeviceModel{
				ID:                types.StringPointerValue(testDevice.ID),
				TargetName:        types.StringPointerValue(testDevice.TargetName),
				DefaultDevice:     types.BoolPointerValue(testDevice.DefaultDevice),
				Delay:             types.Int32PointerValue(testDevice.Delay),
				DeviceType:        types.StringPointerValue(testDevice.DeviceType),
				EmailAddress:      customTypes.StringPointerValue(testDevice.EmailAddress),
				ExternalKey:       customTypes.StringPointerValue(testDevice.ExternalKey),
				ExternallyOwned:   types.BoolPointerValue(testDevice.ExternallyOwned),
				Name:              customTypes.StringPointerValue(testDevice.Name),
				Owner:             types.StringPointerValue(testDevice.Owner.ID),
				PhoneNumber:       types.StringPointerValue(testDevice.PhoneNumber),
				PIN:               customTypes.StringPointerValue(testDevice.PIN),
				PriorityThreshold: customTypes.StringPointerValue(testDevice.PriorityThreshold),
				Status:            types.StringPointerValue(testDevice.Status),
				TestStatus:        customTypes.StringPointerValue(testDevice.TestStatus),
				Timeframes: types.SetValueMust(utils.TimeframeObjectType,
					[]attr.Value{
						types.ObjectValueMust(utils.TimeframeObjectType.AttrTypes,
							map[string]attr.Value{
								"name":                customTypes.StringPointerValue(testDevice.Timeframes[0].Name),
								"start_time":          types.StringPointerValue(testDevice.Timeframes[0].StartTime),
								"duration_in_minutes": types.Int32PointerValue(testDevice.Timeframes[0].DurationInMinutes),
								"days": types.SetValueMust(customTypes.CustomStringType{},
									[]attr.Value{
										customTypes.StringPointerValue(testDevice.Timeframes[0].Days[0]),
										customTypes.StringPointerValue(testDevice.Timeframes[0].Days[1]),
									},
								),
								"exclude_holidays": types.BoolPointerValue(testDevice.Timeframes[0].ExcludeHolidays),
							},
						),
					},
				),
				TwoWayDevice: types.BoolPointerValue(testDevice.TwoWayDevice),
			},
		},
	}
	for _, thisTest := range tests {
		actual := device.DeviceToModel(thisTest.args.diags, thisTest.args.in)
		assert.Equal(t, thisTest.expected, actual)
	}
}

func TestDeviceParams(t *testing.T) {
	testDeviceModel := xmatters.PushDeviceParams{
		DefaultDevice:     utils.RandBoolPointer(),
		Delay:             utils.RandInt32Pointer(),
		DeviceType:        utils.RandString(5),
		EmailAddress:      utils.RandString(5),
		ExternalKey:       utils.RandStringPointer(5),
		ExternallyOwned:   utils.RandBoolPointer(),
		Name:              utils.RandString(5),
		Owner:             utils.RandString(5),
		PhoneNumber:       utils.RandString(5),
		PIN:               utils.RandString(5),
		PriorityThreshold: utils.RandString(5),
		Status:            utils.RandString(5),
		TestStatus:        utils.RandString(5),
		Timeframes: []*xmatters.DeviceTimeframe{
			{
				Name:              utils.RandStringPointer(5),
				StartTime:         utils.RandStringPointer(5),
				DurationInMinutes: utils.RandInt32Pointer(),
				Days:              utils.RandStringPointerList(2, 5),
				ExcludeHolidays:   utils.RandBoolPointer(),
			},
		},
		TwoWayDevice: utils.RandBoolPointer(),
	}
	type args struct {
		diags *diag.Diagnostics
		in    device.DeviceModel
	}
	tests := []struct {
		name     string
		args     args
		expected xmatters.PushDeviceParams
	}{
		{
			name: "empty device",
			args: args{
				diags: &diag.Diagnostics{},
				in:    device.DeviceModel{},
			},
			expected: xmatters.PushDeviceParams{},
		},
		{
			name: "valid device",
			args: args{
				diags: &diag.Diagnostics{},
				in: device.DeviceModel{
					DefaultDevice:     types.BoolPointerValue(testDeviceModel.DefaultDevice),
					Delay:             types.Int32PointerValue(testDeviceModel.Delay),
					DeviceType:        types.StringValue(testDeviceModel.DeviceType),
					EmailAddress:      customTypes.StringValue(testDeviceModel.EmailAddress),
					ExternalKey:       customTypes.StringPointerValue(testDeviceModel.ExternalKey),
					ExternallyOwned:   types.BoolPointerValue(testDeviceModel.ExternallyOwned),
					Name:              customTypes.StringValue(testDeviceModel.Name),
					Owner:             types.StringValue(testDeviceModel.Owner),
					PhoneNumber:       types.StringValue(testDeviceModel.PhoneNumber),
					PIN:               customTypes.StringValue(testDeviceModel.PIN),
					PriorityThreshold: customTypes.StringValue(testDeviceModel.PriorityThreshold),
					Status:            types.StringValue(testDeviceModel.Status),
					TestStatus:        customTypes.StringValue(testDeviceModel.TestStatus),
					Timeframes: types.SetValueMust(utils.TimeframeObjectType,
						[]attr.Value{
							types.ObjectValueMust(utils.TimeframeObjectType.AttrTypes,
								map[string]attr.Value{
									"name":                customTypes.StringPointerValue(testDeviceModel.Timeframes[0].Name),
									"start_time":          types.StringPointerValue(testDeviceModel.Timeframes[0].StartTime),
									"duration_in_minutes": types.Int32PointerValue(testDeviceModel.Timeframes[0].DurationInMinutes),
									"days": types.SetValueMust(customTypes.CustomStringType{},
										[]attr.Value{
											customTypes.StringPointerValue(testDeviceModel.Timeframes[0].Days[0]),
											customTypes.StringPointerValue(testDeviceModel.Timeframes[0].Days[1]),
										},
									),
									"exclude_holidays": types.BoolPointerValue(testDeviceModel.Timeframes[0].ExcludeHolidays),
								},
							),
						},
					),
					TwoWayDevice: types.BoolPointerValue(testDeviceModel.TwoWayDevice),
				},
			},
			expected: testDeviceModel,
		},
	}
	for _, thisTest := range tests {
		actual := thisTest.args.in.DeviceParams(thisTest.args.diags)
		assert.Equal(t, thisTest.expected, actual)
	}
}
