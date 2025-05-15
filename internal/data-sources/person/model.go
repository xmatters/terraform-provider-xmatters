package person

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PersonModel describes the resource data model.
type PersonModel struct {
	PersonID        types.String `tfsdk:"person_id"`
	ID              types.String `tfsdk:"id"`
	TargetName      types.String `tfsdk:"target_name"`
	FirstName       types.String `tfsdk:"first_name"`
	LastName        types.String `tfsdk:"last_name"`
	Roles           types.Set    `tfsdk:"roles"`
	Status          types.String `tfsdk:"status"`
	WebLogin        types.String `tfsdk:"web_login"`
	Site            types.String `tfsdk:"site"`
	Timezone        types.String `tfsdk:"timezone"`
	Language        types.String `tfsdk:"language"`
	Supervisors     types.Set    `tfsdk:"supervisors"`
	PhoneLogin      types.String `tfsdk:"phone_login"`
	LicenseType     types.String `tfsdk:"license_type"`
	ExternalKey     types.String `tfsdk:"external_key"`
	ExternallyOwned types.Bool   `tfsdk:"externally_owned"`
	LastLogin       types.String `tfsdk:"last_login"`
}
