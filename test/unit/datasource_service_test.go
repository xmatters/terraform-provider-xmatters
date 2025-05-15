package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/service"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

// ServiceToModel
func TestServiceToModel(t *testing.T) {
	testService := xmatters.Service{
		ID:          utils.RandUUIDPointer(),
		TargetName:  utils.RandStringPointer(5),
		Description: utils.RandStringPointer(5),
		ServiceType: utils.RandStringPointer(5),
		ServiceTier: utils.RandStringPointer(5),
		OwnedBy: &xmatters.GroupReference{
			ID: utils.RandUUIDPointer(),
		},
		ServiceLinks: []*xmatters.ServiceLink{
			{
				Label: utils.RandStringPointer(5),
				URL:   utils.RandStringPointer(5),
			},
			{
				Label: utils.RandStringPointer(5),
				URL:   utils.RandStringPointer(5),
			},
		},
	}
	type args struct {
		diags *diag.Diagnostics
		in    xmatters.Service
	}
	tests := []struct {
		name     string
		args     args
		expected service.ServiceModel
	}{
		{
			name: "empty service",
			args: args{
				diags: &diag.Diagnostics{},
				in:    xmatters.Service{},
			},
			expected: service.ServiceModel{
				ServiceLinks: types.SetValueMust(
					utils.ServiceLinkObjectType,
					[]attr.Value{},
				),
			},
		},
		{
			name: "valid service",
			args: args{
				diags: &diag.Diagnostics{},
				in:    testService,
			},
			expected: service.ServiceModel{
				ID:          types.StringPointerValue(testService.ID),
				Name:        types.StringPointerValue(testService.TargetName),
				Description: types.StringPointerValue(testService.Description),
				Type:        types.StringPointerValue(testService.ServiceType),
				Tier:        types.StringPointerValue(testService.ServiceTier),
				Owner:       types.StringPointerValue(testService.OwnedBy.ID),
				ServiceLinks: types.SetValueMust(
					utils.ServiceLinkObjectType,
					[]attr.Value{
						types.ObjectValueMust(
							utils.ServiceLinkObjectType.AttrTypes,
							map[string]attr.Value{
								"link_text": customTypes.StringPointerValue(testService.ServiceLinks[0].Label),
								"url":       types.StringPointerValue(testService.ServiceLinks[0].URL),
							},
						),
						types.ObjectValueMust(
							utils.ServiceLinkObjectType.AttrTypes,
							map[string]attr.Value{
								"link_text": customTypes.StringPointerValue(testService.ServiceLinks[1].Label),
								"url":       types.StringPointerValue(testService.ServiceLinks[1].URL),
							},
						),
					},
				),
			},
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := service.ServiceToModel(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.expected, got)
		})
	}
}
