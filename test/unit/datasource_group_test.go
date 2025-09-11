package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/group"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

func TestGroupToModel(t *testing.T) {
	testGroup := xmatters.Group{
		ID:              utils.RandUUIDPointer(),
		TargetName:      utils.RandStringPointer(10),
		Status:          utils.RandStringPointer(5),
		Description:     utils.RandStringPointer(20),
		GroupType:       utils.RandStringPointer(8),
		AllowDuplicates: utils.RandBoolPointer(),
		Timezone:        utils.RandStringPointer(10),
		Site:            &xmatters.ReferenceById{ID: utils.RandUUIDPointer()},
		ObservedByAll:   utils.RandBoolPointer(),
		Observers: []*xmatters.ReferenceByName{
			{
				Name: utils.RandStringPointer(10),
			},
		},
		Supervisors: []*xmatters.ReferenceById{
			{
				ID: utils.RandUUIDPointer(),
			},
		},
		Services: []*xmatters.Service{
			{
				ID:         utils.RandUUIDPointer(),
				TargetName: utils.RandStringPointer(10),
			},
		},
		ExternalKey:     utils.RandStringPointer(10),
		ExternallyOwned: utils.RandBoolPointer(),
	}

	tests := []struct {
		name string
		args struct {
			diags *diag.Diagnostics
			in    xmatters.Group
		}
		expected group.GroupModel
	}{
		{
			name: "empty group",
			args: struct {
				diags *diag.Diagnostics
				in    xmatters.Group
			}{
				diags: &diag.Diagnostics{},
				in:    xmatters.Group{},
			},
			expected: group.GroupModel{
				Observers: types.SetValueMust(
					customTypes.CustomStringType{},
					[]attr.Value{},
				),
				Supervisors: types.SetValueMust(
					types.StringType,
					[]attr.Value{},
				),
				Services: types.SetValueMust(
					utils.ServiceObjectType,
					[]attr.Value{},
				),
			},
		},
		{
			name: "full group",
			args: struct {
				diags *diag.Diagnostics
				in    xmatters.Group
			}{
				diags: &diag.Diagnostics{},
				in:    testGroup,
			},
			expected: group.GroupModel{
				ID:              types.StringPointerValue(testGroup.ID),
				TargetName:      types.StringPointerValue(testGroup.TargetName),
				Description:     types.StringPointerValue(testGroup.Description),
				Status:          types.StringPointerValue(testGroup.Status),
				ExternalKey:     types.StringPointerValue(testGroup.ExternalKey),
				ExternallyOwned: types.BoolPointerValue(testGroup.ExternallyOwned),
				AllowDuplicates: types.BoolPointerValue(testGroup.AllowDuplicates),
				Timezone:        types.StringPointerValue(testGroup.Timezone),
				Site:            utils.FlattenReferenceID(testGroup.Site),
				ObservedByAll:   types.BoolPointerValue(testGroup.ObservedByAll),
				GroupType:       types.StringPointerValue(testGroup.GroupType),
				Observers: types.SetValueMust(
					customTypes.CustomStringType{},
					[]attr.Value{
						customTypes.StringValue(*testGroup.Observers[0].Name),
					},
				),
				Supervisors: types.SetValueMust(
					types.StringType,
					[]attr.Value{
						types.StringValue(*testGroup.Supervisors[0].ID),
					},
				),
				Services: types.SetValueMust(
					utils.ServiceObjectType,
					[]attr.Value{
						types.ObjectValueMust(
							utils.ServiceObjectType.AttrTypes,
							map[string]attr.Value{
								"id":          types.StringPointerValue(testGroup.Services[0].ID),
								"name":        customTypes.StringPointerValue(testGroup.Services[0].TargetName),
								"description": customTypes.StringPointerValue(nil),
								"type":        customTypes.StringPointerValue(nil),
								"tier":        types.StringPointerValue(nil),
								"owner":       types.StringPointerValue(nil),
								"links": types.SetValueMust(
									utils.ServiceLinkObjectType,
									[]attr.Value{},
								),
							},
						),
					},
				),
			},
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			actual := group.GroupToModel(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.expected, actual)
		})
	}
}
