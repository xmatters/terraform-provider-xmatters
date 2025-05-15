package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/sites"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

// GetSitesParams
func TestGetSitesParams(t *testing.T) {
	testParams := xmatters.GetSitesParams{
		Search:   utils.RandString(5),
		Fields:   utils.RandString(5),
		Operand:  utils.RandString(5),
		Country:  utils.RandString(5),
		Geocoded: utils.RandBoolPointer(),
		Status:   utils.RandString(5),
	}
	type args struct {
		diags *diag.Diagnostics
		model sites.SitesModel
	}
	tests := []struct {
		name     string
		args     args
		expected xmatters.GetSitesParams
	}{
		{
			name: "empty model",
			args: args{
				diags: &diag.Diagnostics{},
				model: sites.SitesModel{},
			},
			expected: xmatters.GetSitesParams{},
		},
		{
			name: "valid model",
			args: args{
				diags: &diag.Diagnostics{},
				model: sites.SitesModel{
					Search: &sites.SitesSearchModel{
						Terms: types.StringValue(testParams.Search),
						Fields: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.Fields),
							},
						),
						Operand: types.StringValue(testParams.Operand),
					},
					Filters: &sites.SitesFilterModel{
						Country:  types.StringValue(testParams.Country),
						Geocoded: types.BoolPointerValue(testParams.Geocoded),
						Status:   types.StringValue(testParams.Status),
					},
				},
			},
			expected: testParams,
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			// Call the getSitesParams method
			params := thisTest.args.model.APIParams(thisTest.args.diags)
			// Assert the expected values
			assert.Equal(t, thisTest.expected, params)
		})
	}
}
