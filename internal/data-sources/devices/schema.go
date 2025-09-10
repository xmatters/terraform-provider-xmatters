package devices

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *DevicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Returns all devices in your xMatters instance that match the provided criteria.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: describe.DevicesFilters,
				Attributes:          DevicesDataSourceFiltersSchema(),
			},
			"devices": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: describe.DevicesList,
				NestedObject: schema.NestedAttributeObject{
					Attributes: DeviceDataSourceSchema(),
				},
			},
		},
	}
}

// DeviceDataSourceSchema is a helper function to simplify the DevicesDataSource Schema implementation.
// It represents the Provider's implmentation of the xMatters Device object.
func DeviceDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceID,
		},
		"target_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceTargetName,
		},
		"country": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceCountry,
		},
		"default_device": schema.BoolAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceDefaultDevice,
		},
		"delay": schema.Int32Attribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceDelay,
		},
		"device_type": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceDeviceType,
		},
		"email_address": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceEmailAddress,
		},
		"external_key": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceExternalKey,
		},
		"externally_owned": schema.BoolAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceExternallyOwned,
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceName,
		},
		"owner": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceOwner,
		},
		"phone_number": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DevicePhoneNumber,
		},
		"pin": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DevicePIN,
		},
		"priority_threshold": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DevicePriorityThreshold,
		},
		"sequence": schema.Int32Attribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceSequence,
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceStatus,
		},
		"test_status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceTestStatus,
		},
		"timeframes": schema.SetNestedAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceTimeframes,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"days": schema.SetAttribute{
						Computed:            true,
						MarkdownDescription: describe.DeviceTimeframeDays,
						ElementType:         types.StringType,
					},
					"duration_in_minutes": schema.Int32Attribute{
						Computed:            true,
						MarkdownDescription: describe.DeviceTimeframeDuration,
					},
					"exclude_holidays": schema.BoolAttribute{
						Computed:            true,
						MarkdownDescription: describe.DeviceTimeframeHolidays,
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: describe.DeviceTimeframeName,
					},
					"start_time": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: describe.DeviceTimeframeStartTime,
					},
				},
			},
		},
		"two_way_device": schema.BoolAttribute{
			Computed:            true,
			MarkdownDescription: describe.DeviceTwoWayDevice,
		},
	}

}

// DevicesDataSourceFiltersSchema is a helper function to simplify the DevicesDataSource Schema implementation.
// It represents the Provider's implmentation of the xMatters Device filters object.
func DevicesDataSourceFiltersSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"device_status": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.DeviceFiltersStatus,
			Validators: []validator.String{
				utils.StatusValidator{},
			},
		},
		"device_type": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.DeviceFiltersType,
			ElementType:         types.StringType,
		},
		"device_names": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.DeviceFiltersNames,
			ElementType:         types.StringType,
		},
	}
}
