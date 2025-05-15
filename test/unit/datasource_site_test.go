package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/site"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

// SiteToModel
func TestSiteToModel(t *testing.T) {
	testSite := xmatters.Site{
		ID:         utils.RandUUIDPointer(),
		Address1:   utils.RandStringPointer(5),
		Address2:   utils.RandStringPointer(5),
		City:       utils.RandStringPointer(5),
		Country:    utils.RandStringPointer(5),
		Language:   utils.RandStringPointer(5),
		Latitude:   utils.RandomLatitudePointer(),
		Longitude:  utils.RandomLongitudePointer(),
		Name:       utils.RandStringPointer(5),
		PostalCode: utils.RandStringPointer(5),
		State:      utils.RandStringPointer(5),
		Status:     utils.RandStringPointer(5),
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
				ID:         types.StringPointerValue(testSite.ID),
				Address1:   types.StringPointerValue(testSite.Address1),
				Address2:   types.StringPointerValue(testSite.Address2),
				City:       types.StringPointerValue(testSite.City),
				Country:    types.StringPointerValue(testSite.Country),
				Language:   types.StringPointerValue(testSite.Language),
				Latitude:   types.Float64PointerValue(testSite.Latitude),
				Longitude:  types.Float64PointerValue(testSite.Longitude),
				Name:       types.StringPointerValue(testSite.Name),
				PostalCode: types.StringPointerValue(testSite.PostalCode),
				State:      types.StringPointerValue(testSite.State),
				Status:     types.StringPointerValue(testSite.Status),
				Timezone:   types.StringPointerValue(testSite.Timezone),
			},
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			actual := site.SiteToModel(thisTest.args)
			assert.Equal(t, thisTest.expected, actual)
		})
	}
}
