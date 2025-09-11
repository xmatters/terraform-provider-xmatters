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
				TargetName:      customTypes.StringPointerValue(testGroup.TargetName),
				Status:          types.StringPointerValue(testGroup.Status),
				Description:     customTypes.StringPointerValue(testGroup.Description),
				GroupType:       types.StringPointerValue(testGroup.GroupType),
				AllowDuplicates: types.BoolPointerValue(testGroup.AllowDuplicates),
				Site:            types.StringPointerValue(testGroup.Site.ID),
				ObservedByAll:   types.BoolPointerValue(testGroup.ObservedByAll),
				Observers: types.SetValueMust(
					customTypes.CustomStringType{},
					[]attr.Value{
						customTypes.StringValue(*testGroup.Observers[0].Name),
					},
				),
				UseDefaultDevices: types.BoolPointerValue(testGroup.UseDefaultDevices),
				Supervisors: types.SetValueMust(
					types.StringType,
					[]attr.Value{
						types.StringValue(*testGroup.Supervisors[0].ID),
					},
				),
				ExternalKey:     customTypes.StringPointerValue(testGroup.ExternalKey),
				ExternallyOwned: types.BoolPointerValue(testGroup.ExternallyOwned),
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
					TargetName:      customTypes.StringValue(testGroupParams.TargetName),
					AllowDuplicates: types.BoolPointerValue(testGroupParams.AllowDuplicates),
					Description:     customTypes.StringValue(testGroupParams.Description),
					ExternalKey:     customTypes.StringValue(testGroupParams.ExternalKey),
					ExternallyOwned: types.BoolPointerValue(testGroupParams.ExternallyOwned),
					GroupType:       types.StringValue(testGroupParams.GroupType),
					ObservedByAll:   types.BoolPointerValue(testGroupParams.ObservedByAll),
					Observers: types.SetValueMust(
						types.StringType,
						[]attr.Value{
							types.StringValue(*testGroupParams.Observers[0].Name),
						},
					),
					Site:              types.StringValue(testGroupParams.Site),
					Status:            types.StringValue(testGroupParams.Status),
					UseDefaultDevices: types.BoolPointerValue(testGroupParams.UseDefaultDevices),
					Supervisors: types.SetValueMust(
						types.StringType,
						[]attr.Value{
							types.StringValue(*testGroupParams.Supervisors[0].ID),
						},
					),
				},
			},
			expected: testGroupParams,
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			actual := thisTest.args.in.GroupParams(thisTest.args.diags)
			assert.Equal(t, thisTest.expected, actual)
		})
	}
}
