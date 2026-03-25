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
		ExternalKey:       utils.RandStringPointer(10),
		ExternallyOwned:   utils.RandBoolPointer(),
		UseDefaultDevices: utils.RandBoolPointer(),
		Criteria: &xmatters.SearchCriteria{
			Operand: utils.RandOperandPointer(),
			Criterion: []*xmatters.SearchCriterion{
				{
					CriterionType: utils.RandStringPointer(10),
					Field:         utils.RandStringPointer(10),
					Operand:       utils.RandOperandPointer(),
					Value:         utils.RandStringPointer(15),
				},
			},
		},
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
				Criteria: types.ObjectNull(
					utils.GroupCriteriaObjectType.AttrTypes,
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
				ID:                types.StringPointerValue(testGroup.ID),
				Name:              types.StringPointerValue(testGroup.TargetName),
				Description:       types.StringPointerValue(testGroup.Description),
				Status:            types.StringPointerValue(testGroup.Status),
				ExternalKey:       types.StringPointerValue(testGroup.ExternalKey),
				ExternallyOwned:   types.BoolPointerValue(testGroup.ExternallyOwned),
				AllowDuplicates:   types.BoolPointerValue(testGroup.AllowDuplicates),
				UseDefaultDevices: types.BoolPointerValue(testGroup.UseDefaultDevices),
				Site:              utils.FlattenReferenceID(testGroup.Site),
				ObservedByAll:     types.BoolPointerValue(testGroup.ObservedByAll),
				GroupType:         types.StringPointerValue(testGroup.GroupType),
				Observers: types.SetValueMust(
					customTypes.CustomStringType{},
					[]attr.Value{
						customTypes.StringPointerValue(testGroup.Observers[0].Name),
					},
				),
				Supervisors: types.SetValueMust(
					types.StringType,
					[]attr.Value{
						types.StringPointerValue(testGroup.Supervisors[0].ID),
					},
				),
				Criteria: types.ObjectValueMust(
					utils.GroupCriteriaObjectType.AttrTypes,
					map[string]attr.Value{
						"operand": types.StringPointerValue(testGroup.Criteria.Operand),
						"criterion": types.SetValueMust(
							utils.GroupCriterionObjectType,
							[]attr.Value{
								types.ObjectValueMust(
									utils.GroupCriterionObjectType.AttrTypes,
									map[string]attr.Value{
										"criterion_type": types.StringPointerValue(testGroup.Criteria.Criterion[0].CriterionType),
										"field":          types.StringPointerValue(testGroup.Criteria.Criterion[0].Field),
										"operand":        customTypes.StringPointerValue(testGroup.Criteria.Criterion[0].Operand),
										"value":          customTypes.StringPointerValue(testGroup.Criteria.Criterion[0].Value),
									},
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
