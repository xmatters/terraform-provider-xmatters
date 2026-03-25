package utils

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/helpers"
	"github.com/xmatters/xmatters-go"
)

// ------------------------------------------------------------
// List Expanders
// ------------------------------------------------------------

func TestExpandStringPointerSliceList(t *testing.T) {
	testObservers := []*xmatters.Role{
		{
			ID: RandUUIDPointer(),
		},
		{
			ID: RandUUIDPointer(),
		},
	}

	type args struct {
		diags diag.Diagnostics
		in    basetypes.ListValue
	}
	tests := []struct {
		name string
		args args
		want []*string
	}{
		{
			name: "empty",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.ListNull(types.StringType),
			},
			want: nil,
		},
		{
			name: "invalid item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.Int32Type, []attr.Value{
					types.Int32Value(5),
				}),
			},
			want: nil,
		},
		{
			name: "single item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue(*testObservers[0].ID),
				}),
			},
			want: []*string{
				testObservers[0].ID,
			},
		},
		{
			name: "multiple items",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue(*testObservers[0].ID),
					types.StringValue(*testObservers[1].ID),
				}),
			},
			want: []*string{
				testObservers[0].ID,
				testObservers[1].ID,
			},
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandStringPointerSliceList(&thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestExpandStringSliceList(t *testing.T) {
	testObservers := []*xmatters.Role{
		{
			ID: RandUUIDPointer(),
		},
		{
			ID: RandUUIDPointer(),
		},
	}

	type args struct {
		diags diag.Diagnostics
		in    basetypes.ListValue
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.ListNull(types.StringType),
			},
			want: nil,
		},
		{
			name: "invalid item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.Int32Type, []attr.Value{
					types.Int32Value(5),
				}),
			},
			want: nil,
		},
		{
			name: "single item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue(*testObservers[0].ID),
				}),
			},
			want: []string{
				*testObservers[0].ID,
			},
		},
		{
			name: "multiple items",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue(*testObservers[0].ID),
					types.StringValue(*testObservers[1].ID),
				}),
			},
			want: []string{
				*testObservers[0].ID,
				*testObservers[1].ID,
			},
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandStringSliceList(&thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestExpandStringList(t *testing.T) {
	testStrings := RandStringList(2, 5)
	type args struct {
		diags diag.Diagnostics
		in    basetypes.ListValue
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.ListNull(types.StringType),
			},
			want: "",
		},
		{
			name: "invalid item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.Int32Type, []attr.Value{
					types.Int32Value(5),
				}),
			},
			want: "",
		},
		{
			name: "single item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue(testStrings[0]),
				}),
			},
			want: testStrings[0],
		},
		{
			name: "multiple items",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue(testStrings[0]),
					types.StringValue(testStrings[1]),
				}),
			},
			want: fmt.Sprintf("%s,%s", testStrings[0], testStrings[1]),
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandStringList(&thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestExpandEncodedStringList(t *testing.T) {
	testStrings := []string{"group 1", "group,2", "group3"}
	type args struct {
		diags diag.Diagnostics
		in    basetypes.ListValue
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.ListNull(types.StringType),
			},
			want: "",
		},
		{
			name: "invalid item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.Int32Type, []attr.Value{
					types.Int32Value(5),
				}),
			},
			want: "",
		},
		{
			name: "single item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue(testStrings[0]),
				}),
			},
			want: "group+1",
		},
		{
			name: "multiple items",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue(testStrings[0]),
					types.StringValue(testStrings[1]),
					types.StringValue(testStrings[2]),
				}),
			},
			want: "group+1,group%2C2,group3",
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandEncodedStringList(&thisTest.args.diags, thisTest.args.in)

			assert.Equal(t, thisTest.want, got)
		})
	}
}

// ------------------------------------------------------------
// Set Expanders
// ------------------------------------------------------------

