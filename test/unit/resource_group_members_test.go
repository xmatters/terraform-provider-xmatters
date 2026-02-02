package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/resources/groupMembers"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

func TestResourceGroupMembersToModel(t *testing.T) {
	testRoster := xmatters.GroupMembers{
		ID: utils.RandUUIDPointer(),
		Group: &xmatters.GroupReference{
			ID:            utils.RandUUIDPointer(),
			TargetName:    utils.RandStringPointer(5),
			RecipientType: utils.RandStringPointer(5),
			GroupType:     utils.RandStringPointer(5),
		},
		Members: []*xmatters.GroupMember{
			{
				ID:         utils.RandUUIDPointer(),
				MemberType: utils.RandStringPointer(5),
			},
			{
				ID:         utils.RandUUIDPointer(),
				MemberType: utils.RandStringPointer(5),
			},
			{
				ID:         utils.RandUUIDPointer(),
				MemberType: utils.RandStringPointer(5),
			},
		},
	}
	type args struct {
		diags *diag.Diagnostics
		in    xmatters.GroupMembers
	}

	tests := []struct {
		name     string
		args     args
		expected groupMembers.GroupMembersModel
	}{
		{
			name: "empty roster",
			args: args{
				diags: &diag.Diagnostics{},
				in:    xmatters.GroupMembers{},
			},
			expected: groupMembers.GroupMembersModel{
				Group: types.ObjectNull(
					utils.GroupReferenceObjectType.AttrTypes,
				),
				Members: types.SetValueMust(
					utils.GroupMemberObjectType,
					[]attr.Value{},
				),
			},
		},
		{
			name: "full roster",
			args: args{
				diags: &diag.Diagnostics{},
				in:    testRoster,
			},
			expected: groupMembers.GroupMembersModel{
				ID: types.StringPointerValue(testRoster.ID),
				Group: types.ObjectValueMust(
					utils.GroupReferenceObjectType.AttrTypes,
					map[string]attr.Value{
						"id":             types.StringPointerValue(testRoster.Group.ID),
						"name":           types.StringPointerValue(testRoster.Group.TargetName),
						"recipient_type": types.StringPointerValue(testRoster.Group.RecipientType),
						"group_type":     types.StringPointerValue(testRoster.Group.GroupType),
					},
				),
				Members: types.SetValueMust(
					utils.GroupMemberObjectType,
					[]attr.Value{
						types.ObjectValueMust(
							utils.GroupMemberObjectType.AttrTypes,
							map[string]attr.Value{
								"id":          types.StringPointerValue(testRoster.Members[0].ID),
								"member_type": types.StringPointerValue(testRoster.Members[0].MemberType),
							},
						),
						types.ObjectValueMust(
							utils.GroupMemberObjectType.AttrTypes,
							map[string]attr.Value{
								"id":          types.StringPointerValue(testRoster.Members[1].ID),
								"member_type": types.StringPointerValue(testRoster.Members[1].MemberType),
							},
						),
						types.ObjectValueMust(
							utils.GroupMemberObjectType.AttrTypes,
							map[string]attr.Value{
								"id":          types.StringPointerValue(testRoster.Members[2].ID),
								"member_type": types.StringPointerValue(testRoster.Members[2].MemberType),
							},
						),
					},
				),
			},
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			actual := groupMembers.GroupMembersToModel(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.expected, actual)
		})
	}
}
