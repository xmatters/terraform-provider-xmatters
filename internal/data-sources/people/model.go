package people

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

// PeopleModel represents an xMatters People object in the Provider.
type PeopleModel struct {
	Search  *PeopleSearchModel  `tfsdk:"search" tf:"optional"`
	Filters *PeopleFiltersModel `tfsdk:"filters" tf:"optional"`
	Options *PeopleOptionsModel `tfsdk:"options" tf:"optional"`
	People  types.List          `tfsdk:"people" tf:"computed"`
}

// PeoplesSearchModel contains the search fields for the Provider's People data source.
type PeopleSearchModel struct {
	Terms   types.String `tfsdk:"terms" tf:"optional"`
	Operand types.String `tfsdk:"operand" tf:"optional"`
	Fields  types.List   `tfsdk:"fields" tf:"optional"`
}

// PeopleFiltersModel contains the filter fields for the Provider's People data source.
type PeopleFiltersModel struct {
	CreatedFrom        types.String `tfsdk:"created_from" tf:"optional"`
	CreatedTo          types.String `tfsdk:"created_to" tf:"optional"`
	CreatedBefore      types.String `tfsdk:"created_before" tf:"optional"`
	CreatedAfter       types.String `tfsdk:"created_after" tf:"optional"`
	DevicesExists      types.Bool   `tfsdk:"devices_exists" tf:"optional"`
	DevicesEmailExists types.Bool   `tfsdk:"devices_email_exists" tf:"optional"`
	DevicesFailsafe    types.Bool   `tfsdk:"devices_failsafe_exists" tf:"optional"`
	DevicesMobile      types.Bool   `tfsdk:"devices_mobile_exists" tf:"optional"`
	DevicesSMS         types.Bool   `tfsdk:"devices_sms_exists" tf:"optional"`
	DevicesVoice       types.Bool   `tfsdk:"devices_voice_exists" tf:"optional"`
	DevicesStatus      types.String `tfsdk:"devices_status" tf:"optional"`
	DevicesTestStatus  types.String `tfsdk:"devices_test_status" tf:"optional"`
	EmailAddress       types.String `tfsdk:"email_address" tf:"optional"`
	FirstName          types.String `tfsdk:"first_name" tf:"optional"`
	Groups             types.List   `tfsdk:"groups" tf:"optional"`
	GroupsExists       types.Bool   `tfsdk:"groups_exists" tf:"optional"`
	LastName           types.String `tfsdk:"last_name" tf:"optional"`
	LicenseType        types.String `tfsdk:"license_type" tf:"optional"`
	PhoneNumber        types.String `tfsdk:"phone_number" tf:"optional"`
	Roles              types.List   `tfsdk:"roles" tf:"optional"`
	Site               types.String `tfsdk:"site" tf:"optional"`
	Status             types.String `tfsdk:"status" tf:"optional"`
	Supervisors        types.List   `tfsdk:"supervisors" tf:"optional"`
	SupervisorsExists  types.Bool   `tfsdk:"supervisors_exists" tf:"optional"`
	TargetName         types.String `tfsdk:"target_name" tf:"optional"`
	WebLogin           types.String `tfsdk:"web_login" tf:"optional"`
}

// PeopleOptionsModel contains the options fields for the Provider's People data source.
type PeopleOptionsModel struct {
	SortBy    types.String `tfsdk:"sort_by" tf:"optional"`
	SortOrder types.String `tfsdk:"sort_order" tf:"optional"`
}

// APIParams returns the xmatters.GetPeopleParams object based on the PeoplesModel instance.
func (in PeopleModel) APIParams(diags *diag.Diagnostics) xmatters.GetPeopleParams {
	peopleParams := xmatters.GetPeopleParams{
		Embed: "roles,supervisors",
	}
	// Check for user provided Search fields
	if in.Search != nil {
		peopleParams.Terms = in.Search.Terms.ValueString()
		peopleParams.Operand = in.Search.Operand.ValueString()
		peopleParams.Fields = utils.ExpandStringList(diags, in.Search.Fields)
	}
	// Check for user provided Filter fields
	if in.Filters != nil {
		peopleParams.CreatedFrom = in.Filters.CreatedFrom.ValueString()
		peopleParams.CreatedTo = in.Filters.CreatedTo.ValueString()
		peopleParams.CreatedBefore = in.Filters.CreatedBefore.ValueString()
		peopleParams.CreatedAfter = in.Filters.CreatedAfter.ValueString()
		peopleParams.DevicesExists = in.Filters.DevicesExists.ValueBoolPointer()
		peopleParams.DevicesEmailExists = in.Filters.DevicesEmailExists.ValueBoolPointer()
		peopleParams.DevicesFailsafe = in.Filters.DevicesFailsafe.ValueBoolPointer()
		peopleParams.DevicesMobile = in.Filters.DevicesMobile.ValueBoolPointer()
		peopleParams.DevicesSMS = in.Filters.DevicesSMS.ValueBoolPointer()
		peopleParams.DevicesVoice = in.Filters.DevicesVoice.ValueBoolPointer()
		peopleParams.DevicesStatus = in.Filters.DevicesStatus.ValueString()
		peopleParams.DevicesTestStatus = in.Filters.DevicesTestStatus.ValueString()
		peopleParams.EmailAddress = in.Filters.EmailAddress.ValueString()
		peopleParams.FirstName = in.Filters.FirstName.ValueString()
		peopleParams.Groups = utils.ExpandEncodedStringList(diags, in.Filters.Groups) // Due to the way the API is designed, we need to URL-encode the group names prior to sending the request
		peopleParams.GroupsExists = in.Filters.GroupsExists.ValueBoolPointer()
		peopleParams.LastName = in.Filters.LastName.ValueString()
		peopleParams.LicenseType = in.Filters.LicenseType.ValueString()
		peopleParams.PhoneNumber = in.Filters.PhoneNumber.ValueString()
		peopleParams.Roles = utils.ExpandStringList(diags, in.Filters.Roles)
		peopleParams.Site = in.Filters.Site.ValueString()
		peopleParams.Status = in.Filters.Status.ValueString()
		peopleParams.Supervisors = utils.ExpandStringList(diags, in.Filters.Supervisors)
		peopleParams.SupervisorsExists = in.Filters.SupervisorsExists.ValueBoolPointer()
		peopleParams.TargetName = in.Filters.TargetName.ValueString()
		peopleParams.WebLogin = in.Filters.WebLogin.ValueString()
	}
	// Check for user provided Options fields
	if in.Options != nil {
		peopleParams.SortBy = in.Options.SortBy.ValueString()
		peopleParams.SortOrder = in.Options.SortOrder.ValueString()
	}
	return peopleParams
}
