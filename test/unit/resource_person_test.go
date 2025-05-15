package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/resources/person"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

// PersonToModel
func TestResourcePersonToModel(t *testing.T) {
	testPerson := xmatters.Person{
		ID:         utils.RandUUIDPointer(),
		TargetName: utils.RandStringPointer(5),
		FirstName:  utils.RandStringPointer(5),
		LastName:   utils.RandStringPointer(5),
		Roles: []*xmatters.Role{
			{
				Name: utils.RandUUIDPointer(),
			},
			{
				Name: utils.RandUUIDPointer(),
			},
		},
		Status:   utils.RandStringPointer(5),
		WebLogin: utils.RandStringPointer(5),
		Site: &xmatters.ReferenceById{
			ID: utils.RandUUIDPointer(),
		},
		Timezone: utils.RandStringPointer(5),
		Language: utils.RandStringPointer(5),
		Supervisors: []*xmatters.Person{
			{
				ID: utils.RandUUIDPointer(),
			},
		},
		PhoneLogin:      utils.RandNumStringPointer(5),
		LicenseType:     utils.RandStringPointer(5),
		ExternalKey:     utils.RandStringPointer(5),
		ExternallyOwned: utils.RandBoolPointer(),
	}
	type args struct {
		diags  *diag.Diagnostics
		person xmatters.Person
	}
	tests := []struct {
		name     string
		args     args
		expected person.PersonModel
	}{
		{
			name: "empty person",
			args: args{
				diags:  &diag.Diagnostics{},
				person: xmatters.Person{},
			},
			expected: person.PersonModel{
				Roles: types.SetValueMust(
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
			name: "valid person",
			args: args{
				diags:  &diag.Diagnostics{},
				person: testPerson,
			},
			expected: person.PersonModel{
				ID:         types.StringPointerValue(testPerson.ID),
				TargetName: customTypes.StringPointerValue(testPerson.TargetName),
				FirstName:  customTypes.StringPointerValue(testPerson.FirstName),
				LastName:   customTypes.StringPointerValue(testPerson.LastName),
				Roles: types.SetValueMust(customTypes.CustomStringType{},
					[]attr.Value{
						customTypes.StringPointerValue(testPerson.Roles[0].Name),
						customTypes.StringPointerValue(testPerson.Roles[1].Name),
					},
				),
				Status:   types.StringPointerValue(testPerson.Status),
				WebLogin: customTypes.StringPointerValue(testPerson.WebLogin),
				Site:     types.StringPointerValue(testPerson.Site.ID),
				Timezone: types.StringPointerValue(testPerson.Timezone),
				Language: types.StringPointerValue(testPerson.Language),
				Supervisors: types.SetValueMust(types.StringType,
					[]attr.Value{
						types.StringPointerValue(testPerson.Supervisors[0].ID),
					},
				),
				PhoneLogin:      types.StringPointerValue(testPerson.PhoneLogin),
				LicenseType:     customTypes.StringPointerValue(testPerson.LicenseType),
				ExternalKey:     customTypes.StringPointerValue(testPerson.ExternalKey),
				ExternallyOwned: types.BoolPointerValue(testPerson.ExternallyOwned),
			},
		},
	}
	for _, thisTest := range tests {
		actual := person.PersonToModel(thisTest.args.diags, thisTest.args.person, thisTest.expected)
		assert.Equal(t, thisTest.expected, actual)
	}
}

// PersonParams
func TestPersonParams(t *testing.T) {
	testParams := xmatters.PushPersonParams{
		ID:              utils.RandUUID(),
		TargetName:      utils.RandString(5),
		FirstName:       utils.RandString(5),
		LastName:        utils.RandString(5),
		Roles:           []*string{utils.RandStringPointer(5), utils.RandStringPointer(5)},
		Status:          utils.RandString(5),
		WebLogin:        utils.RandString(5),
		Site:            utils.RandString(5),
		Timezone:        utils.RandString(5),
		Language:        utils.RandString(5),
		Supervisors:     []*string{utils.RandStringPointer(5)},
		PhoneLogin:      utils.RandNumStringPointer(5),
		PhonePin:        utils.RandNumString(5),
		LicenseType:     utils.RandString(5),
		ExternalKey:     utils.RandStringPointer(5),
		ExternallyOwned: utils.RandBoolPointer(),
	}
	type args struct {
		diags *diag.Diagnostics
	}
	tests := []struct {
		name     string
		args     args
		model    person.PersonModel
		expected *xmatters.PushPersonParams
	}{
		{
			name: "empty model",
			args: args{
				diags: &diag.Diagnostics{},
			},
			model:    person.PersonModel{},
			expected: &xmatters.PushPersonParams{},
		},
		{
			name: "valid model",
			args: args{
				diags: &diag.Diagnostics{},
			},
			model: person.PersonModel{
				ID:         types.StringValue(testParams.ID),
				TargetName: customTypes.StringValue(testParams.TargetName),
				FirstName:  customTypes.StringValue(testParams.FirstName),
				LastName:   customTypes.StringValue(testParams.LastName),
				Roles: types.SetValueMust(types.StringType,
					[]attr.Value{
						types.StringPointerValue(testParams.Roles[0]),
						types.StringPointerValue(testParams.Roles[1]),
					},
				),
				Status:   types.StringValue(testParams.Status),
				WebLogin: customTypes.StringValue(testParams.WebLogin),
				Site:     types.StringValue(testParams.Site),
				Timezone: types.StringValue(testParams.Timezone),
				Language: types.StringValue(testParams.Language),
				Supervisors: types.SetValueMust(types.StringType,
					[]attr.Value{
						types.StringPointerValue(testParams.Supervisors[0]),
					},
				),
				PhoneLogin:      types.StringPointerValue(testParams.PhoneLogin),
				PhonePin:        types.StringValue(testParams.PhonePin),
				LicenseType:     customTypes.StringValue(testParams.LicenseType),
				ExternalKey:     customTypes.StringPointerValue(testParams.ExternalKey),
				ExternallyOwned: types.BoolPointerValue(testParams.ExternallyOwned),
			},
			expected: &testParams,
		},
	}
	for _, thisTest := range tests {
		actual := thisTest.model.PersonParams(thisTest.args.diags)
		assert.Equal(t, thisTest.expected, actual)
	}
}