func TestExpandServiceLinkSet(t *testing.T) {
	testServiceLinks := []*xmatters.ServiceLink{
		{
			Label: RandStringPointer(5),
			URL:   RandStringPointer(5),
		},
		{
			Label: RandStringPointer(5),
			URL:   RandStringPointer(5),
		},
	}
	type args struct {
		diags diag.Diagnostics
		in    basetypes.SetValue
	}
	tests := []struct {
		name string
		args args
		want []*xmatters.ServiceLink
	}{
		{
			name: "empty list",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.SetNull(ServiceLinkObjectType),
			},
			want: []*xmatters.ServiceLink{},
		},
		{
			name: "invalid item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(types.Int32Type, []attr.Value{
					types.Int32Value(5),
				}),
			},
			want: nil,
		},
		{
			name: "single item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(ServiceLinkObjectType, []attr.Value{
					types.ObjectValueMust(ServiceLinkObjectType.AttrTypes, map[string]attr.Value{
						"link_text": customTypes.StringValue(*testServiceLinks[0].Label),
						"url":       types.StringValue(*testServiceLinks[0].URL),
					}),
				}),
			},
			want: []*xmatters.ServiceLink{
				testServiceLinks[0],
			},
		},
		{
			name: "multiple items",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(ServiceLinkObjectType, []attr.Value{
					types.ObjectValueMust(ServiceLinkObjectType.AttrTypes, map[string]attr.Value{
						"link_text": customTypes.StringValue(*testServiceLinks[0].Label),
						"url":       types.StringValue(*testServiceLinks[0].URL),
					}),
					types.ObjectValueMust(ServiceLinkObjectType.AttrTypes, map[string]attr.Value{
						"link_text": customTypes.StringValue(*testServiceLinks[1].Label),
						"url":       types.StringValue(*testServiceLinks[1].URL),
					}),
				}),
			},
			want: testServiceLinks,
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandServiceLinkSet(&thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestExpandStringSliceSet(t *testing.T) {
	testStrings := []*string{
		RandStringPointer(5),
		RandStringPointer(5),
	}

	type args struct {
		diags diag.Diagnostics
		in    basetypes.SetValue
	}
	tests := []struct {
		name string
		args args
		want []*string
	}{
		{
			name: "empty",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.SetNull(types.StringType),
			},
			want: nil,
		},
		{
			name: "invalid item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(types.Int32Type, []attr.Value{
					types.Int32Value(5),
				}),
			},
			want: nil,
		},
		{
			name: "single item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(types.StringType, []attr.Value{
					types.StringValue(*testStrings[0]),
				}),
			},
			want: []*string{
				testStrings[0],
			},
		},
		{
			name: "multiple items",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(types.StringType, []attr.Value{
					types.StringValue(*testStrings[0]),
					types.StringValue(*testStrings[1]),
				}),
			},
			want: []*string{
				testStrings[0],
				testStrings[1],
			},
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandStringSliceSet(&thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestExpandTimeframeSet(t *testing.T) {
	testTimeframes := []*xmatters.DeviceTimeframe{
		{
			Days:              []*string{RandStringPointer(3), RandStringPointer(3)},
			DurationInMinutes: helpers.Int32Pointer(60),
			ExcludeHolidays:   RandBoolPointer(),
			Name:              RandStringPointer(5),
			StartTime:         RandStringPointer(5),
		},
		{
			Days:              []*string{RandStringPointer(3)},
			DurationInMinutes: helpers.Int32Pointer(120),
			ExcludeHolidays:   RandBoolPointer(),
			Name:              RandStringPointer(5),
			StartTime:         RandStringPointer(5),
		},
	}
	type args struct {
		diags diag.Diagnostics
		in    basetypes.SetValue
	}
	tests := []struct {
		name string
		args args
		want []*xmatters.DeviceTimeframe
	}{
		{
			name: "empty",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.SetNull(TimeframeObjectType),
			},
			want: nil,
		},
		{
			name: "invalid item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(types.Int32Type, []attr.Value{
					types.Int32Value(5),
				}),
			},
			want: nil,
		},
		{
			name: "single item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(TimeframeObjectType, []attr.Value{
					types.ObjectValueMust(TimeframeObjectType.AttrTypes, map[string]attr.Value{
						"days":                types.SetValueMust(customTypes.CustomStringType{}, []attr.Value{customTypes.StringValue(*testTimeframes[0].Days[0]), customTypes.StringValue(*testTimeframes[0].Days[1])}),
						"duration_in_minutes": types.Int32PointerValue(testTimeframes[0].DurationInMinutes),
						"exclude_holidays":    types.BoolPointerValue(testTimeframes[0].ExcludeHolidays),
						"name":                customTypes.StringPointerValue(testTimeframes[0].Name),
						"start_time":          types.StringPointerValue(testTimeframes[0].StartTime),
					}),
				}),
			},
			want: []*xmatters.DeviceTimeframe{
				testTimeframes[0],
			},
		},
		{
			name: "multiple items",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(TimeframeObjectType, []attr.Value{
					types.ObjectValueMust(TimeframeObjectType.AttrTypes, map[string]attr.Value{
						"days":                types.SetValueMust(customTypes.CustomStringType{}, []attr.Value{customTypes.StringValue(*testTimeframes[0].Days[0]), customTypes.StringValue(*testTimeframes[0].Days[1])}),
						"duration_in_minutes": types.Int32PointerValue(testTimeframes[0].DurationInMinutes),
						"exclude_holidays":    types.BoolPointerValue(testTimeframes[0].ExcludeHolidays),
						"name":                customTypes.StringPointerValue(testTimeframes[0].Name),
						"start_time":          types.StringPointerValue(testTimeframes[0].StartTime),
					}),
					types.ObjectValueMust(TimeframeObjectType.AttrTypes, map[string]attr.Value{
						"days":                types.SetValueMust(customTypes.CustomStringType{}, []attr.Value{customTypes.StringValue(*testTimeframes[1].Days[0])}),
						"duration_in_minutes": types.Int32PointerValue(testTimeframes[1].DurationInMinutes),
						"exclude_holidays":    types.BoolPointerValue(testTimeframes[1].ExcludeHolidays),
						"name":                customTypes.StringPointerValue(testTimeframes[1].Name),
						"start_time":          types.StringPointerValue(testTimeframes[1].StartTime),
					}),
				}),
			},
			want: testTimeframes,
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandTimeframeSet(&thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestExpandReferenceNameSet(t *testing.T) {
	testRefNames := []*xmatters.ReferenceByName{
		{
			Name: RandStringPointer(5),
		},
		{
			Name: RandStringPointer(5),
		},
	}
	type args struct {
		diags diag.Diagnostics
		in    basetypes.SetValue
	}
	tests := []struct {
		name string
		args args
		want []*xmatters.ReferenceByName
	}{
		{
			name: "empty",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.SetNull(customTypes.CustomStringType{}),
			},
			want: nil,
		},
		{
			name: "invalid item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(types.Int32Type, []attr.Value{
					types.Int32Value(5),
				}),
			},
			want: nil,
		},
		{
			name: "single item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(customTypes.CustomStringType{}, []attr.Value{
					customTypes.StringValue(*testRefNames[0].Name),
				}),
			},
			want: []*xmatters.ReferenceByName{
				testRefNames[0],
			},
		},
		{
			name: "multiple items",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(customTypes.CustomStringType{}, []attr.Value{
					customTypes.StringValue(*testRefNames[0].Name),
					customTypes.StringValue(*testRefNames[1].Name),
				}),
			},
			want: []*xmatters.ReferenceByName{
				testRefNames[0],
				testRefNames[1],
			},
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandReferenceNameSet(&thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestExpandReferenceIDSet(t *testing.T) {
	testRefIds := []*xmatters.ReferenceById{
		{
			ID: RandStringPointer(5),
		},
		{
			ID: RandStringPointer(5),
		},
	}
	type args struct {
		diags diag.Diagnostics
		in    basetypes.SetValue
	}
	tests := []struct {
		name string
		args args
		want []*xmatters.ReferenceById
	}{
		{
			name: "empty",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.SetNull(types.StringType),
			},
			want: nil,
		},
		{
			name: "invalid item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(types.Int32Type, []attr.Value{
					types.Int32Value(5),
				}),
			},
			want: nil,
		},
		{
			name: "single item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(types.StringType, []attr.Value{
					types.StringValue(*testRefIds[0].ID),
				}),
			},
			want: []*xmatters.ReferenceById{
				testRefIds[0],
			},
		},
		{
			name: "multiple items",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(types.StringType, []attr.Value{
					types.StringValue(*testRefIds[0].ID),
					types.StringValue(*testRefIds[1].ID),
				}),
			},
			want: []*xmatters.ReferenceById{
				testRefIds[0],
				testRefIds[1],
			},
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandReferenceIDSet(&thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestExpandGroupMemberSet(t *testing.T) {
	testGroupMembers := []*xmatters.GroupMember{
		{
			ID:         RandStringPointer(5),
			MemberType: helpers.StringPointer("PERSON"),
		},
		{
			ID:         RandStringPointer(5),
			MemberType: helpers.StringPointer("DEVICE"),
		},
		{
			ID:         RandStringPointer(5),
			MemberType: helpers.StringPointer("GROUP"),
		},
	}
	type args struct {
		diags diag.Diagnostics
		in    basetypes.SetValue
	}
	tests := []struct {
		name string
		args args
		want []*xmatters.GroupMember
	}{
		{
			name: "empty list",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.SetNull(GroupMemberObjectType),
			},
			want: nil,
		},
		{
			name: "invalid item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(types.Int32Type, []attr.Value{
					types.Int32Value(5),
				}),
			},
			want: nil,
		},
		{
			name: "single item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(GroupMemberObjectType, []attr.Value{
					types.ObjectValueMust(GroupMemberObjectType.AttrTypes, map[string]attr.Value{
						"id":          types.StringValue(*testGroupMembers[0].ID),
						"member_type": types.StringValue(*testGroupMembers[0].MemberType),
					}),
				}),
			},
			want: []*xmatters.GroupMember{
				testGroupMembers[0],
			},
		},
		{
			name: "multiple items",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(GroupMemberObjectType, []attr.Value{
					types.ObjectValueMust(GroupMemberObjectType.AttrTypes, map[string]attr.Value{
						"id":          types.StringValue(*testGroupMembers[0].ID),
						"member_type": types.StringValue(*testGroupMembers[0].MemberType),
					}),
					types.ObjectValueMust(GroupMemberObjectType.AttrTypes, map[string]attr.Value{
						"id":          types.StringValue(*testGroupMembers[1].ID),
						"member_type": types.StringValue(*testGroupMembers[1].MemberType),
					}),
					types.ObjectValueMust(GroupMemberObjectType.AttrTypes, map[string]attr.Value{
						"id":          types.StringValue(*testGroupMembers[2].ID),
						"member_type": types.StringValue(*testGroupMembers[2].MemberType),
					}),
				}),
			},
			want: testGroupMembers,
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandGroupMemberSet(&thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestExpandGroupCriterionSet(t *testing.T) {
	testCriteria := []*xmatters.SearchCriterion{
		{
			CriterionType: RandStringPointer(5),
			Field:         RandStringPointer(5),
			Operand:       RandStringPointer(3),
			Value:         RandStringPointer(5),
		},
		{
			CriterionType: RandStringPointer(5),
			Field:         RandStringPointer(5),
			Operand:       RandStringPointer(3),
			Value:         RandStringPointer(5),
		},
	}
	type args struct {
		diags diag.Diagnostics
		in    basetypes.SetValue
	}
	tests := []struct {
		name string
		args args
		want []*xmatters.SearchCriterion
	}{
		{
			name: "empty",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.SetNull(GroupCriterionObjectType),
			},
			want: nil,
		},
		{
			name: "invalid item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(types.Int32Type, []attr.Value{
					types.Int32Value(5),
				}),
			},
			want: nil,
		},
		{
			name: "single item",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(GroupCriterionObjectType, []attr.Value{
					types.ObjectValueMust(GroupCriterionObjectType.AttrTypes, map[string]attr.Value{
						"criterion_type": types.StringPointerValue(testCriteria[0].CriterionType),
						"field":          types.StringPointerValue(testCriteria[0].Field),
						"operand":        customTypes.StringPointerValue(testCriteria[0].Operand),
						"value":          customTypes.StringPointerValue(testCriteria[0].Value),
					}),
				}),
			},
			want: []*xmatters.SearchCriterion{
				testCriteria[0],
			},
		},
		{
			name: "multiple items",
			args: args{
				diags: diag.Diagnostics{},
				in: types.SetValueMust(GroupCriterionObjectType, []attr.Value{
					types.ObjectValueMust(GroupCriterionObjectType.AttrTypes, map[string]attr.Value{
						"criterion_type": types.StringPointerValue(testCriteria[0].CriterionType),
						"field":          types.StringPointerValue(testCriteria[0].Field),
						"operand":        customTypes.StringPointerValue(testCriteria[0].Operand),
						"value":          customTypes.StringPointerValue(testCriteria[0].Value),
					}),
					types.ObjectValueMust(GroupCriterionObjectType.AttrTypes, map[string]attr.Value{
						"criterion_type": types.StringPointerValue(testCriteria[1].CriterionType),
						"field":          types.StringPointerValue(testCriteria[1].Field),
						"operand":        customTypes.StringPointerValue(testCriteria[1].Operand),
						"value":          customTypes.StringPointerValue(testCriteria[1].Value),
					}),
				}),
			},
			want: testCriteria,
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandGroupCriterionSet(&thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

// ------------------------------------------------------------
// Primitive Expanders
// ------------------------------------------------------------

func TestExpandGroupReference(t *testing.T) {
	testGroupReference := xmatters.GroupReference{
		ID: RandUUIDPointer(),
	}
	type args struct {
		in basetypes.StringValue
	}
	tests := []struct {
		name string
		args args
		want *xmatters.GroupReference
	}{
		{
			name: "null params",
			args: args{
				in: basetypes.NewStringNull(),
			},
			want: nil,
		},
		{
			name: "unknown params",
			args: args{
				in: basetypes.NewStringUnknown(),
			},
			want: nil,
		},
		{
			name: "valid",
			args: args{
				in: types.StringPointerValue(testGroupReference.ID),
			},
			want: &testGroupReference,
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandGroupReferenceId(thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestExpandGroupCriteriaObject(t *testing.T) {
	testCriteria := &xmatters.SearchCriteria{
		Operand: RandStringPointer(3),
		Criterion: []*xmatters.SearchCriterion{
			{
				CriterionType: RandStringPointer(5),
				Field:         RandStringPointer(5),
				Operand:       RandStringPointer(3),
				Value:         RandStringPointer(5),
			},
		},
	}
	type args struct {
		diags diag.Diagnostics
		in    basetypes.ObjectValue
	}
	tests := []struct {
		name string
		args args
		want *xmatters.SearchCriteria
	}{
		{
			name: "null params",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.ObjectNull(GroupCriteriaObjectType.AttrTypes),
			},
			want: nil,
		},
		{
			name: "unknown params",
			args: args{
				diags: diag.Diagnostics{},
				in:    types.ObjectUnknown(GroupCriteriaObjectType.AttrTypes),
			},
			want: nil,
		},
		{
			name: "valid params",
			args: args{
				diags: diag.Diagnostics{},
				in: types.ObjectValueMust(GroupCriteriaObjectType.AttrTypes, map[string]attr.Value{
					"operand": types.StringPointerValue(testCriteria.Operand),
					"criterion": types.SetValueMust(GroupCriterionObjectType, []attr.Value{
						types.ObjectValueMust(GroupCriterionObjectType.AttrTypes, map[string]attr.Value{
							"criterion_type": types.StringPointerValue(testCriteria.Criterion[0].CriterionType),
							"field":          types.StringPointerValue(testCriteria.Criterion[0].Field),
							"operand":        customTypes.StringPointerValue(testCriteria.Criterion[0].Operand),
							"value":          customTypes.StringPointerValue(testCriteria.Criterion[0].Value),
						}),
					}),
				}),
			},
			want: testCriteria,
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := ExpandGroupCriteriaObject(&thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}
