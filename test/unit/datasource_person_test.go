package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/person"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

func TestPersonToModel(t *testing.T) {
	testPerson := xmatters.Person{
		ID:         utils.RandUUIDPointer(),
		TargetName: utils.RandStringPointer(5),
		FirstName:  utils.RandStringPointer(5),
		LastName:   utils.RandStringPointer(5),
		Roles: []*xmatters.Role{
			{
				ID:   utils.RandUUIDPointer(),
				Name: utils.RandStringPointer(5),
			},
			{
				ID:   utils.RandUUIDPointer(),
				Name: utils.RandStringPointer(5),
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
			{
				ID: utils.RandUUIDPointer(),
			},
		},
		PhoneLogin:      utils.RandStringPointer(5),
		LicenseType:     utils.RandStringPointer(5),
		ExternalKey:     utils.RandStringPointer(5),
		ExternallyOwned: utils.RandBoolPointer(),
		LastLogin:       utils.RandStringPointer(5),
	}
	type args struct {
		diags *diag.Diagnostics
		in    xmatters.Person
	}
	tests := []struct {
		name     string
		args     args
		expected person.PersonModel
	}{
		{
			name: "empty person",
			args: args{
				diags: &diag.Diagnostics{},
				in:    xmatters.Person{},
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
				diags: &diag.Diagnostics{},
				in:    testPerson,
			},
			expected: person.PersonModel{
				ID:         types.StringPointerValue(testPerson.ID),
				TargetName: types.StringPointerValue(testPerson.TargetName),
				FirstName:  types.StringPointerValue(testPerson.FirstName),
				LastName:   types.StringPointerValue(testPerson.LastName),
				Roles: types.SetValueMust(
					customTypes.CustomStringType{},
					[]attr.Value{
						customTypes.StringPointerValue(testPerson.Roles[0].Name),
						customTypes.StringPointerValue(testPerson.Roles[1].Name),
					},
				),
				Status:   types.StringPointerValue(testPerson.Status),
				WebLogin: types.StringPointerValue(testPerson.WebLogin),
				Site:     types.StringPointerValue(testPerson.Site.ID),
				Timezone: types.StringPointerValue(testPerson.Timezone),
				Language: types.StringPointerValue(testPerson.Language),
				Supervisors: types.SetValueMust(
					types.StringType,
					[]attr.Value{
						types.StringPointerValue(testPerson.Supervisors[0].ID),
						types.StringPointerValue(testPerson.Supervisors[1].ID),
					},
				),
				PhoneLogin:      types.StringPointerValue(testPerson.PhoneLogin),
				LicenseType:     types.StringPointerValue(testPerson.LicenseType),
				ExternalKey:     types.StringPointerValue(testPerson.ExternalKey),
				ExternallyOwned: types.BoolPointerValue(testPerson.ExternallyOwned),
				LastLogin:       types.StringPointerValue(testPerson.LastLogin),
			},
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := person.PersonToModel(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.expected, got)
		})
	}
}
