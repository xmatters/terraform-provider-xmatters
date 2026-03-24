package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/groups"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

func TestGroupsAPIParams(t *testing.T) {
	testParams := xmatters.GetGroupsParams{
		Embed:             "supervisors,observers",
		Terms:             utils.RandString(5),
		Operand:           utils.RandString(5),
		Fields:            utils.RandString(5),
		CreatedAfter:      utils.RandString(5),
		CreatedBefore:     utils.RandString(5),
		CreatedFrom:       utils.RandString(5),
		CreatedTo:         utils.RandString(5),
		GroupType:         utils.RandString(5),
		MemberExists:      utils.RandString(5),
		MemberLicenseType: utils.RandString(5),
		Members:           utils.RandString(8),
		Sites:             utils.RandString(8),
		Status:            utils.RandString(5),
		Supervisors:       utils.RandString(8),
		SortBy:            utils.RandString(4),
		SortOrder:         utils.RandString(4),
	}

	type args struct {
		diags *diag.Diagnostics
		model groups.GroupsModel
	}

	tests := []struct {
		name   string
		args   args
		expect xmatters.GetGroupsParams
	}{
		{
			name: "empty model",
			args: args{
				diags: &diag.Diagnostics{},
				model: groups.GroupsModel{},
			},
			expect: xmatters.GetGroupsParams{
				Embed: "supervisors,observers",
			},
		},
		{
			name: "with search params",
			args: args{
				diags: &diag.Diagnostics{},
				model: groups.GroupsModel{
					Search: &groups.GroupsSearchModel{
						Terms:   types.StringValue(testParams.Terms),
						Operand: types.StringValue(testParams.Operand),
						Fields: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.Fields),
							},
						),
					},
				},
			},
			expect: xmatters.GetGroupsParams{
				Embed:   "supervisors,observers",
				Terms:   testParams.Terms,
				Operand: testParams.Operand,
				Fields:  testParams.Fields,
			},
		},
		{
			name: "with filter params",
			args: args{
				diags: &diag.Diagnostics{},
				model: groups.GroupsModel{
					Filters: &groups.GroupsFiltersModel{
						CreatedAfter:       types.StringValue(testParams.CreatedAfter),
						CreatedBefore:      types.StringValue(testParams.CreatedBefore),
						CreatedFrom:        types.StringValue(testParams.CreatedFrom),
						CreatedTo:          types.StringValue(testParams.CreatedTo),
						GroupType:          types.StringValue(testParams.GroupType),
						MemberExists:       types.StringValue(testParams.MemberExists),
						MembersLicenseType: types.StringValue(testParams.MemberLicenseType),
						Members: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.Members),
							},
						),
						Sites: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.Sites),
							},
						),
						Status: types.StringValue(testParams.Status),
						Supervisors: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.Supervisors),
							},
						),
					},
				},
			},
			expect: xmatters.GetGroupsParams{
				Embed:             "supervisors,observers",
				CreatedAfter:      testParams.CreatedAfter,
				CreatedBefore:     testParams.CreatedBefore,
				CreatedFrom:       testParams.CreatedFrom,
				CreatedTo:         testParams.CreatedTo,
				GroupType:         testParams.GroupType,
				MemberExists:      testParams.MemberExists,
				MemberLicenseType: testParams.MemberLicenseType,
				Members:           testParams.Members,
				Sites:             testParams.Sites,
				Status:            testParams.Status,
				Supervisors:       testParams.Supervisors,
			},
		},
		{
			name: "with multi-site names requiring encoding",
			args: args{
				diags: &diag.Diagnostics{},
				model: groups.GroupsModel{
					Filters: &groups.GroupsFiltersModel{
						Sites: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue("Default Site"),
								types.StringValue("Great White North"),
							},
						),
					},
				},
			},
			expect: xmatters.GetGroupsParams{
				Embed: "supervisors,observers",
				Sites: "Default+Site,Great+White+North",
			},
		},
		{
			name: "with options params",
			args: args{
				diags: &diag.Diagnostics{},
				model: groups.GroupsModel{
					Options: &groups.GroupsOptionsModel{
						SortBy:    types.StringValue(testParams.SortBy),
						SortOrder: types.StringValue(testParams.SortOrder),
					},
				},
			},
			expect: xmatters.GetGroupsParams{
				Embed:     "supervisors,observers",
				SortBy:    testParams.SortBy,
				SortOrder: testParams.SortOrder,
			},
		},
		{
			name: "full model",
			args: args{
				diags: &diag.Diagnostics{},
				model: groups.GroupsModel{
					Search: &groups.GroupsSearchModel{
						Terms:   types.StringValue(testParams.Terms),
						Operand: types.StringValue(testParams.Operand),
						Fields: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.Fields),
							},
						),
					},
					Filters: &groups.GroupsFiltersModel{
						CreatedAfter:       types.StringValue(testParams.CreatedAfter),
						CreatedBefore:      types.StringValue(testParams.CreatedBefore),
						CreatedFrom:        types.StringValue(testParams.CreatedFrom),
						CreatedTo:          types.StringValue(testParams.CreatedTo),
						GroupType:          types.StringValue(testParams.GroupType),
						MemberExists:       types.StringValue(testParams.MemberExists),
						MembersLicenseType: types.StringValue(testParams.MemberLicenseType),
						Members: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.Members),
							},
						),
						Sites: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.Sites),
							},
						),
						Status: types.StringValue(testParams.Status),
						Supervisors: types.ListValueMust(
							types.StringType,
							[]attr.Value{
								types.StringValue(testParams.Supervisors),
							},
						),
					},
					Options: &groups.GroupsOptionsModel{
						SortBy:    types.StringValue(testParams.SortBy),
						SortOrder: types.StringValue(testParams.SortOrder),
					},
				},
			},
			expect: testParams,
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			actual := thisTest.args.model.APIParams(thisTest.args.diags)
			assert.Equal(t, thisTest.expect, actual)
		})
	}
}
