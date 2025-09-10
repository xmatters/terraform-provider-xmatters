package device

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/helpers"
	"github.com/xmatters/xmatters-go"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                   = &DeviceResource{}
	_ resource.ResourceWithConfigure      = &DeviceResource{}
	_ resource.ResourceWithImportState    = &DeviceResource{}
	_ resource.ResourceWithValidateConfig = &DeviceResource{}
)

// NewDeviceResource is a helper function to simplify the provider implementation.
func NewDeviceResource() resource.Resource {
	return &DeviceResource{}
}

// DeviceResource defines the resource implementation.
type DeviceResource struct {
	client *xmatters.XMattersAPI
}

// Metadata returns the resource type name.
func (r *DeviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

// Configure adds the provider configured client to the resource.
func (r *DeviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*xmatters.XMattersAPI)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Create creates the resource and sets the initial Terraform state.
func (r *DeviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	createParams := plan.DeviceParams(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new xMatters Device
	deviceReturn, err := r.client.PushDevice(createParams)
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("device", err))
		return
	}

	// Map response body to schema and populate Computed attribute values
	state := DeviceToModel(&resp.Diagnostics, deviceReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Created xMatters Device", map[string]any{"success": true})
}

// Read refreshes the Terraform state with the latest data.
func (r *DeviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceModel

	// Get current state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed Device from xMatters
	deviceReturn, err := r.client.GetDevice(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DatasourceErrorDiagnostic("device", err))
		return
	}

	// Capture state `lastUpdated` timestamp
	lastUpdated := state.LastUpdated
	// Overwrite Device with refreshed state
	state = DeviceToModel(&resp.Diagnostics, deviceReturn)
	if resp.Diagnostics.HasError() {
		return
	}

	// Ensure `last_updated` is preserved and set to the new state
	state.LastUpdated = lastUpdated

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Retrieved xMatters Device data", map[string]any{"success": true})
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *DeviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DeviceModel

	// Retrieve values from plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	updateParams := plan.DeviceParams(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing xMatters Device and fetch updated Device
	deviceReturn, err := r.client.PushDevice(updateParams)
	if err != nil {
		resp.Diagnostics.Append(helpers.ResourceErrorDiagnostic("device", err))
		return
	}

	// Updated resource state with updated data and timestamp
	state := DeviceToModel(&resp.Diagnostics, deviceReturn)
	if resp.Diagnostics.HasError() {
		return
	}
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Save state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Updated xMatters Device", map[string]any{"success": true})
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *DeviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceModel

	// Retrieve values from state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing Device
	err := r.client.DeleteDevice(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.Append(helpers.DeleteErrorDiagnostic("device", err))
		return
	}
}

