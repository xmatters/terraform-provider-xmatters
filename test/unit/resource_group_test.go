package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/resources/group"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

func TestResourceGroupToModel(t *testing.T) {
	testGroup := xmatters.Group{
		TargetName:      utils.RandStringPointer(10),
		Status:          utils.RandStringPointer(5),
		Description:     utils.RandStringPointer(20),
		GroupType:       utils.RandStringPointer(5),
		AllowDuplicates: utils.RandBoolPointer(),
		Site: &xmatters.ReferenceById{
			ID: utils.RandUUIDPointer(),
		},
		ObservedByAll: utils.RandBoolPointer(),
		Observers: []*xmatters.ReferenceByName{
			{
				Name: utils.RandStringPointer(10),
			},
		},
		UseDefaultDevices: utils.RandBoolPointer(),
		Supervisors: []*xmatters.ReferenceById{
			{
				ID: utils.RandUUIDPointer(),
			},
		},
		ExternalKey:     utils.RandStringPointer(10),
		ExternallyOwned: utils.RandBoolPointer(),
		Criteria: &xmatters.SearchCriteria{
			Operand: utils.RandStringPointer(3),
			Criterion: []*xmatters.SearchCriterion{
				{
					CriterionType: utils.RandStringPointer(10),
					Field:         utils.RandStringPointer(10),
					Operand:       utils.RandStringPointer(5),
					Value:         utils.RandStringPointer(10),
				},
			},
		},
	}

	type args struct {
		diags *diag.Diagnostics
		in    xmatters.Group
	}

	tests := []struct {
		name     string
		args     args
		expected group.GroupModel
	}{
		{
			name: "empty group",
			args: args{
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
			name: "valid group",
			args: args{
				diags: &diag.Diagnostics{},
				in:    testGroup,
			},
			expected: group.GroupModel{
				ID:              types.StringPointerValue(testGroup.ID),
				Name:            customTypes.StringPointerValue(testGroup.TargetName),
				Status:          types.StringPointerValue(testGroup.Status),
				Description:     customTypes.StringPointerValue(testGroup.Description),
				GroupType:       types.StringPointerValue(testGroup.GroupType),
				AllowDuplicates: types.BoolPointerValue(testGroup.AllowDuplicates),
				Site:            types.StringPointerValue(testGroup.Site.ID),
				ObservedByAll:   types.BoolPointerValue(testGroup.ObservedByAll),
				Observers: types.SetValueMust(
					customTypes.CustomStringType{},
					[]attr.Value{
						customTypes.StringPointerValue(testGroup.Observers[0].Name),
					},
				),
				UseDefaultDevices: types.BoolPointerValue(testGroup.UseDefaultDevices),
				Supervisors: types.SetValueMust(
					types.StringType,
					[]attr.Value{
						types.StringPointerValue(testGroup.Supervisors[0].ID),
					},
				),
				ExternalKey:     customTypes.StringPointerValue(testGroup.ExternalKey),
				ExternallyOwned: types.BoolPointerValue(testGroup.ExternallyOwned),
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

func TestGroupParams(t *testing.T) {
	testGroupParams := xmatters.PushGroupParams{
		TargetName:      utils.RandString(10),
		AllowDuplicates: utils.RandBoolPointer(),
		Description:     utils.RandString(10),
		ExternalKey:     utils.RandString(10),
		ExternallyOwned: utils.RandBoolPointer(),
		GroupType:       utils.RandString(5),
		ObservedByAll:   utils.RandBoolPointer(),
		Observers: []*xmatters.ReferenceByName{
			{
				Name: utils.RandStringPointer(10),
			},
		},
		Site:              utils.RandString(5),
		Status:            utils.RandString(5),
		UseDefaultDevices: utils.RandBoolPointer(),
		Supervisors: []*xmatters.ReferenceById{
			{
				ID: utils.RandUUIDPointer(),
			},
		},
		Criteria: &xmatters.SearchCriteria{
			Operand: utils.RandStringPointer(3),
			Criterion: []*xmatters.SearchCriterion{
				{
					CriterionType: utils.RandStringPointer(10),
					Field:         utils.RandStringPointer(10),
					Operand:       utils.RandStringPointer(5),
					Value:         utils.RandStringPointer(10),
				},
			},
		},
	}
	type args struct {
		diags *diag.Diagnostics
		in    group.GroupModel
	}
	tests := []struct {
		name     string
		args     args
		expected xmatters.PushGroupParams
	}{
		{
			name: "empty group model",
			args: args{
				diags: &diag.Diagnostics{},
				in:    group.GroupModel{},
			},
			expected: xmatters.PushGroupParams{},
		},
		{
			name: "valid model",
			args: args{
				diags: &diag.Diagnostics{},
				in: group.GroupModel{
					Name:            customTypes.StringValue(testGroupParams.TargetName),
					AllowDuplicates: types.BoolPointerValue(testGroupParams.AllowDuplicates),
					Description:     customTypes.StringValue(testGroupParams.Description),
					ExternalKey:     customTypes.StringValue(testGroupParams.ExternalKey),
					ExternallyOwned: types.BoolPointerValue(testGroupParams.ExternallyOwned),
					GroupType:       types.StringValue(testGroupParams.GroupType),
					ObservedByAll:   types.BoolPointerValue(testGroupParams.ObservedByAll),
					Observers: types.SetValueMust(
						types.StringType,
						[]attr.Value{
							types.StringPointerValue(testGroupParams.Observers[0].Name),
						},
					),
					Site:              types.StringValue(testGroupParams.Site),
					Status:            types.StringValue(testGroupParams.Status),
					UseDefaultDevices: types.BoolPointerValue(testGroupParams.UseDefaultDevices),
					Supervisors: types.SetValueMust(
						types.StringType,
						[]attr.Value{
							types.StringPointerValue(testGroupParams.Supervisors[0].ID),
						},
					),
					Criteria: types.ObjectValueMust(
						utils.GroupCriteriaObjectType.AttrTypes,
						map[string]attr.Value{
							"operand": types.StringPointerValue(testGroupParams.Criteria.Operand),
							"criterion": types.SetValueMust(
								utils.GroupCriterionObjectType,
								[]attr.Value{
									types.ObjectValueMust(
										utils.GroupCriterionObjectType.AttrTypes,
										map[string]attr.Value{
											"criterion_type": types.StringPointerValue(testGroupParams.Criteria.Criterion[0].CriterionType),
											"field":          types.StringPointerValue(testGroupParams.Criteria.Criterion[0].Field),
											"operand":        customTypes.StringPointerValue(testGroupParams.Criteria.Criterion[0].Operand),
											"value":          customTypes.StringPointerValue(testGroupParams.Criteria.Criterion[0].Value),
										},
									),
								},
							),
						},
					),
				},
			},
			expected: testGroupParams,
		},
		{
			name: "explicit empty supervisors",
			args: args{
				diags: &diag.Diagnostics{},
				in: group.GroupModel{
					Name: customTypes.StringValue(""),
					Supervisors: types.SetValueMust(
						types.StringType,
						[]attr.Value{},
					),
				},
			},
			expected: xmatters.PushGroupParams{
				TargetName:  "",
				Description: "",
				Supervisors: []*xmatters.ReferenceById{},
			},
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			actual := thisTest.args.in.GroupParams(thisTest.args.diags)
			assert.Equal(t, thisTest.expected, actual)
		})
	}
}
