package site

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

// SiteModel represents an xMatters Site object in the Provider.
type SiteModel struct {
	ID          types.String                   `tfsdk:"id"`
	Address1    customTypes.CustomStringValue  `tfsdk:"address1"`
	Address2    customTypes.CustomStringValue  `tfsdk:"address2"`
	City        customTypes.CustomStringValue  `tfsdk:"city"`
	Country     customTypes.CustomCountryValue `tfsdk:"country"`
	Language    types.String                   `tfsdk:"language"`
	Latitude    types.Float64                  `tfsdk:"latitude"`
	Longitude   types.Float64                  `tfsdk:"longitude"`
	Name        customTypes.CustomStringValue  `tfsdk:"name"`
	PostalCode  customTypes.CustomStringValue  `tfsdk:"postal_code"`
	State       customTypes.CustomStringValue  `tfsdk:"state"`
	Status      types.String                   `tfsdk:"status"`
	Timezone    types.String                   `tfsdk:"timezone"`
	LastUpdated types.String                   `tfsdk:"last_updated"`
}

// SiteParams is a method that takes the proposed configuration changes `SiteModel` and builds the API representation in the form of `*xmatters.PushSiteParams`.
// The reverse of this method is `SiteToModel` which handles building a state representation using the API response.
func (in SiteModel) SiteParams() xmatters.PushSiteParams {
	return xmatters.PushSiteParams{
		ID:         in.ID.ValueString(),
		Address1:   in.Address1.ValueStringPointer(),
		Address2:   in.Address2.ValueStringPointer(),
		City:       in.City.ValueStringPointer(),
		Country:    in.Country.ValueString(),
		Language:   in.Language.ValueString(),
		Latitude:   in.Latitude.ValueFloat64Pointer(),
		Longitude:  in.Longitude.ValueFloat64Pointer(),
		Name:       in.Name.ValueString(),
		PostalCode: in.PostalCode.ValueStringPointer(),
		State:      in.State.ValueStringPointer(),
		Status:     in.Status.ValueString(),
		Timezone:   in.Timezone.ValueString(),
	}
}
