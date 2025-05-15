package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/people"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

// GetPeopleParams
func TestGetPeopleParams(t *testing.T) {
	testParams := xmatters.GetPeopleParams{
		Terms:              utils.RandString(5),
		Operand:            utils.RandString(5),
		Fields:             utils.RandString(5),
		CreatedFrom:        utils.RandString(5),
		CreatedTo:          utils.RandString(5),
		CreatedBefore:      utils.RandString(5),
		CreatedAfter:       utils.RandString(5),
		DevicesExists:      utils.RandBoolPointer(),
		DevicesEmailExists: utils.RandBoolPointer(),
		DevicesFailsafe:    utils.RandBoolPointer(),
		DevicesMobile:      utils.RandBoolPointer(),
		DevicesSMS:         utils.RandBoolPointer(),
		DevicesVoice:       utils.RandBoolPointer(),
		DevicesStatus:      utils.RandString(5),
		DevicesTestStatus:  utils.RandString(5),
		EmailAddress:       utils.RandString(5),
		FirstName:          utils.RandString(5),
		Groups:             utils.RandString(5),
		GroupsExists:       utils.RandBoolPointer(),
		LastName:           utils.RandString(5),
		LicenseType:        utils.RandString(5),
		PhoneNumber:        utils.RandString(5),
		Roles:              utils.RandString(5),
		Site:               utils.RandString(5),
		Status:             utils.RandString(5),
		Supervisors:        utils.RandString(5),
		SupervisorsExists:  utils.RandBoolPointer(),
		TargetName:         utils.RandString(5),
		WebLogin:           utils.RandString(5),
		SortBy:             utils.RandString(5),
		SortOrder:          utils.RandString(5),
	}
	tests := []struct {
		name   string
		model  people.PeopleModel
		expect xmatters.GetPeopleParams
	}{
		{
			name:  "Empty model",
			model: people.PeopleModel{},
			expect: xmatters.GetPeopleParams{
				Embed: "roles,supervisors",
			},
		},
		{
			name: "With search fields",
			model: people.PeopleModel{
				Search: &people.PeopleSearchModel{
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
			expect: xmatters.GetPeopleParams{
				Embed:   "roles,supervisors",
				Terms:   testParams.Terms,
				Operand: testParams.Operand,
				Fields:  testParams.Fields,
			},
		},
		{
			name: "With filters",
			model: people.PeopleModel{
				Filters: &people.PeopleFiltersModel{
					CreatedFrom:        types.StringValue(testParams.CreatedFrom),
					CreatedTo:          types.StringValue(testParams.CreatedTo),
					CreatedBefore:      types.StringValue(testParams.CreatedBefore),
					CreatedAfter:       types.StringValue(testParams.CreatedAfter),
					DevicesExists:      types.BoolPointerValue(testParams.DevicesExists),
					DevicesEmailExists: types.BoolPointerValue(testParams.DevicesEmailExists),
					DevicesFailsafe:    types.BoolPointerValue(testParams.DevicesFailsafe),
					DevicesMobile:      types.BoolPointerValue(testParams.DevicesMobile),
					DevicesSMS:         types.BoolPointerValue(testParams.DevicesSMS),
					DevicesVoice:       types.BoolPointerValue(testParams.DevicesVoice),
					DevicesStatus:      types.StringValue(testParams.DevicesStatus),
					DevicesTestStatus:  types.StringValue(testParams.DevicesTestStatus),
					EmailAddress:       types.StringValue(testParams.EmailAddress),
					FirstName:          types.StringValue(testParams.FirstName),
					Groups: types.ListValueMust(types.StringType,
						[]attr.Value{
							types.StringValue(testParams.Groups),
						},
					),
					GroupsExists: types.BoolPointerValue(testParams.GroupsExists),
					LastName:     types.StringValue(testParams.LastName),
					LicenseType:  types.StringValue(testParams.LicenseType),
					PhoneNumber:  types.StringValue(testParams.PhoneNumber),
					Roles: types.ListValueMust(types.StringType,
						[]attr.Value{
							types.StringValue(testParams.Roles),
						},
					),
					Site:   types.StringValue(testParams.Site),
					Status: types.StringValue(testParams.Status),
					Supervisors: types.ListValueMust(types.StringType,
						[]attr.Value{
							types.StringValue(testParams.Supervisors),
						},
					),
					SupervisorsExists: types.BoolPointerValue(testParams.SupervisorsExists),
					TargetName:        types.StringValue(testParams.TargetName),
					WebLogin:          types.StringValue(testParams.WebLogin),
				},
			},
			expect: xmatters.GetPeopleParams{
				Embed:              "roles,supervisors",
				CreatedFrom:        testParams.CreatedFrom,
				CreatedTo:          testParams.CreatedTo,
				CreatedBefore:      testParams.CreatedBefore,
				CreatedAfter:       testParams.CreatedAfter,
				DevicesExists:      testParams.DevicesExists,
				DevicesEmailExists: testParams.DevicesEmailExists,
				DevicesFailsafe:    testParams.DevicesFailsafe,
				DevicesMobile:      testParams.DevicesMobile,
				DevicesSMS:         testParams.DevicesSMS,
				DevicesVoice:       testParams.DevicesVoice,
				DevicesStatus:      testParams.DevicesStatus,
				DevicesTestStatus:  testParams.DevicesTestStatus,
				EmailAddress:       testParams.EmailAddress,
				FirstName:          testParams.FirstName,
				Groups:             testParams.Groups,
				GroupsExists:       testParams.GroupsExists,
				LastName:           testParams.LastName,
				LicenseType:        testParams.LicenseType,
				PhoneNumber:        testParams.PhoneNumber,
				Roles:              testParams.Roles,
				Site:               testParams.Site,
				Status:             testParams.Status,
				Supervisors:        testParams.Supervisors,
				SupervisorsExists:  testParams.SupervisorsExists,
				TargetName:         testParams.TargetName,
				WebLogin:           testParams.WebLogin,
			},
		},
		{
			name: "With options",
			model: people.PeopleModel{
				Options: &people.PeopleOptionsModel{
					SortBy:    types.StringValue(testParams.SortBy),
					SortOrder: types.StringValue(testParams.SortOrder),
				},
			},
			expect: xmatters.GetPeopleParams{
				Embed:     "roles,supervisors",
				SortBy:    testParams.SortBy,
				SortOrder: testParams.SortOrder,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := tt.model.APIParams(&diag.Diagnostics{})
			assert.Equal(t, tt.expect, params)
		})
	}
}
