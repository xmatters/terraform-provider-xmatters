package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/services"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

// GetServicesParams
func TestGetServicesParams(t *testing.T) {
	testParams := xmatters.GetServicesParams{
		Search:  utils.RandString(5),
		Fields:  utils.RandString(5),
		Operand: utils.RandString(5),
		OwnedBy: utils.RandUUID(),
	}
	type args struct {
		diags *diag.Diagnostics
		model services.ServicesModel
	}
	tests := []struct {
		name     string
		args     args
		expected xmatters.GetServicesParams
	}{
		{
			name: "empty model",
			args: args{
				diags: &diag.Diagnostics{},
				model: services.ServicesModel{},
			},
			expected: xmatters.GetServicesParams{},
		},
		{
			name: "valid model",
			args: args{
				diags: &diag.Diagnostics{},
				model: services.ServicesModel{
					Search: &services.ServicesSearchModel{
						Terms: types.StringValue(testParams.Search),
						Fields: types.ListValueMust(types.StringType,
							[]attr.Value{
								types.StringValue(testParams.Fields),
							},
						),
						Operand: types.StringValue(testParams.Operand),
					},
					Owner: types.StringValue(testParams.OwnedBy),
				},
			},
			expected: testParams,
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			// Call the getServicesParams method
			params := thisTest.args.model.APIParams(thisTest.args.diags)
			// Assert the expected values
			assert.Equal(t, thisTest.expected, params)
		})
	}
}
