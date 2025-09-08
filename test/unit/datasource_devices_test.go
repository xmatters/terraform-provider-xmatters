package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/devices"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

func TestDevicesAPIParams(t *testing.T) {
	testParams := xmatters.GetDevicesParams{
		DeviceStatus: utils.RandString(10),
		DeviceType:   utils.RandString(10),
		DeviceNames:  utils.RandString(15),
	}

	type args struct {
		diags *diag.Diagnostics
		in    devices.DevicesModel
	}
	tests := []struct {
		name   string
		args   args
		expect xmatters.GetDevicesParams
	}{
		{
			name: "Empty model",
			args: args{
				diags: &diag.Diagnostics{},
				in:    devices.DevicesModel{},
			},
			expect: xmatters.GetDevicesParams{
				Embed: "timeframes",
			},
		},
		{
			name: "With device status filter",
			args: args{
				diags: &diag.Diagnostics{},
				in: devices.DevicesModel{
					Filters: &devices.DevicesFilterModel{
						DeviceStatus: types.StringValue(testParams.DeviceStatus),
					},
				},
			},
			expect: xmatters.GetDevicesParams{
				Embed:        "timeframes",
				DeviceStatus: testParams.DeviceStatus,
			},
		},
		{
			name: "With device type filter",
			args: args{
				diags: &diag.Diagnostics{},
				in: devices.DevicesModel{
					Filters: &devices.DevicesFilterModel{
						DeviceType: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.DeviceType),
							},
						),
					},
				},
			},
			expect: xmatters.GetDevicesParams{
				Embed:      "timeframes",
				DeviceType: testParams.DeviceType,
			},
		},
		{
			name: "With device names filter",
			args: args{
				diags: &diag.Diagnostics{},
				in: devices.DevicesModel{
					Filters: &devices.DevicesFilterModel{
						DeviceNames: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.DeviceNames),
							},
						),
					},
				},
			},
			expect: xmatters.GetDevicesParams{
				Embed:       "timeframes",
				DeviceNames: testParams.DeviceNames,
			},
		},
		{
			name: "With all filters",
			args: args{
				diags: &diag.Diagnostics{},
				in: devices.DevicesModel{
					Filters: &devices.DevicesFilterModel{
						DeviceStatus: types.StringValue(testParams.DeviceStatus),
						DeviceType: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.DeviceType),
							},
						),
						DeviceNames: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.DeviceNames),
							},
						),
					},
				},
			},
			expect: xmatters.GetDevicesParams{
				Embed:        "timeframes",
				DeviceStatus: testParams.DeviceStatus,
				DeviceType:   testParams.DeviceType,
				DeviceNames:  testParams.DeviceNames,
			},
		},
		{
			name: "With multiple device types",
			args: args{
				diags: &diag.Diagnostics{},
				in: devices.DevicesModel{
					Filters: &devices.DevicesFilterModel{
						DeviceType: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue("EMAIL"),
								types.StringValue("VOICE"),
								types.StringValue("TEXT_PHONE"),
							},
						),
					},
				},
			},
			expect: xmatters.GetDevicesParams{
				Embed:      "timeframes",
				DeviceType: "EMAIL,VOICE,TEXT_PHONE",
			},
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := thisTest.args.in.APIParams(thisTest.args.diags)
			assert.Equal(t, thisTest.expect, got)
		})
	}
}
