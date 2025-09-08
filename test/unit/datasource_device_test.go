package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/device"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

func TestDeviceToModel(t *testing.T) {
	testDevice := xmatters.Device{
		ID:              utils.RandUUIDPointer(),
		TargetName:      utils.RandStringPointer(10),
		Country:         utils.RandStringPointer(3),
		DefaultDevice:   utils.RandBoolPointer(),
		Delay:           utils.RandInt32Pointer(),
		DeviceType:      utils.RandStringPointer(10),
		EmailAddress:    utils.RandStringPointer(15),
		ExternalKey:     utils.RandStringPointer(10),
		ExternallyOwned: utils.RandBoolPointer(),
		Name:            utils.RandStringPointer(10),
		Owner: &xmatters.PersonReference{
			ID: utils.RandUUIDPointer(),
		},
		PhoneNumber:       utils.RandStringPointer(12),
		PIN:               utils.RandStringPointer(6),
		PriorityThreshold: utils.RandStringPointer(5),
		Sequence:          utils.RandInt32Pointer(),
		Status:            utils.RandStringPointer(10),
		TestStatus:        utils.RandStringPointer(10),
		Timeframes: []*xmatters.DeviceTimeframe{
			{
				Name:              utils.RandStringPointer(10),
				StartTime:         utils.RandStringPointer(5),
				DurationInMinutes: utils.RandInt32Pointer(),
				Days:              utils.RandStringPointerList(3, 10),
				ExcludeHolidays:   utils.RandBoolPointer(),
			},
			{
				Name:              utils.RandStringPointer(10),
				StartTime:         utils.RandStringPointer(5),
				DurationInMinutes: utils.RandInt32Pointer(),
				Days:              utils.RandStringPointerList(2, 10),
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
			name: "full device",
			args: args{
				diags: &diag.Diagnostics{},
				in:    testDevice,
			},
			expected: device.DeviceModel{
				ID:                types.StringPointerValue(testDevice.ID),
				TargetName:        types.StringPointerValue(testDevice.TargetName),
				Country:           types.StringPointerValue(testDevice.Country),
				DefaultDevice:     types.BoolPointerValue(testDevice.DefaultDevice),
				Delay:             types.Int32PointerValue(testDevice.Delay),
				DeviceType:        types.StringPointerValue(testDevice.DeviceType),
				EmailAddress:      types.StringPointerValue(testDevice.EmailAddress),
				ExternalKey:       types.StringPointerValue(testDevice.ExternalKey),
				ExternallyOwned:   types.BoolPointerValue(testDevice.ExternallyOwned),
				Name:              types.StringPointerValue(testDevice.Name),
				Owner:             types.StringPointerValue(testDevice.Owner.ID),
				PhoneNumber:       types.StringPointerValue(testDevice.PhoneNumber),
				PIN:               types.StringPointerValue(testDevice.PIN),
				PriorityThreshold: types.StringPointerValue(testDevice.PriorityThreshold),
				Sequence:          types.Int32PointerValue(testDevice.Sequence),
				Status:            types.StringPointerValue(testDevice.Status),
				TestStatus:        types.StringPointerValue(testDevice.TestStatus),
				Timeframes: types.SetValueMust(
					utils.TimeframeObjectType,
					[]attr.Value{
						types.ObjectValueMust(
							utils.TimeframeObjectType.AttrTypes,
							map[string]attr.Value{
								"name":                customTypes.StringPointerValue(testDevice.Timeframes[0].Name),
								"start_time":          types.StringPointerValue(testDevice.Timeframes[0].StartTime),
								"duration_in_minutes": types.Int32PointerValue(testDevice.Timeframes[0].DurationInMinutes),
								"days": types.SetValueMust(
									customTypes.CustomStringType{},
									[]attr.Value{
										customTypes.StringPointerValue(testDevice.Timeframes[0].Days[0]),
										customTypes.StringPointerValue(testDevice.Timeframes[0].Days[1]),
										customTypes.StringPointerValue(testDevice.Timeframes[0].Days[2]),
									},
								),
								"exclude_holidays": types.BoolPointerValue(testDevice.Timeframes[0].ExcludeHolidays),
							},
						),
						types.ObjectValueMust(
							utils.TimeframeObjectType.AttrTypes,
							map[string]attr.Value{
								"name":                customTypes.StringPointerValue(testDevice.Timeframes[1].Name),
								"start_time":          types.StringPointerValue(testDevice.Timeframes[1].StartTime),
								"duration_in_minutes": types.Int32PointerValue(testDevice.Timeframes[1].DurationInMinutes),
								"days": types.SetValueMust(
									customTypes.CustomStringType{},
									[]attr.Value{
										customTypes.StringPointerValue(testDevice.Timeframes[1].Days[0]),
										customTypes.StringPointerValue(testDevice.Timeframes[1].Days[1]),
									},
								),
								"exclude_holidays": types.BoolPointerValue(testDevice.Timeframes[1].ExcludeHolidays),
							},
						),
					},
				),
				TwoWayDevice: types.BoolPointerValue(testDevice.TwoWayDevice),
			},
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := device.DeviceToModel(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.expected, got)
		})
	}
}
