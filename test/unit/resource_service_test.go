package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/resources/service"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/helpers"
	"github.com/xmatters/xmatters-go"
)

// ServiceToModel
func TestResourceServiceToModel(t *testing.T) {
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
		diags   *diag.Diagnostics
		service xmatters.Service
	}
	tests := []struct {
		name     string
		args     args
		expected service.ServiceModel
	}{
		{
			name: "empty service",
			args: args{
				diags:   &diag.Diagnostics{},
				service: xmatters.Service{},
			},
			expected: service.ServiceModel{
				ServiceLinks: types.SetValueMust(utils.ServiceLinkObjectType,
					[]attr.Value{},
				),
			},
		},
		{
			name: "valid service",
			args: args{
				diags:   &diag.Diagnostics{},
				service: testService,
			},
			expected: service.ServiceModel{
				ID:          types.StringPointerValue(testService.ID),
				Name:        customTypes.StringPointerValue(testService.TargetName),
				Description: customTypes.StringPointerValue(testService.Description),
				Type:        customTypes.StringPointerValue(testService.ServiceType),
				Tier:        types.StringPointerValue(testService.ServiceTier),
				Owner:       types.StringPointerValue(testService.OwnedBy.ID),
				ServiceLinks: types.SetValueMust(utils.ServiceLinkObjectType,
					[]attr.Value{
						types.ObjectValueMust(utils.ServiceLinkObjectType.AttrTypes,
							map[string]attr.Value{
								"link_text": customTypes.StringPointerValue(testService.ServiceLinks[0].Label),
								"url":       types.StringPointerValue(testService.ServiceLinks[0].URL),
							},
						),
						types.ObjectValueMust(utils.ServiceLinkObjectType.AttrTypes,
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
		actual := service.ServiceToModel(thisTest.args.diags, thisTest.args.service)
		assert.Equal(t, thisTest.expected, actual)
	}
}

// ServiceParams
func TestServiceParams(t *testing.T) {
	testParams := xmatters.PushServiceParams{
		TargetName:  utils.RandString(5),
		Description: helpers.StringPointer(utils.RandString(5)),
		OwnedBy: &xmatters.GroupReference{
			ID: utils.RandUUIDPointer(),
		},
		ServiceLinks: []*xmatters.ServiceLink{
			{
				Label: utils.RandStringPointer(5),
				URL:   utils.RandStringPointer(5),
			},
		},
	}
	type args struct {
		diags *diag.Diagnostics
	}
	tests := []struct {
		name     string
		args     args
		model    service.ServiceModel
		expected *xmatters.PushServiceParams
	}{
		{
			name: "empty model",
			args: args{
				diags: &diag.Diagnostics{},
			},
			model: service.ServiceModel{},
			expected: &xmatters.PushServiceParams{
				ServiceLinks: []*xmatters.ServiceLink{},
			},
		},
		{
			name: "valid model",
			args: args{
				diags: &diag.Diagnostics{},
			},
			model: service.ServiceModel{
				Name:        customTypes.StringValue(testParams.TargetName),
				Description: customTypes.StringPointerValue(testParams.Description),
				Owner:       types.StringPointerValue(testParams.OwnedBy.ID),
				ServiceLinks: types.SetValueMust(utils.ServiceLinkObjectType,
					[]attr.Value{
						types.ObjectValueMust(utils.ServiceLinkObjectType.AttrTypes,
							map[string]attr.Value{
								"link_text": customTypes.StringPointerValue(testParams.ServiceLinks[0].Label),
								"url":       types.StringPointerValue(testParams.ServiceLinks[0].URL),
							},
						),
					},
				),
			},
			expected: &testParams,
		},
	}
	for _, thisTest := range tests {
		actual := thisTest.model.ServiceParams(thisTest.args.diags)
		assert.Equal(t, *thisTest.expected, actual)
	}
}
