package utils

import (
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
// Object Flatteners
// ------------------------------------------------------------

func TestFlattenServiceLinkObject(t *testing.T) {
	testServiceLink := xmatters.ServiceLink{
		Label: RandStringPointer(5),
		URL:   RandStringPointer(5),
	}
	type args struct {
		diags *diag.Diagnostics
		in    *xmatters.ServiceLink
	}
	tests := []struct {
		name string
		args args
		want basetypes.ObjectValue
	}{
		{
			name: "nil params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    nil,
			},
			want: types.ObjectNull(ServiceLinkObjectType.AttrTypes),
		},
		{
			name: "valid params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    &testServiceLink,
			},
			want: types.ObjectValueMust(
				ServiceLinkObjectType.AttrTypes,
				map[string]attr.Value{
					"link_text": customTypes.StringPointerValue(testServiceLink.Label),
					"url":       types.StringPointerValue(testServiceLink.URL),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenServiceLinkObject(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenServiceObject(t *testing.T) {
	testService := xmatters.Service{
		ID:          RandUUIDPointer(),
		TargetName:  RandStringPointer(5),
		Description: RandStringPointer(5),
		ServiceType: RandStringPointer(5),
		ServiceTier: RandStringPointer(5),
		OwnedBy: &xmatters.GroupReference{
			ID: RandUUIDPointer(),
		},
		ServiceLinks: []*xmatters.ServiceLink{
			{
				Label: RandStringPointer(5),
				URL:   RandStringPointer(5),
			},
			{
				Label: RandStringPointer(5),
				URL:   RandStringPointer(5),
			},
		},
	}
	type args struct {
		diags *diag.Diagnostics
		in    *xmatters.Service
	}
	tests := []struct {
		name string
		args args
		want basetypes.ObjectValue
	}{
		{
			name: "nil params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    nil,
			},
			want: types.ObjectNull(ServiceObjectType.AttrTypes),
		},
		{
			name: "valid params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    &testService,
			},
			want: types.ObjectValueMust(
				ServiceObjectType.AttrTypes,
				map[string]attr.Value{
					"id":          types.StringPointerValue(testService.ID),
					"name":        customTypes.StringPointerValue(testService.TargetName),
					"description": customTypes.StringPointerValue(testService.Description),
					"type":        customTypes.StringPointerValue(testService.ServiceType),
					"tier":        types.StringPointerValue(testService.ServiceTier),
					"owner":       types.StringPointerValue(testService.OwnedBy.ID),
					"links": types.SetValueMust(ServiceLinkObjectType, []attr.Value{
						types.ObjectValueMust(
							ServiceLinkObjectType.AttrTypes,
							map[string]attr.Value{
								"link_text": customTypes.StringPointerValue(testService.ServiceLinks[0].Label),
								"url":       types.StringPointerValue(testService.ServiceLinks[0].URL),
							},
						),
						types.ObjectValueMust(
							ServiceLinkObjectType.AttrTypes,
							map[string]attr.Value{
								"link_text": customTypes.StringPointerValue(testService.ServiceLinks[1].Label),
								"url":       types.StringPointerValue(testService.ServiceLinks[1].URL),
							},
						),
					}),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenServiceObject(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenQuotaDetailsObject(t *testing.T) {
	testQuotaDetails := xmatters.QuotaDetails{
		Total:  RandInt64Pointer(),
		Active: RandInt64Pointer(),
		Unused: RandInt64Pointer(),
	}
	type args struct {
		diags *diag.Diagnostics
		in    *xmatters.QuotaDetails
	}
	tests := []struct {
		name string
		args args
		want basetypes.ObjectValue
	}{
		{
			name: "nil params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    nil,
			},
			want: types.ObjectNull(UserQuotaDetailsObjectType.AttrTypes),
		},
		{
			name: "valid params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    &testQuotaDetails,
			},
			want: types.ObjectValueMust(
				UserQuotaDetailsObjectType.AttrTypes,
				map[string]attr.Value{
					"total":  types.Int64PointerValue(testQuotaDetails.Total),
					"active": types.Int64PointerValue(testQuotaDetails.Active),
					"unused": types.Int64PointerValue(testQuotaDetails.Unused),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenQuotaDetailsObject(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenGroupObject(t *testing.T) {
	testGroup := xmatters.Group{
		ID:              RandUUIDPointer(),
		TargetName:      RandStringPointer(10),
		Description:     RandStringPointer(10),
		Status:          RandStringPointer(10),
		ExternalKey:     RandStringPointer(10),
		ExternallyOwned: RandBoolPointer(),
		AllowDuplicates: RandBoolPointer(),
		Timezone:        RandStringPointer(10),
		Site: &xmatters.ReferenceById{
			ID: RandUUIDPointer(),
		},
		ObservedByAll: RandBoolPointer(),
		Observers: []*xmatters.ReferenceByName{
			{
				Name: RandUUIDPointer(),
			},
		},
		Supervisors: []*xmatters.ReferenceById{
			{
				ID: RandUUIDPointer(),
			},
		},
		Services: []*xmatters.Service{
			{
				ID:          RandUUIDPointer(),
				TargetName:  RandStringPointer(10),
				Description: RandStringPointer(10),
				ServiceType: RandStringPointer(10),
				ServiceTier: RandStringPointer(10),
				OwnedBy: &xmatters.GroupReference{
					ID: RandUUIDPointer(),
				},
				ServiceLinks: []*xmatters.ServiceLink{
					{
						Label: RandStringPointer(10),
						URL:   RandStringPointer(10),
					},
				},
			},
		},
		GroupType: RandStringPointer(10),
	}
	type args struct {
		diags *diag.Diagnostics
		in    *xmatters.Group
	}
	tests := []struct {
		name string
		args args
		want basetypes.ObjectValue
	}{
		{
			name: "nil params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    nil,
			},
			want: types.ObjectNull(GroupObjectType.AttrTypes),
		},
		{
			name: "valid params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    &testGroup,
			},
			want: types.ObjectValueMust(
				GroupObjectType.AttrTypes,
				map[string]attr.Value{
					"id":               types.StringPointerValue(testGroup.ID),
					"name":             types.StringPointerValue(testGroup.TargetName),
					"description":      types.StringPointerValue(testGroup.Description),
					"status":           types.StringPointerValue(testGroup.Status),
					"external_key":     types.StringPointerValue(testGroup.ExternalKey),
					"externally_owned": types.BoolPointerValue(testGroup.ExternallyOwned),
					"allow_duplicates": types.BoolPointerValue(testGroup.AllowDuplicates),
					"timezone":         types.StringPointerValue(testGroup.Timezone),
					"site":             types.StringPointerValue(testGroup.Site.ID),
					"observed_by_all":  types.BoolPointerValue(testGroup.ObservedByAll),
					"observers": types.SetValueMust(
						customTypes.CustomStringType{},
						[]attr.Value{
							customTypes.StringPointerValue(testGroup.Observers[0].Name),
						},
					),
					"supervisors": types.SetValueMust(
						types.StringType,
						[]attr.Value{
							types.StringPointerValue(testGroup.Supervisors[0].ID),
						},
					),
					"services": types.SetValueMust(
						ServiceObjectType,
						[]attr.Value{
							types.ObjectValueMust(
								ServiceObjectType.AttrTypes,
								map[string]attr.Value{
									"id":          types.StringPointerValue(testGroup.Services[0].ID),
									"name":        customTypes.StringPointerValue(testGroup.Services[0].TargetName),
									"description": customTypes.StringPointerValue(testGroup.Services[0].Description),
									"type":        customTypes.StringPointerValue(testGroup.Services[0].ServiceType),
									"tier":        types.StringPointerValue(testGroup.Services[0].ServiceTier),
									"owner":       types.StringPointerValue(testGroup.Services[0].OwnedBy.ID),
									"links": types.SetValueMust(
										ServiceLinkObjectType,
										[]attr.Value{
											types.ObjectValueMust(
												ServiceLinkObjectType.AttrTypes,
												map[string]attr.Value{
													"link_text": customTypes.StringPointerValue(testGroup.Services[0].ServiceLinks[0].Label),
													"url":       types.StringPointerValue(testGroup.Services[0].ServiceLinks[0].URL),
												},
											),
										},
									),
								},
							),
						},
					),
					"group_type": types.StringPointerValue(testGroup.GroupType),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenGroupObject(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenGroupMemberObject(t *testing.T) {
	testGroupMember := xmatters.GroupMember{
		ID:         RandUUIDPointer(),
		MemberType: helpers.StringPointer("PERSON"),
	}
	type args struct {
		diags *diag.Diagnostics
		in    *xmatters.GroupMember
	}
	tests := []struct {
		name string
		args args
		want basetypes.ObjectValue
	}{
		{
			name: "nil params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    nil,
			},
			want: types.ObjectNull(GroupMemberObjectType.AttrTypes),
		},
		{
			name: "valid params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    &testGroupMember,
			},
			want: types.ObjectValueMust(
				GroupMemberObjectType.AttrTypes,
				map[string]attr.Value{
					"id":          types.StringPointerValue(testGroupMember.ID),
					"member_type": types.StringPointerValue(testGroupMember.MemberType),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenGroupMemberObject(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenGroupReferenceObject(t *testing.T) {
	testGroupReference := xmatters.GroupReference{
		ID:            RandUUIDPointer(),
		TargetName:    RandStringPointer(5),
		RecipientType: RandStringPointer(5),
		GroupType:     RandStringPointer(5),
	}
	type args struct {
		diags *diag.Diagnostics
		in    *xmatters.GroupReference
	}
	tests := []struct {
		name string
		args args
		want basetypes.ObjectValue
	}{
		{
			name: "nil params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    nil,
			},
			want: types.ObjectNull(GroupReferenceObjectType.AttrTypes),
		},
		{
			name: "valid params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    &testGroupReference,
			},
			want: types.ObjectValueMust(
				GroupReferenceObjectType.AttrTypes,
				map[string]attr.Value{
					"id":             types.StringPointerValue(testGroupReference.ID),
					"name":           types.StringPointerValue(testGroupReference.TargetName),
					"recipient_type": types.StringPointerValue(testGroupReference.RecipientType),
					"group_type":     types.StringPointerValue(testGroupReference.GroupType),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenGroupReferenceObject(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenPersonObject(t *testing.T) {
	testPerson := &xmatters.Person{
		ID:         RandUUIDPointer(),
		TargetName: RandStringPointer(5),
		FirstName:  RandStringPointer(5),
		LastName:   RandStringPointer(5),
		Roles: []*xmatters.Role{
			{
				Name: RandStringPointer(5),
			},
			{
				Name: RandStringPointer(5),
			},
		},
		Status:   RandStringPointer(5),
		WebLogin: RandStringPointer(5),
		Site: &xmatters.ReferenceById{
			ID: RandUUIDPointer(),
		},
		Timezone: RandStringPointer(5),
		Language: RandStringPointer(5),
		Supervisors: []*xmatters.Person{
			{
				ID: RandUUIDPointer(),
			},
			{
				ID: RandUUIDPointer(),
			},
		},
		PhoneLogin:      RandStringPointer(5),
		LicenseType:     RandStringPointer(5),
		ExternalKey:     RandStringPointer(5),
		ExternallyOwned: RandBoolPointer(),
		LastLogin:       RandStringPointer(5),
	}
	type args struct {
		diags *diag.Diagnostics
		in    *xmatters.Person
	}
	tests := []struct {
		name     string
		args     args
		expected basetypes.ObjectValue
	}{
		{
			name: "nil params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    nil,
			},
			expected: types.ObjectNull(PersonObjectType.AttrTypes),
		},
		{
			name: "valid params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    testPerson,
			},
			expected: types.ObjectValueMust(PersonObjectType.AttrTypes,
				map[string]attr.Value{
					"id":          types.StringPointerValue(testPerson.ID),
					"target_name": customTypes.StringPointerValue(testPerson.TargetName),
					"first_name":  customTypes.StringPointerValue(testPerson.FirstName),
					"last_name":   customTypes.StringPointerValue(testPerson.LastName),
					"roles": types.SetValueMust(
						customTypes.CustomStringType{},
						[]attr.Value{
							customTypes.StringPointerValue(testPerson.Roles[0].Name),
							customTypes.StringPointerValue(testPerson.Roles[1].Name),
						},
					),
					"status":    types.StringPointerValue(testPerson.Status),
					"web_login": customTypes.StringPointerValue(testPerson.WebLogin),
					"site":      types.StringPointerValue(testPerson.Site.ID),
					"timezone":  customTypes.StringPointerValue(testPerson.Timezone),
					"language":  customTypes.StringPointerValue(testPerson.Language),
					"supervisors": types.SetValueMust(
						types.StringType,
						[]attr.Value{
							types.StringPointerValue(testPerson.Supervisors[0].ID),

							types.StringPointerValue(testPerson.Supervisors[1].ID),
						},
					),
					"phone_login":      types.StringPointerValue(testPerson.PhoneLogin),
					"license_type":     customTypes.StringPointerValue(testPerson.LicenseType),
					"external_key":     customTypes.StringPointerValue(testPerson.ExternalKey),
					"externally_owned": types.BoolPointerValue(testPerson.ExternallyOwned),
					"last_login":       types.StringPointerValue(testPerson.LastLogin),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			model := FlattenPersonObject(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.expected, model)
		})
	}
}

func TestFlattenSiteObject(t *testing.T) {
	testSite := xmatters.Site{
		Address1:   RandStringPointer(5),
		Address2:   RandStringPointer(5),
		City:       RandStringPointer(5),
		Country:    RandStringPointer(5),
		ID:         RandUUIDPointer(),
		Language:   RandStringPointer(5),
		Latitude:   RandomLatitudePointer(),
		Longitude:  RandomLongitudePointer(),
		Name:       RandStringPointer(5),
		PostalCode: RandStringPointer(5),
		State:      RandStringPointer(5),
		Status:     helpers.StringPointer("ACTIVE"),
		Timezone:   RandStringPointer(5),
	}
	type args struct {
		diags *diag.Diagnostics
		in    *xmatters.Site
	}
	tests := []struct {
		name string
		args args
		want basetypes.ObjectValue
	}{
		{
			name: "nil params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    nil,
			},
			want: types.ObjectNull(SiteObjectType.AttrTypes),
		},
		{
			name: "valid params",
			args: args{
				diags: &diag.Diagnostics{},
				in:    &testSite,
			},
			want: types.ObjectValueMust(
				SiteObjectType.AttrTypes,
				map[string]attr.Value{
					"address1":    types.StringPointerValue(testSite.Address1),
					"address2":    types.StringPointerValue(testSite.Address2),
					"city":        types.StringPointerValue(testSite.City),
					"country":     types.StringPointerValue(testSite.Country),
					"id":          types.StringPointerValue(testSite.ID),
					"language":    types.StringPointerValue(testSite.Language),
					"latitude":    types.Float64PointerValue(testSite.Latitude),
					"longitude":   types.Float64PointerValue(testSite.Longitude),
					"name":        types.StringPointerValue(testSite.Name),
					"postal_code": types.StringPointerValue(testSite.PostalCode),
					"state":       types.StringPointerValue(testSite.State),
					"status":      types.StringPointerValue(testSite.Status),
					"timezone":    types.StringPointerValue(testSite.Timezone),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenSiteObject(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

// ------------------------------------------------------------
// List Flatteners
// ------------------------------------------------------------

func TestFlattenPersonList(t *testing.T) {
	testPerson := []*xmatters.Person{
		{
			ID:         RandUUIDPointer(),
			TargetName: RandStringPointer(5),
			FirstName:  RandStringPointer(5),
			LastName:   RandStringPointer(5),
			Roles: []*xmatters.Role{
				{
					Name: RandStringPointer(5),
				},
				{
					Name: RandStringPointer(5),
				},
			},
			Status:   RandStringPointer(5),
			WebLogin: RandStringPointer(5),
			Site: &xmatters.ReferenceById{
				ID: RandUUIDPointer(),
			},
			Timezone: RandStringPointer(5),
			Language: RandStringPointer(5),
			Supervisors: []*xmatters.Person{
				{
					ID: RandUUIDPointer(),
				},
				{
					ID: RandUUIDPointer(),
				},
			},
			PhoneLogin:      RandStringPointer(5),
			LicenseType:     RandStringPointer(5),
			ExternalKey:     RandStringPointer(5),
			ExternallyOwned: RandBoolPointer(),
			LastLogin:       RandStringPointer(5),
		},
		{
			ID:         RandUUIDPointer(),
			TargetName: RandStringPointer(5),
			FirstName:  RandStringPointer(5),
			LastName:   RandStringPointer(5),
			Roles: []*xmatters.Role{
				{
					Name: RandStringPointer(5),
				},
				{
					Name: RandStringPointer(5),
				},
			},
			Status:   RandStringPointer(5),
			WebLogin: RandStringPointer(5),
			Site: &xmatters.ReferenceById{
				ID: RandUUIDPointer(),
			},
			Timezone: RandStringPointer(5),
			Language: RandStringPointer(5),
			Supervisors: []*xmatters.Person{
				{
					ID: RandUUIDPointer(),
				},
				{
					ID: RandUUIDPointer(),
				},
			},
			PhoneLogin:      RandStringPointer(5),
			LicenseType:     RandStringPointer(5),
			ExternalKey:     RandStringPointer(5),
			ExternallyOwned: RandBoolPointer(),
			LastLogin:       RandStringPointer(5),
		},
	}
	type args struct {
		diags *diag.Diagnostics
		in    []*xmatters.Person
	}
	tests := []struct {
		name     string
		args     args
		expected types.List
	}{
		{
			name: "empty person list",
			args: args{
				diags: &diag.Diagnostics{},
				in:    []*xmatters.Person{},
			},
			expected: types.ListValueMust(
				PersonObjectType,
				[]attr.Value{},
			),
		},
		{
			name: "valid person list",
			args: args{
				diags: &diag.Diagnostics{},
				in:    testPerson,
			},
			expected: types.ListValueMust(
				PersonObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						PersonObjectType.AttrTypes,
						map[string]attr.Value{
							"id":          types.StringPointerValue(testPerson[0].ID),
							"target_name": customTypes.StringPointerValue(testPerson[0].TargetName),
							"first_name":  customTypes.StringPointerValue(testPerson[0].FirstName),
							"last_name":   customTypes.StringPointerValue(testPerson[0].LastName),
							"roles": types.SetValueMust(
								customTypes.CustomStringType{},
								[]attr.Value{
									customTypes.StringPointerValue(testPerson[0].Roles[0].Name),
									customTypes.StringPointerValue(testPerson[0].Roles[1].Name),
								},
							),
							"status":    types.StringPointerValue(testPerson[0].Status),
							"web_login": customTypes.StringPointerValue(testPerson[0].WebLogin),
							"site":      types.StringPointerValue(testPerson[0].Site.ID),
							"timezone":  customTypes.StringPointerValue(testPerson[0].Timezone),
							"language":  customTypes.StringPointerValue(testPerson[0].Language),
							"supervisors": types.SetValueMust(
								types.StringType,
								[]attr.Value{
									types.StringPointerValue(testPerson[0].Supervisors[0].ID),

									types.StringPointerValue(testPerson[0].Supervisors[1].ID),
								},
							),
							"phone_login":      types.StringPointerValue(testPerson[0].PhoneLogin),
							"license_type":     customTypes.StringPointerValue(testPerson[0].LicenseType),
							"external_key":     customTypes.StringPointerValue(testPerson[0].ExternalKey),
							"externally_owned": types.BoolPointerValue(testPerson[0].ExternallyOwned),
							"last_login":       types.StringPointerValue(testPerson[0].LastLogin),
						},
					),
					types.ObjectValueMust(
						PersonObjectType.AttrTypes,
						map[string]attr.Value{
							"id":          types.StringPointerValue(testPerson[1].ID),
							"target_name": customTypes.StringPointerValue(testPerson[1].TargetName),
							"first_name":  customTypes.StringPointerValue(testPerson[1].FirstName),
							"last_name":   customTypes.StringPointerValue(testPerson[1].LastName),
							"roles": types.SetValueMust(
								customTypes.CustomStringType{},
								[]attr.Value{
									customTypes.StringPointerValue(testPerson[1].Roles[0].Name),
									customTypes.StringPointerValue(testPerson[1].Roles[1].Name),
								},
							),
							"status":    types.StringPointerValue(testPerson[1].Status),
							"web_login": customTypes.StringPointerValue(testPerson[1].WebLogin),
							"site":      types.StringPointerValue(testPerson[1].Site.ID),
							"timezone":  customTypes.StringPointerValue(testPerson[1].Timezone),
							"language":  customTypes.StringPointerValue(testPerson[1].Language),
							"supervisors": types.SetValueMust(
								types.StringType,
								[]attr.Value{
									types.StringPointerValue(testPerson[1].Supervisors[0].ID),
									types.StringPointerValue(testPerson[1].Supervisors[1].ID),
								},
							),
							"phone_login":      types.StringPointerValue(testPerson[1].PhoneLogin),
							"license_type":     customTypes.StringPointerValue(testPerson[1].LicenseType),
							"external_key":     customTypes.StringPointerValue(testPerson[1].ExternalKey),
							"externally_owned": types.BoolPointerValue(testPerson[1].ExternallyOwned),
							"last_login":       types.StringPointerValue(testPerson[1].LastLogin),
						},
					),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			model := FlattenPersonList(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.expected, model)
		})
	}
}

func TestFlattenServiceList(t *testing.T) {
	testService := xmatters.Service{
		ID:          RandUUIDPointer(),
		TargetName:  RandStringPointer(5),
		Description: RandStringPointer(5),
		ServiceType: RandStringPointer(5),
		ServiceTier: RandStringPointer(5),
		OwnedBy: &xmatters.GroupReference{
			ID: RandUUIDPointer(),
		},
		ServiceLinks: []*xmatters.ServiceLink{
			{
				Label: RandStringPointer(5),
				URL:   RandStringPointer(5),
			},
			{
				Label: RandStringPointer(5),
				URL:   RandStringPointer(5),
			},
		},
	}
	type args struct {
		diags *diag.Diagnostics
		in    []*xmatters.Service
	}
	tests := []struct {
		name     string
		args     args
		expected types.List
	}{
		{
			name: "empty service list",
			args: args{
				diags: &diag.Diagnostics{},
				in:    []*xmatters.Service{},
			},
			expected: types.ListValueMust(
				ServiceObjectType,
				[]attr.Value{},
			),
		},
		{
			name: "valid service list",
			args: args{
				diags: &diag.Diagnostics{},
				in: []*xmatters.Service{
					&testService,
				},
			},
			expected: types.ListValueMust(
				ServiceObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						ServiceObjectType.AttrTypes,
						map[string]attr.Value{
							"id":          types.StringPointerValue(testService.ID),
							"name":        customTypes.StringPointerValue(testService.TargetName),
							"description": customTypes.StringPointerValue(testService.Description),
							"type":        customTypes.StringPointerValue(testService.ServiceType),
							"tier":        types.StringPointerValue(testService.ServiceTier),
							"owner":       types.StringPointerValue(testService.OwnedBy.ID),
							"links": types.SetValueMust(
								ServiceLinkObjectType,
								[]attr.Value{
									types.ObjectValueMust(
										ServiceLinkObjectType.AttrTypes,
										map[string]attr.Value{
											"link_text": customTypes.StringPointerValue(testService.ServiceLinks[0].Label),
											"url":       types.StringPointerValue(testService.ServiceLinks[0].URL),
										},
									),
									types.ObjectValueMust(
										ServiceLinkObjectType.AttrTypes,
										map[string]attr.Value{
											"link_text": customTypes.StringPointerValue(testService.ServiceLinks[1].Label),
											"url":       types.StringPointerValue(testService.ServiceLinks[1].URL),
										},
									),
								},
							),
						},
					),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			model := FlattenServiceList(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.expected, model)
		})
	}
}

func TestFlattenGroupList(t *testing.T) {
	testGroups := []*xmatters.Group{
		{
			ID:              RandUUIDPointer(),
			TargetName:      RandStringPointer(10),
			Description:     RandStringPointer(10),
			Status:          RandStringPointer(10),
			ExternalKey:     RandStringPointer(10),
			ExternallyOwned: RandBoolPointer(),
			AllowDuplicates: RandBoolPointer(),
			Timezone:        RandStringPointer(10),
			Site: &xmatters.ReferenceById{
				ID: RandUUIDPointer(),
			},
			ObservedByAll: RandBoolPointer(),
			Observers: []*xmatters.ReferenceByName{
				{
					Name: RandUUIDPointer(),
				},
			},
			Supervisors: []*xmatters.ReferenceById{
				{
					ID: RandUUIDPointer(),
				},
			},
			Services: []*xmatters.Service{
				{
					ID:          RandUUIDPointer(),
					TargetName:  RandStringPointer(10),
					Description: RandStringPointer(10),
					ServiceType: RandStringPointer(10),
					ServiceTier: RandStringPointer(10),
					OwnedBy: &xmatters.GroupReference{
						ID: RandUUIDPointer(),
					},
					ServiceLinks: []*xmatters.ServiceLink{
						{
							Label: RandStringPointer(10),
							URL:   RandStringPointer(10),
						},
					},
				},
			},
			GroupType: RandStringPointer(10),
		},
		{
			ID:              RandUUIDPointer(),
			TargetName:      RandStringPointer(10),
			Description:     RandStringPointer(10),
			Status:          RandStringPointer(10),
			ExternalKey:     RandStringPointer(10),
			ExternallyOwned: RandBoolPointer(),
			AllowDuplicates: RandBoolPointer(),
			Timezone:        RandStringPointer(10),
			Site: &xmatters.ReferenceById{
				ID: RandUUIDPointer(),
			},
			ObservedByAll: RandBoolPointer(),
			Observers: []*xmatters.ReferenceByName{
				{
					Name: RandUUIDPointer(),
				},
			},
			Supervisors: []*xmatters.ReferenceById{
				{
					ID: RandUUIDPointer(),
				},
			},
			Services: []*xmatters.Service{
				{
					ID:          RandUUIDPointer(),
					TargetName:  RandStringPointer(10),
					Description: RandStringPointer(10),
					ServiceType: RandStringPointer(10),
					ServiceTier: RandStringPointer(10),
					OwnedBy: &xmatters.GroupReference{
						ID: RandUUIDPointer(),
					},
					ServiceLinks: []*xmatters.ServiceLink{
						{
							Label: RandStringPointer(10),
							URL:   RandStringPointer(10),
						},
					},
				},
			},
			GroupType: RandStringPointer(10),
		},
	}
	type args struct {
		diags  *diag.Diagnostics
		groups []*xmatters.Group
	}
	tests := []struct {
		name string
		args args
		want types.List
	}{
		{
			name: "nil params",
			args: args{
				diags:  &diag.Diagnostics{},
				groups: nil,
			},
			want: types.ListValueMust(
				GroupObjectType,
				[]attr.Value{},
			),
		},
		{
			name: "valid params",
			args: args{
				diags:  &diag.Diagnostics{},
				groups: testGroups,
			},
			want: types.ListValueMust(
				GroupObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						GroupObjectType.AttrTypes,
						map[string]attr.Value{
							"id":               types.StringPointerValue(testGroups[0].ID),
							"name":             types.StringPointerValue(testGroups[0].TargetName),
							"description":      types.StringPointerValue(testGroups[0].Description),
							"status":           types.StringPointerValue(testGroups[0].Status),
							"external_key":     types.StringPointerValue(testGroups[0].ExternalKey),
							"externally_owned": types.BoolPointerValue(testGroups[0].ExternallyOwned),
							"allow_duplicates": types.BoolPointerValue(testGroups[0].AllowDuplicates),
							"timezone":         types.StringPointerValue(testGroups[0].Timezone),
							"site":             types.StringPointerValue(testGroups[0].Site.ID),
							"observed_by_all":  types.BoolPointerValue(testGroups[0].ObservedByAll),
							"observers": types.SetValueMust(
								customTypes.CustomStringType{},
								[]attr.Value{
									customTypes.StringPointerValue(testGroups[0].Observers[0].Name),
								},
							),
							"supervisors": types.SetValueMust(
								types.StringType,
								[]attr.Value{
									types.StringPointerValue(testGroups[0].Supervisors[0].ID),
								},
							),
							"services": types.SetValueMust(
								ServiceObjectType,
								[]attr.Value{
									types.ObjectValueMust(
										ServiceObjectType.AttrTypes,
										map[string]attr.Value{
											"id":          types.StringPointerValue(testGroups[0].Services[0].ID),
											"name":        customTypes.StringPointerValue(testGroups[0].Services[0].TargetName),
											"description": customTypes.StringPointerValue(testGroups[0].Services[0].Description),
											"type":        customTypes.StringPointerValue(testGroups[0].Services[0].ServiceType),
											"tier":        types.StringPointerValue(testGroups[0].Services[0].ServiceTier),
											"owner":       types.StringPointerValue(testGroups[0].Services[0].OwnedBy.ID),
											"links": types.SetValueMust(
												ServiceLinkObjectType,
												[]attr.Value{
													types.ObjectValueMust(
														ServiceLinkObjectType.AttrTypes,
														map[string]attr.Value{
															"link_text": customTypes.StringPointerValue(testGroups[0].Services[0].ServiceLinks[0].Label),
															"url":       types.StringPointerValue(testGroups[0].Services[0].ServiceLinks[0].URL),
														},
													),
												},
											),
										},
									),
								},
							),
							"group_type": types.StringPointerValue(testGroups[0].GroupType),
						},
					),
					types.ObjectValueMust(
						GroupObjectType.AttrTypes,
						map[string]attr.Value{
							"id":               types.StringPointerValue(testGroups[1].ID),
							"name":             types.StringPointerValue(testGroups[1].TargetName),
							"description":      types.StringPointerValue(testGroups[1].Description),
							"status":           types.StringPointerValue(testGroups[1].Status),
							"external_key":     types.StringPointerValue(testGroups[1].ExternalKey),
							"externally_owned": types.BoolPointerValue(testGroups[1].ExternallyOwned),
							"allow_duplicates": types.BoolPointerValue(testGroups[1].AllowDuplicates),
							"timezone":         types.StringPointerValue(testGroups[1].Timezone),
							"site":             types.StringPointerValue(testGroups[1].Site.ID),
							"observed_by_all":  types.BoolPointerValue(testGroups[1].ObservedByAll),
							"observers": types.SetValueMust(
								customTypes.CustomStringType{},
								[]attr.Value{
									customTypes.StringPointerValue(testGroups[1].Observers[0].Name),
								},
							),
							"supervisors": types.SetValueMust(
								types.StringType,
								[]attr.Value{
									types.StringPointerValue(testGroups[1].Supervisors[0].ID),
								},
							),
							"services": types.SetValueMust(
								ServiceObjectType,
								[]attr.Value{
									types.ObjectValueMust(
										ServiceObjectType.AttrTypes,
										map[string]attr.Value{
											"id":          types.StringPointerValue(testGroups[1].Services[0].ID),
											"name":        customTypes.StringPointerValue(testGroups[1].Services[0].TargetName),
											"description": customTypes.StringPointerValue(testGroups[1].Services[0].Description),
											"type":        customTypes.StringPointerValue(testGroups[1].Services[0].ServiceType),
											"tier":        types.StringPointerValue(testGroups[1].Services[0].ServiceTier),
											"owner":       types.StringPointerValue(testGroups[1].Services[0].OwnedBy.ID),
											"links": types.SetValueMust(
												ServiceLinkObjectType,
												[]attr.Value{
													types.ObjectValueMust(
														ServiceLinkObjectType.AttrTypes,
														map[string]attr.Value{
															"link_text": customTypes.StringPointerValue(testGroups[1].Services[0].ServiceLinks[0].Label),
															"url":       types.StringPointerValue(testGroups[1].Services[0].ServiceLinks[0].URL),
														},
													),
												},
											),
										},
									),
								},
							),
							"group_type": types.StringPointerValue(testGroups[1].GroupType),
						},
					),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenGroupList(thisTest.args.diags, thisTest.args.groups)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenSiteList(t *testing.T) {
	testSites := []*xmatters.Site{
		{
			Address1:   RandStringPointer(5),
			Address2:   RandStringPointer(5),
			City:       RandStringPointer(5),
			Country:    RandStringPointer(5),
			ID:         RandUUIDPointer(),
			Language:   RandStringPointer(5),
			Latitude:   RandomLatitudePointer(),
			Longitude:  RandomLongitudePointer(),
			Name:       RandStringPointer(5),
			PostalCode: RandStringPointer(5),
			State:      RandStringPointer(5),
			Status:     helpers.StringPointer("ACTIVE"),
			Timezone:   RandStringPointer(5),
		},
		{
			Address1:   RandStringPointer(5),
			Address2:   RandStringPointer(5),
			City:       RandStringPointer(5),
			Country:    RandStringPointer(5),
			ID:         RandUUIDPointer(),
			Language:   RandStringPointer(5),
			Latitude:   RandomLatitudePointer(),
			Longitude:  RandomLongitudePointer(),
			Name:       RandStringPointer(5),
			PostalCode: RandStringPointer(5),
			State:      RandStringPointer(5),
			Status:     helpers.StringPointer("ACTIVE"),
			Timezone:   RandStringPointer(5),
		},
	}
	type args struct {
		diags *diag.Diagnostics
		sites []*xmatters.Site
	}
	tests := []struct {
		name string
		args args
		want types.List
	}{
		{
			name: "nil params",
			args: args{
				diags: &diag.Diagnostics{},
				sites: nil,
			},
			want: types.ListValueMust(
				SiteObjectType,
				[]attr.Value{},
			),
		},
		{
			name: "valid params",
			args: args{
				diags: &diag.Diagnostics{},
				sites: testSites,
			},
			want: types.ListValueMust(
				SiteObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						SiteObjectType.AttrTypes,
						map[string]attr.Value{
							"address1":    types.StringPointerValue(testSites[0].Address1),
							"address2":    types.StringPointerValue(testSites[0].Address2),
							"city":        types.StringPointerValue(testSites[0].City),
							"country":     types.StringPointerValue(testSites[0].Country),
							"id":          types.StringPointerValue(testSites[0].ID),
							"language":    types.StringPointerValue(testSites[0].Language),
							"latitude":    types.Float64PointerValue(testSites[0].Latitude),
							"longitude":   types.Float64PointerValue(testSites[0].Longitude),
							"name":        types.StringPointerValue(testSites[0].Name),
							"postal_code": types.StringPointerValue(testSites[0].PostalCode),
							"state":       types.StringPointerValue(testSites[0].State),
							"status":      types.StringPointerValue(testSites[0].Status),
							"timezone":    types.StringPointerValue(testSites[0].Timezone),
						},
					),
					types.ObjectValueMust(
						SiteObjectType.AttrTypes,
						map[string]attr.Value{
							"address1":    types.StringPointerValue(testSites[1].Address1),
							"address2":    types.StringPointerValue(testSites[1].Address2),
							"city":        types.StringPointerValue(testSites[1].City),
							"country":     types.StringPointerValue(testSites[1].Country),
							"id":          types.StringPointerValue(testSites[1].ID),
							"language":    types.StringPointerValue(testSites[1].Language),
							"latitude":    types.Float64PointerValue(testSites[1].Latitude),
							"longitude":   types.Float64PointerValue(testSites[1].Longitude),
							"name":        types.StringPointerValue(testSites[1].Name),
							"postal_code": types.StringPointerValue(testSites[1].PostalCode),
							"state":       types.StringPointerValue(testSites[1].State),
							"status":      types.StringPointerValue(testSites[1].Status),
							"timezone":    types.StringPointerValue(testSites[1].Timezone),
						},
					),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenSiteList(thisTest.args.diags, thisTest.args.sites)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

// ------------------------------------------------------------
// Set Flatteners
// ------------------------------------------------------------

func TestFlattenServiceLinkSet(t *testing.T) {
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
		diags *diag.Diagnostics
		links []*xmatters.ServiceLink
	}
	tests := []struct {
		name string
		args args
		want types.Set
	}{
		{
			name: "nil params",
			args: args{
				diags: &diag.Diagnostics{},
				links: nil,
			},
			want: types.SetValueMust(
				ServiceLinkObjectType,
				[]attr.Value{},
			),
		},
		{
			name: "valid params",
			args: args{
				diags: &diag.Diagnostics{},
				links: testServiceLinks,
			},
			want: types.SetValueMust(
				ServiceLinkObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						ServiceLinkObjectType.AttrTypes,
						map[string]attr.Value{
							"link_text": customTypes.StringPointerValue(testServiceLinks[0].Label),
							"url":       types.StringPointerValue(testServiceLinks[0].URL),
						},
					),
					types.ObjectValueMust(
						ServiceLinkObjectType.AttrTypes,
						map[string]attr.Value{
							"link_text": customTypes.StringPointerValue(testServiceLinks[1].Label),
							"url":       types.StringPointerValue(testServiceLinks[1].URL),
						},
					),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenServiceLinkSet(thisTest.args.diags, thisTest.args.links)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenSupervisorSet(t *testing.T) {
	testSupervisors := []*xmatters.Person{
		{
			ID: RandUUIDPointer(),
		},
		{
			ID: RandUUIDPointer(),
		},
	}
	type args struct {
		diags       *diag.Diagnostics
		supervisors []*xmatters.Person
	}
	tests := []struct {
		name string
		args args
		want types.Set
	}{
		{
			name: "nil params",
			args: args{
				diags:       &diag.Diagnostics{},
				supervisors: nil,
			},
			want: types.SetValueMust(
				types.StringType,
				[]attr.Value{},
			),
		},
		{
			name: "valid params",
			args: args{
				diags:       &diag.Diagnostics{},
				supervisors: testSupervisors,
			},
			want: types.SetValueMust(
				types.StringType,
				[]attr.Value{
					types.StringPointerValue(testSupervisors[0].ID),
					types.StringPointerValue(testSupervisors[1].ID),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenSupervisorSet(thisTest.args.diags, thisTest.args.supervisors)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenServiceSet(t *testing.T) {
	testService := xmatters.Service{
		ID:          RandUUIDPointer(),
		TargetName:  RandStringPointer(5),
		Description: RandStringPointer(5),
		ServiceType: RandStringPointer(5),
		ServiceTier: RandStringPointer(5),
		OwnedBy: &xmatters.GroupReference{
			ID: RandUUIDPointer(),
		},
		ServiceLinks: []*xmatters.ServiceLink{
			{
				Label: RandStringPointer(5),
				URL:   RandStringPointer(5),
			},
			{
				Label: RandStringPointer(5),
				URL:   RandStringPointer(5),
			},
		},
	}
	type args struct {
		diags *diag.Diagnostics
		in    []*xmatters.Service
	}
	tests := []struct {
		name     string
		args     args
		expected types.Set
	}{
		{
			name: "empty service list",
			args: args{
				diags: &diag.Diagnostics{},
				in:    []*xmatters.Service{},
			},
			expected: types.SetValueMust(
				ServiceObjectType,
				[]attr.Value{},
			),
		},
		{
			name: "valid service list",
			args: args{
				diags: &diag.Diagnostics{},
				in: []*xmatters.Service{
					&testService,
				},
			},
			expected: types.SetValueMust(
				ServiceObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						ServiceObjectType.AttrTypes,
						map[string]attr.Value{
							"id":          types.StringPointerValue(testService.ID),
							"name":        customTypes.StringPointerValue(testService.TargetName),
							"description": customTypes.StringPointerValue(testService.Description),
							"type":        customTypes.StringPointerValue(testService.ServiceType),
							"tier":        types.StringPointerValue(testService.ServiceTier),
							"owner":       types.StringPointerValue(testService.OwnedBy.ID),
							"links": types.SetValueMust(
								ServiceLinkObjectType,
								[]attr.Value{
									types.ObjectValueMust(
										ServiceLinkObjectType.AttrTypes,
										map[string]attr.Value{
											"link_text": customTypes.StringPointerValue(testService.ServiceLinks[0].Label),
											"url":       types.StringPointerValue(testService.ServiceLinks[0].URL),
										},
									),
									types.ObjectValueMust(
										ServiceLinkObjectType.AttrTypes,
										map[string]attr.Value{
											"link_text": customTypes.StringPointerValue(testService.ServiceLinks[1].Label),
											"url":       types.StringPointerValue(testService.ServiceLinks[1].URL),
										},
									),
								},
							),
						},
					),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			model := FlattenServiceSet(thisTest.args.diags, thisTest.args.in)
			assert.Equal(t, thisTest.expected, model)
		})
	}
}

func TestFlattenRoleSet(t *testing.T) {
	testRoles := []*xmatters.Role{
		{
			Name: RandStringPointer(5),
		},
		{
			Name: RandStringPointer(5),
		},
	}
	type args struct {
		diags *diag.Diagnostics
		roles []*xmatters.Role
	}
	tests := []struct {
		name string
		args args
		want types.Set
	}{
		{
			name: "nil params",
			args: args{
				diags: &diag.Diagnostics{},
				roles: nil,
			},
			want: types.SetValueMust(
				customTypes.CustomStringType{},
				[]attr.Value{},
			),
		},
		{
			name: "valid params",
			args: args{
				diags: &diag.Diagnostics{},
				roles: testRoles,
			},
			want: types.SetValueMust(
				customTypes.CustomStringType{},
				[]attr.Value{
					customTypes.StringPointerValue(testRoles[0].Name),
					customTypes.StringPointerValue(testRoles[1].Name),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenRoleSet(thisTest.args.diags, thisTest.args.roles)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenReferenceIDSet(t *testing.T) {
	testReferences := []*xmatters.ReferenceById{
		{
			ID: RandUUIDPointer(),
		},
		{
			ID: RandUUIDPointer(),
		},
	}
	type args struct {
		diags      *diag.Diagnostics
		references []*xmatters.ReferenceById
	}
	tests := []struct {
		name string
		args args
		want types.Set
	}{
		{
			name: "nil params",
			args: args{
				diags:      &diag.Diagnostics{},
				references: nil,
			},
			want: types.SetValueMust(
				types.StringType,
				[]attr.Value{},
			),
		},
		{
			name: "valid params",
			args: args{
				diags:      &diag.Diagnostics{},
				references: testReferences,
			},
			want: types.SetValueMust(
				types.StringType,
				[]attr.Value{
					types.StringPointerValue(testReferences[0].ID),
					types.StringPointerValue(testReferences[1].ID),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenReferenceIDSet(thisTest.args.diags, thisTest.args.references)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenReferenceNameSet(t *testing.T) {
	testReferences := []*xmatters.ReferenceByName{
		{
			Name: RandStringPointer(5),
		},
		{
			Name: RandStringPointer(5),
		},
	}
	type args struct {
		diags      *diag.Diagnostics
		references []*xmatters.ReferenceByName
	}
	tests := []struct {
		name string
		args args
		want types.Set
	}{
		{
			name: "nil params",
			args: args{
				diags:      &diag.Diagnostics{},
				references: nil,
			},
			want: types.SetValueMust(
				customTypes.CustomStringType{},
				[]attr.Value{},
			),
		},
		{
			name: "valid params",
			args: args{
				diags:      &diag.Diagnostics{},
				references: testReferences,
			},
			want: types.SetValueMust(
				customTypes.CustomStringType{},
				[]attr.Value{
					customTypes.StringPointerValue(testReferences[0].Name),
					customTypes.StringPointerValue(testReferences[1].Name),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenReferenceNameSet(thisTest.args.diags, thisTest.args.references)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenGroupMemberSet(t *testing.T) {
	testMembers := []*xmatters.GroupMember{
		{
			ID:         RandUUIDPointer(),
			MemberType: helpers.StringPointer("PERSON"),
		},
		{
			ID:         RandUUIDPointer(),
			MemberType: helpers.StringPointer("PERSON"),
		},
	}
	type args struct {
		diags   *diag.Diagnostics
		members []*xmatters.GroupMember
	}
	tests := []struct {
		name string
		args args
		want types.Set
	}{
		{
			name: "nil params",
			args: args{
				diags:   &diag.Diagnostics{},
				members: nil,
			},
			want: types.SetValueMust(
				GroupMemberObjectType,
				[]attr.Value{},
			),
		},
		{
			name: "valid params",
			args: args{
				diags:   &diag.Diagnostics{},
				members: testMembers,
			},
			want: types.SetValueMust(
				GroupMemberObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						GroupMemberObjectType.AttrTypes,
						map[string]attr.Value{
							"id":          types.StringPointerValue(testMembers[0].ID),
							"member_type": types.StringPointerValue(testMembers[0].MemberType),
						},
					),
					types.ObjectValueMust(
						GroupMemberObjectType.AttrTypes,
						map[string]attr.Value{
							"id":          types.StringPointerValue(testMembers[1].ID),
							"member_type": types.StringPointerValue(testMembers[1].MemberType),
						},
					),
				},
			),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenGroupMemberSet(thisTest.args.diags, thisTest.args.members)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

// ------------------------------------------------------------
// Map Flatteners
// ------------------------------------------------------------

// ------------------------------------------------------------
// Primitive Flatteners
// ------------------------------------------------------------

func TestFlattenGroupReferenceID(t *testing.T) {
	testGroupReference := xmatters.GroupReference{
		ID: RandUUIDPointer(),
	}
	type args struct {
		in *xmatters.GroupReference
	}
	tests := []struct {
		name string
		args args
		want basetypes.StringValue
	}{
		{
			name: "nil params",
			args: args{
				in: nil,
			},
			want: types.StringNull(),
		},
		{
			name: "valid params",
			args: args{
				in: &testGroupReference,
			},
			want: types.StringPointerValue(testGroupReference.ID),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenGroupReferenceID(thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenServiceReferenceID(t *testing.T) {
	testServiceReference := xmatters.ServiceReference{
		ID: RandUUIDPointer(),
	}
	type args struct {
		in *xmatters.ServiceReference
	}
	tests := []struct {
		name string
		args args
		want basetypes.StringValue
	}{
		{
			name: "nil params",
			args: args{
				in: nil,
			},
			want: types.StringNull(),
		},
		{
			name: "valid params",
			args: args{
				in: &testServiceReference,
			},
			want: types.StringPointerValue(testServiceReference.ID),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenServiceReferenceID(thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenPersonReferenceID(t *testing.T) {
	testPersonReference := xmatters.PersonReference{
		ID: RandUUIDPointer(),
	}
	type args struct {
		in *xmatters.PersonReference
	}
	tests := []struct {
		name string
		args args
		want basetypes.StringValue
	}{
		{
			name: "nil params",
			args: args{
				in: nil,
			},
			want: types.StringNull(),
		},
		{
			name: "valid params",
			args: args{
				in: &testPersonReference,
			},
			want: types.StringPointerValue(testPersonReference.ID),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenPersonReferenceID(thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}

func TestFlattenReferenceID(t *testing.T) {
	testIDReference := xmatters.ReferenceById{
		ID: RandUUIDPointer(),
	}
	type args struct {
		in *xmatters.ReferenceById
	}
	tests := []struct {
		name string
		args args
		want basetypes.StringValue
	}{
		{
			name: "nil params",
			args: args{
				in: nil,
			},
			want: types.StringNull(),
		},
		{
			name: "valid params",
			args: args{
				in: &testIDReference,
			},
			want: types.StringPointerValue(testIDReference.ID),
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			got := FlattenReferenceID(thisTest.args.in)
			assert.Equal(t, thisTest.want, got)
		})
	}
}
