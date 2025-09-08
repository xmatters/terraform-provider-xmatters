package device

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *DeviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Create or update a device in xMatters. Provide fields in the request body that are common to all devices as well as fields that are specific to the type of device you want to create, for example, include the 'emailAddress' field when creating email devices and include the 'phoneNumber' field for phone, text (SMS), or public address devices.\n\nMobile app devices such as iPhone, iPad, and Android cannot be created using Terraform. These devices are added to a user's profile in xMatters automatically after they install the mobile app on their device and use it to log on to xMatters for the first time. However, once a mobile app device has been added to a user's profile it may be modified through Terraform.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: describe.DeviceID,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"device_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: describe.DeviceDeviceType,
				Validators: []validator.String{
					stringvalidator.OneOf("EMAIL", "VOICE", "VOICE_IVR", "TEXT_PHONE", "TEXT_PAGER"),
				},
			},
			"name": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.DeviceResourceName,
			},
			"owner": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: describe.DeviceOwner,
				Validators: []validator.String{
					utils.UUIDValidator{},
				},
			},
			"target_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: describe.DeviceTargetName,
			},
			"default_device": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: describe.DeviceResourceDefaultDevice,
			},
			"delay": schema.Int32Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: describe.DeviceDelay,
				Default:             int32default.StaticInt32(0),
			},
			"email_address": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Optional:            true,
				MarkdownDescription: describe.DeviceResourceEmailAddress,
				Validators: []validator.String{
					utils.EmailValidator{},
				},
			},
			"external_key": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Optional:            true,
				MarkdownDescription: describe.DeviceExternalKey,
			},
			"externally_owned": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: describe.DeviceResourceExternallyOwned,
				Default:             booldefault.StaticBool(true),
			},
			"phone_number": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: describe.DeviceResourcePhoneNumber,
			},
			"pin": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Optional:            true,
				MarkdownDescription: describe.DevicePIN,
			},
			"priority_threshold": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.DevicePriorityThreshold,
			},
			"status": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: describe.DeviceStatus,
				Validators: []validator.String{
					utils.StatusValidator{},
				},
			},
			"test_status": schema.StringAttribute{
				CustomType:          customTypes.CustomStringType{},
				Required:            true,
				MarkdownDescription: describe.DeviceTestStatus,
			},
			"timeframes": schema.SetNestedAttribute{
				Required:            true,
				MarkdownDescription: describe.DeviceResourceTimeframes,
				NestedObject: schema.NestedAttributeObject{
					Attributes: DeviceTimeframeSchema(),
				},
			},
			"two_way_device": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: describe.DeviceTwoWayDevice,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// DeviceTimeframeSchema returns the schema for a DeviceTimeframe.
func DeviceTimeframeSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"days": schema.SetAttribute{
			Required:            true,
			MarkdownDescription: describe.DeviceResourceTimeframeDays,
			ElementType:         customTypes.CustomStringType{},
		},
		"duration_in_minutes": schema.Int32Attribute{
			Required:            true,
			MarkdownDescription: describe.DeviceTimeframeDuration,
		},
		"exclude_holidays": schema.BoolAttribute{
			Required:            true,
			MarkdownDescription: describe.DeviceTimeframeHolidays,
		},
		"name": schema.StringAttribute{
			CustomType:          customTypes.CustomStringType{},
			Required:            true,
			MarkdownDescription: describe.DeviceTimeframeName,
		},
		"start_time": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: describe.DeviceResourceTimeframeStartTime,
			Validators: []validator.String{
				stringvalidator.RegexMatches(
					regexp.MustCompile(`^(?:[01]\d|2[0-3]):[0-5]\d$`),
					"Time must be in the format HH:mm",
				),
			},
		},
	}
}
