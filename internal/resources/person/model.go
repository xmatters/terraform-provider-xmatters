package person

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

// PersonModel describes the resource data model.
type PersonModel struct {
	ID              types.String                  `tfsdk:"id"`
	TargetName      customTypes.CustomStringValue `tfsdk:"target_name"`
	FirstName       customTypes.CustomStringValue `tfsdk:"first_name"`
	LastName        customTypes.CustomStringValue `tfsdk:"last_name"`
	Roles           types.Set                     `tfsdk:"roles"`
	Status          types.String                  `tfsdk:"status"`
	WebLogin        customTypes.CustomStringValue `tfsdk:"web_login"`
	Site            types.String                  `tfsdk:"site"`
	Timezone        types.String                  `tfsdk:"timezone"`
	Language        types.String                  `tfsdk:"language"`
	Supervisors     types.Set                     `tfsdk:"supervisors"`
	PhoneLogin      types.String                  `tfsdk:"phone_login"`
	PhonePin        types.String                  `tfsdk:"phone_pin"`
	LicenseType     customTypes.CustomStringValue `tfsdk:"license_type"`
	ExternalKey     customTypes.CustomStringValue `tfsdk:"external_key"`
	ExternallyOwned types.Bool                    `tfsdk:"externally_owned"`
	LastUpdated     types.String                  `tfsdk:"last_updated"`
}

// PersonParams is a method that takes the proposed configuration changes `PersonModel` and builds the API representation in the form of `*xmatters.PushPersonParams`.
// The reverse of this method is `PersonToModel` which handles building a state representation using the API response.
func (in PersonModel) PersonParams(diags *diag.Diagnostics) xmatters.PushPersonParams {
	return xmatters.PushPersonParams{
		ID:              in.ID.ValueString(),
		TargetName:      in.TargetName.ValueString(),
		FirstName:       in.FirstName.ValueString(),
		LastName:        in.LastName.ValueString(),
		Roles:           utils.ExpandStringSliceSet(diags, in.Roles),
		Status:          in.Status.ValueString(),
		WebLogin:        in.WebLogin.ValueString(),
		Site:            in.Site.ValueString(),
		Timezone:        in.Timezone.ValueString(),
		Language:        in.Language.ValueString(),
		Supervisors:     utils.ExpandStringSliceSet(diags, in.Supervisors),
		PhoneLogin:      in.PhoneLogin.ValueStringPointer(),
		PhonePin:        in.PhonePin.ValueString(),
		LicenseType:     in.LicenseType.ValueString(),
		ExternalKey:     in.ExternalKey.ValueStringPointer(),
		ExternallyOwned: in.ExternallyOwned.ValueBoolPointer(),
	}
}