// ImportState imports a resource from an existing xMatters resource.
func (r *DeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// ValidateConfig validates the configuration of the resource.
func (r *DeviceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data DeviceModel

	// Get configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate DeviceType requirements
	switch data.DeviceType.ValueString() {
	case "EMAIL":
		// Email Address is required
		if data.EmailAddress.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("email_address"),
				"Missing Attribute Configuration",
				"Expected 'email_address' to be configured when 'device_type' is \"EMAIL\".",
			)
		}
		// PhoneNumber, PIN, and TwoWayDevice are not allowed
		if !data.PhoneNumber.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("phone_number"),
				"Invalid Attribute Configuration",
				"Expected 'phone_number' to be empty when 'device_type' is \"EMAIL\".",
			)
		}
		if !data.PIN.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("pin"),
				"Invalid Attribute Configuration",
				"Expected 'pin' to be empty when 'device_type' is \"EMAIL\".",
			)
		}
		if !data.TwoWayDevice.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("two_way_device"),
				"Invalid Attribute Configuration",
				"Expected 'two_way_device' to be empty when 'device_type' is \"EMAIL\".",
			)
		}
	case "VOICE":
		// PhoneNumber is required
		if data.PhoneNumber.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("phone_number"),
				"Missing Attribute Configuration",
				"Expected 'phone_number' to be configured when 'device_type' is \"VOICE\".",
			)
		}
		// Email Address, PIN, and TwoWayDevice are not allowed
		if !data.EmailAddress.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("email_address"),
				"Invalid Attribute Configuration",
				"Expected 'email_address' to be empty when 'device_type' is \"VOICE\". ",
			)
		}
		if !data.PIN.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("pin"),
				"Invalid Attribute Configuration",
				"Expected 'pin' to be empty when 'device_type' is \"VOICE\". ",
			)
		}
		if !data.TwoWayDevice.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("two_way_device"),
				"Invalid Attribute Configuration",
				"Expected 'two_way_device' to be empty when 'device_type' is \"VOICE\". ",
			)
		}
	case "VOICE_IVR":
		// PhoneNumber is required
		if data.PhoneNumber.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("phone_number"),
				"Missing Attribute Configuration",
				"Expected 'phone_number' to be configured when 'device_type' is \"VOICE_IVR\".",
			)
		}
		// Email Address, PIN, and TwoWayDevice are not allowed
		if !data.EmailAddress.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("email_address"),
				"Invalid Attribute Configuration",
				"Expected 'email_address' to be empty when 'device_type' is \"VOICE_IVR\". ",
			)
		}
		if !data.PIN.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("pin"),
				"Invalid Attribute Configuration",
				"Expected 'pin' to be empty when 'device_type' is \"VOICE_IVR\". ",
			)
		}
		if !data.TwoWayDevice.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("two_way_device"),
				"Invalid Attribute Configuration",
				"Expected 'two_way_device' to be empty when 'device_type' is \"VOICE_IVR\". ",
			)
		}
	case "TEXT_PHONE":
		// PhoneNumber is required
		if data.PhoneNumber.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("phone_number"),
				"Missing Attribute Configuration",
				"Expected 'phone_number' to be configured when 'device_type' is \"TEXT_PHONE\".",
			)
		}
		// Email Address, Country, PIN, and TwoWayDevice are not allowed
		if !data.EmailAddress.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("email_address"),
				"Invalid Attribute Configuration",
				"Expected 'email_address' to be empty when 'device_type' is \"TEXT_PHONE\". ",
			)
		}
		if !data.PIN.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("pin"),
				"Invalid Attribute Configuration",
				"Expected 'pin' to be empty when 'device_type' is \"TEXT_PHONE\". ",
			)
		}
		if !data.TwoWayDevice.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("two_way_device"),
				"Invalid Attribute Configuration",
				"Expected 'two_way_device' to be empty when 'device_type' is \"TEXT_PHONE\". ",
			)
		}
	case "TEXT_PAGER":
		// PIN and TwoWayDevice are required
		if data.PIN.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("pin"),
				"Missing Attribute Configuration",
				"Expected 'pin' to be configured when 'device_type' is \"TEXT_PAGER\".",
			)
		}
		if data.TwoWayDevice.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("two_way_device"),
				"Missing Attribute Configuration",
				"Expected 'two_way_device' to be configured when 'device_type' is \"TEXT_PAGER\".",
			)
		}
		// Email Address and PhoneNumber are not allowed
		if !data.EmailAddress.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("email_address"),
				"Invalid Attribute Configuration",
				"Expected 'email_address' to be empty when 'device_type' is \"TEXT_PAGER\". ",
			)
		}
		if !data.PhoneNumber.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("phone_number"),
				"Invalid Attribute Configuration",
				"Expected 'phone_number' to be empty when 'device_type' is \"TEXT_PAGER\". ",
			)
		}
	}

	// Validate PhoneNumber format when configured
	if !data.PhoneNumber.IsNull() && !data.PhoneNumber.IsUnknown() {
		if !regexp.MustCompile(`^\+[1-9]\d{1,14}$`).MatchString(data.PhoneNumber.ValueString()) {
			resp.Diagnostics.AddAttributeError(
				path.Root("phone_number"),
				"Invalid Attribute Configuration",
				"Expected 'phone_number' to be in E.164 format: +[1-9]{1,14}.",
			)
		}
	}
}

// DeviceToModel is a function that takes the API payload `xmatters.Device` and builds the state representation in the form of `DeviceModel`.
// The reverse of this method is `DeviceParams` which handles building an API representation using the proposed config.
func DeviceToModel(diags *diag.Diagnostics, device xmatters.Device) DeviceModel {
	model := DeviceModel{
		ID:                types.StringPointerValue(device.ID),
		TargetName:        types.StringPointerValue(device.TargetName),
		DefaultDevice:     types.BoolPointerValue(device.DefaultDevice),
		Delay:             types.Int32PointerValue(device.Delay),
		DeviceType:        types.StringPointerValue(device.DeviceType),
		EmailAddress:      customTypes.StringPointerValue(device.EmailAddress),
		ExternalKey:       customTypes.StringPointerValue(device.ExternalKey),
		ExternallyOwned:   types.BoolPointerValue(device.ExternallyOwned),
		Name:              customTypes.StringPointerValue(device.Name),
		Owner:             utils.FlattenPersonReferenceID(device.Owner),
		PhoneNumber:       types.StringPointerValue(device.PhoneNumber),
		PIN:               customTypes.StringPointerValue(device.PIN),
		PriorityThreshold: customTypes.StringPointerValue(device.PriorityThreshold),
		Status:            types.StringPointerValue(device.Status),
		TestStatus:        customTypes.StringPointerValue(device.TestStatus),
		Timeframes:        utils.FlattenTimeframeSet(diags, device.Timeframes),
		TwoWayDevice:      types.BoolPointerValue(device.TwoWayDevice),
	}
	if diags.HasError() {
		return DeviceModel{}
	}
	return model
}
