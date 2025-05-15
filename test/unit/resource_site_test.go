package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/resources/site"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

// SiteToModel
func TestResourceSiteToModel(t *testing.T) {
	testSite := xmatters.Site{
		Address1:   utils.RandStringPointer(5),
		Address2:   utils.RandStringPointer(5),
		City:       utils.RandStringPointer(5),
		Country:    utils.RandStringPointer(5),
		ID:         utils.RandUUIDPointer(),
		Language:   utils.RandStringPointer(5),
		Latitude:   utils.RandomLatitudePointer(),
		Longitude:  utils.RandomLongitudePointer(),
		Name:       utils.RandStringPointer(5),
		PostalCode: utils.RandStringPointer(5),
		State:      utils.RandStringPointer(5),
		Status:     utils.RandStringPointer(5),
		Timezone:   utils.RandStringPointer(5),
	}
	tests := []struct {
		name     string
		args     xmatters.Site
		expected site.SiteModel
	}{
		{
			name:     "empty site",
			args:     xmatters.Site{},
			expected: site.SiteModel{},
		},
		{
			name: "valid site",
			args: testSite,
			expected: site.SiteModel{
				Address1:   customTypes.StringPointerValue(testSite.Address1),
				Address2:   customTypes.StringPointerValue(testSite.Address2),
				City:       customTypes.StringPointerValue(testSite.City),
				Country:    customTypes.CountryPointerValue(testSite.Country),
				ID:         types.StringPointerValue(testSite.ID),
				Language:   types.StringPointerValue(testSite.Language),
				Latitude:   types.Float64PointerValue(testSite.Latitude),
				Longitude:  types.Float64PointerValue(testSite.Longitude),
				Name:       customTypes.StringPointerValue(testSite.Name),
				PostalCode: customTypes.StringPointerValue(testSite.PostalCode),
				State:      customTypes.StringPointerValue(testSite.State),
				Status:     types.StringPointerValue(testSite.Status),
				Timezone:   types.StringPointerValue(testSite.Timezone),
			},
		},
	}
	for _, thisTest := range tests {
		actual := site.SiteToModel(thisTest.args)
		assert.Equal(t, thisTest.expected, actual)
	}

}

// SiteParams
func TestResourceSiteParams(t *testing.T) {
	testParams := xmatters.PushSiteParams{
		Address1:   utils.RandStringPointer(5),
		Address2:   utils.RandStringPointer(5),
		City:       utils.RandStringPointer(5),
		Country:    utils.RandString(5),
		ID:         utils.RandUUID(),
		Language:   utils.RandString(5),
		Latitude:   utils.RandomLatitudePointer(),
		Longitude:  utils.RandomLongitudePointer(),
		Name:       utils.RandString(5),
		PostalCode: utils.RandStringPointer(5),
		State:      utils.RandStringPointer(5),
		Status:     utils.RandString(5),
		Timezone:   utils.RandString(5),
	}
	tests := []struct {
		name     string
		model    site.SiteModel
		expected xmatters.PushSiteParams
	}{
		{
			name:     "empty model",
			model:    site.SiteModel{},
			expected: xmatters.PushSiteParams{},
		},
		{
			name: "valid model",
			model: site.SiteModel{
				Address1:   customTypes.StringPointerValue(testParams.Address1),
				Address2:   customTypes.StringPointerValue(testParams.Address2),
				City:       customTypes.StringPointerValue(testParams.City),
				Country:    customTypes.CountryValue(testParams.Country),
				ID:         types.StringValue(testParams.ID),
				Language:   types.StringValue(testParams.Language),
				Latitude:   types.Float64PointerValue(testParams.Latitude),
				Longitude:  types.Float64PointerValue(testParams.Longitude),
				Name:       customTypes.StringValue(testParams.Name),
				PostalCode: customTypes.StringPointerValue(testParams.PostalCode),
				State:      customTypes.StringPointerValue(testParams.State),
				Status:     types.StringValue(testParams.Status),
				Timezone:   types.StringValue(testParams.Timezone),
			},
			expected: testParams,
		},
	}
	for _, thisTest := range tests {
		actual := thisTest.model.SiteParams()
		assert.Equal(t, thisTest.expected, actual)
	}
}
