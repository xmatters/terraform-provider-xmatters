package utils

import (
	"context"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
	"github.com/xmatters/xmatters-go"
)

// ------------------------------------------------------------
// Terraform Plugin Framework Objects
// ------------------------------------------------------------

var (
	// Define the provider Auth Config object type.
	AuthConfigObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auth_type": types.StringType,
			"username":  types.StringType,
			"password":  types.StringType,
			"token":     types.StringType,
		},
	}

	// Define the service link object type.
	ServiceLinkObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"link_text": customTypes.CustomStringType{},
			"url":       types.StringType,
		},
	}

	// Define the service object type.
	ServiceObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":          types.StringType,
			"name":        customTypes.CustomStringType{},
			"description": customTypes.CustomStringType{},
			"type":        customTypes.CustomStringType{},
			"tier":        types.StringType,
			"owner":       types.StringType,
			"links":       types.SetType{ElemType: ServiceLinkObjectType},
		},
	}

	// Define the user quota details object type.
	UserQuotaDetailsObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"total":  types.Int64Type,
			"active": types.Int64Type,
			"unused": types.Int64Type,
		},
	}

	// Define the group object type.
	GroupObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":                  types.StringType,
			"name":                types.StringType,
			"description":         types.StringType,
			"status":              types.StringType,
			"external_key":        types.StringType,
			"externally_owned":    types.BoolType,
			"allow_duplicates":    types.BoolType,
			"site":                types.StringType,
			"observed_by_all":     types.BoolType,
			"observers":           types.SetType{ElemType: customTypes.CustomStringType{}},
			"supervisors":         types.SetType{ElemType: types.StringType},
			"group_type":          types.StringType,
			"use_default_devices": types.BoolType,
			"criteria":            GroupCriteriaObjectType,
		},
	}

	// Define the group member object type.
	GroupMemberObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":          types.StringType,
			"member_type": types.StringType,
		},
	}

	// Define the group reference object type.
	GroupReferenceObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":             types.StringType,
			"name":           types.StringType,
			"recipient_type": types.StringType,
			"group_type":     types.StringType,
		},
	}

	// Define the group criteria object type.
	GroupCriteriaObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"operand":   types.StringType,
			"criterion": types.SetType{ElemType: GroupCriterionObjectType},
		},
	}

	// Define the group criterion object type.
	GroupCriterionObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"criterion_type": types.StringType,
			"field":          types.StringType,
			"operand":        types.StringType,
			"value":          types.StringType,
		},
	}

	// Define the person object type.
	PersonObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":               types.StringType,
			"target_name":      customTypes.CustomStringType{},
			"first_name":       customTypes.CustomStringType{},
			"last_name":        customTypes.CustomStringType{},
			"roles":            types.SetType{ElemType: customTypes.CustomStringType{}},
			"status":           types.StringType,
			"web_login":        customTypes.CustomStringType{},
			"site":             types.StringType,
			"timezone":         customTypes.CustomStringType{},
			"language":         customTypes.CustomStringType{},
			"supervisors":      types.SetType{ElemType: types.StringType},
			"phone_login":      types.StringType,
			"license_type":     customTypes.CustomStringType{},
			"external_key":     customTypes.CustomStringType{},
			"externally_owned": types.BoolType,
			"last_login":       types.StringType,
		},
	}

	// Define the Site object type.
	SiteObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"address1":    types.StringType,
			"address2":    types.StringType,
			"city":        types.StringType,
			"country":     types.StringType,
			"id":          types.StringType,
			"language":    types.StringType,
			"latitude":    types.Float64Type,
			"longitude":   types.Float64Type,
			"name":        types.StringType,
			"postal_code": types.StringType,
			"state":       types.StringType,
			"status":      types.StringType,
			"timezone":    types.StringType,
		},
	}

	// Define the Timeframe object type.
	TimeframeObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"days":                types.SetType{ElemType: customTypes.CustomStringType{}},
			"duration_in_minutes": types.Int32Type,
			"exclude_holidays":    types.BoolType,
			"name":                customTypes.CustomStringType{},
			"start_time":          types.StringType,
		},
	}

	// Define the Device object type.
	DeviceObjectType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":                 types.StringType,
			"target_name":        types.StringType,
			"country":            types.StringType,
			"default_device":     types.BoolType,
			"delay":              types.Int32Type,
			"device_type":        types.StringType,
			"email_address":      types.StringType,
			"external_key":       types.StringType,
			"externally_owned":   types.BoolType,
			"name":               types.StringType,
			"owner":              types.StringType,
			"phone_number":       types.StringType,
			"pin":                types.StringType,
			"priority_threshold": types.StringType,
			"sequence":           types.Int32Type,
			"status":             types.StringType,
			"test_status":        types.StringType,
			"timeframes":         types.SetType{ElemType: TimeframeObjectType},
			"two_way_device":     types.BoolType,
		},
	}
)

// ------------------------------------------------------------
// Object Translaters
// ------------------------------------------------------------

// FlattenServiceLinkObject returns a new `basetypes.ObjectValue` from the provided `xmatters.ServiceLink` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenServiceLinkObject(diags *diag.Diagnostics, in *xmatters.ServiceLink) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(ServiceLinkObjectType.AttrTypes)
	}
	objReturn, d := types.ObjectValue(ServiceLinkObjectType.AttrTypes,
		map[string]attr.Value{
			"link_text": customTypes.StringPointerValue(in.Label),
			"url":       types.StringPointerValue(in.URL),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(ServiceLinkObjectType.AttrTypes)
	}
	return objReturn
}

// ServiceObject returns a new `basetypes.ObjectValue` from the provided `xmatters.Service` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenServiceObject(diags *diag.Diagnostics, in *xmatters.Service) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(ServiceObjectType.AttrTypes)
	}
	var ownerID *string
	if in.OwnedBy != nil {
		ownerID = in.OwnedBy.ID
	}
	objReturn, d := types.ObjectValue(ServiceObjectType.AttrTypes,
		map[string]attr.Value{
			"id":          types.StringPointerValue(in.ID),
			"name":        customTypes.StringPointerValue(in.TargetName),
			"description": customTypes.StringPointerValue(in.Description),
			"type":        customTypes.StringPointerValue(in.ServiceType),
			"tier":        types.StringPointerValue(in.ServiceTier),
			"owner":       types.StringPointerValue(ownerID),
			"links":       FlattenServiceLinkSet(diags, in.ServiceLinks),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(ServiceObjectType.AttrTypes)
	}
	return objReturn
}

// FlattenAuthConfigObject returns a new `basetypes.ObjectValue` from the provided `xmatters.QuotaDetails` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenQuotaDetailsObject(diags *diag.Diagnostics, in *xmatters.QuotaDetails) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(UserQuotaDetailsObjectType.AttrTypes)
	}
	objReturn, d := types.ObjectValue(UserQuotaDetailsObjectType.AttrTypes,
		map[string]attr.Value{
			"total":  types.Int64PointerValue(in.Total),
			"active": types.Int64PointerValue(in.Active),
			"unused": types.Int64PointerValue(in.Unused),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(UserQuotaDetailsObjectType.AttrTypes)
	}
	return objReturn
}

// FlattenGroupObject returns a new `basetypes.ObjectValue` from the provided `xmatters.Group` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenGroupObject(diags *diag.Diagnostics, in *xmatters.Group) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(GroupObjectType.AttrTypes)
	}
	objReturn, d := types.ObjectValue(GroupObjectType.AttrTypes,
		map[string]attr.Value{
			"id":                  types.StringPointerValue(in.ID),
			"name":                types.StringPointerValue(in.TargetName),
			"description":         types.StringPointerValue(in.Description),
			"status":              types.StringPointerValue(in.Status),
			"external_key":        types.StringPointerValue(in.ExternalKey),
			"externally_owned":    types.BoolPointerValue(in.ExternallyOwned),
			"allow_duplicates":    types.BoolPointerValue(in.AllowDuplicates),
			"site":                FlattenReferenceID(in.Site),
			"observed_by_all":     types.BoolPointerValue(in.ObservedByAll),
			"observers":           FlattenReferenceNameSet(diags, in.Observers),
			"supervisors":         FlattenReferenceIDSet(diags, in.Supervisors),
			"group_type":          types.StringPointerValue(in.GroupType),
			"use_default_devices": types.BoolPointerValue(in.UseDefaultDevices),
			"criteria":            FlattenGroupCriteriaObject(diags, in.Criteria),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(GroupObjectType.AttrTypes)
	}
	return objReturn
}

// FlattenGroupMemberObject returns a new `basetypes.ObjectValue` from the provided `xmatters.GroupMember` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenGroupMemberObject(diags *diag.Diagnostics, in *xmatters.GroupMember) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(GroupMemberObjectType.AttrTypes)
	}
	objReturn, d := types.ObjectValue(GroupMemberObjectType.AttrTypes,
		map[string]attr.Value{
			"id":          types.StringPointerValue(in.ID),
			"member_type": types.StringPointerValue(in.MemberType),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(GroupMemberObjectType.AttrTypes)
	}
	return objReturn
}

// FlattenGroupReferenceObject returns a new `basetypes.ObjectValue` from the provided `xmatters.GroupReference` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenGroupReferenceObject(diags *diag.Diagnostics, in *xmatters.GroupReference) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(GroupReferenceObjectType.AttrTypes)
	}
	objReturn, d := types.ObjectValue(GroupReferenceObjectType.AttrTypes,
		map[string]attr.Value{
			"id":             types.StringPointerValue(in.ID),
			"name":           types.StringPointerValue(in.TargetName),
			"recipient_type": types.StringPointerValue(in.RecipientType),
			"group_type":     types.StringPointerValue(in.GroupType),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(GroupReferenceObjectType.AttrTypes)
	}
	return objReturn
}

// FlattenGroupCriteriaObject returns a new `basetypes.ObjectValue` from the provided `xmatters.SearchCriteria` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenGroupCriteriaObject(diags *diag.Diagnostics, in *xmatters.SearchCriteria) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(GroupCriteriaObjectType.AttrTypes)
	}
	objReturn, d := types.ObjectValue(GroupCriteriaObjectType.AttrTypes,
		map[string]attr.Value{
			"operand":   types.StringPointerValue(in.Operand),
			"criterion": FlattenGroupCriterionSet(diags, in.Criterion),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(GroupCriteriaObjectType.AttrTypes)
	}
	return objReturn
}

// FlattenGroupCriterionObject returns a new `basetypes.ObjectValue` from the provided `xmatters.SearchCriterion` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenGroupCriterionObject(diags *diag.Diagnostics, in *xmatters.SearchCriterion) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(GroupCriterionObjectType.AttrTypes)
	}
	objReturn, d := types.ObjectValue(GroupCriterionObjectType.AttrTypes,
		map[string]attr.Value{
			"criterion_type": types.StringPointerValue(in.CriterionType),
			"field":          types.StringPointerValue(in.Field),
			"operand":        types.StringPointerValue(in.Operand),
			"value":          types.StringPointerValue(in.Value),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(GroupCriterionObjectType.AttrTypes)
	}
	return objReturn
}

// FlattenPersonObject returns a new `basetypes.ObjectValue` from the provided `xmatters.Person` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenPersonObject(diags *diag.Diagnostics, in *xmatters.Person) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(PersonObjectType.AttrTypes)
	}
	var siteID *string
	if in.Site != nil {
		siteID = in.Site.ID
	}
	personObject, d := types.ObjectValue(PersonObjectType.AttrTypes,
		map[string]attr.Value{
			"id":               types.StringPointerValue(in.ID),
			"target_name":      customTypes.StringPointerValue(in.TargetName),
			"first_name":       customTypes.StringPointerValue(in.FirstName),
			"last_name":        customTypes.StringPointerValue(in.LastName),
			"roles":            FlattenRoleSet(diags, in.Roles),
			"status":           types.StringPointerValue(in.Status),
			"web_login":        customTypes.StringPointerValue(in.WebLogin),
			"site":             types.StringPointerValue(siteID),
			"timezone":         customTypes.StringPointerValue(in.Timezone),
			"language":         customTypes.StringPointerValue(in.Language),
			"supervisors":      FlattenSupervisorSet(diags, in.Supervisors),
			"phone_login":      types.StringPointerValue(in.PhoneLogin),
			"license_type":     customTypes.StringPointerValue(in.LicenseType),
			"external_key":     customTypes.StringPointerValue(in.ExternalKey),
			"externally_owned": types.BoolPointerValue(in.ExternallyOwned),
			"last_login":       types.StringPointerValue(in.LastLogin),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(PersonObjectType.AttrTypes)
	}
	return personObject
}

// FlattenSiteObject returns a new `basetypes.ObjectValue` from the provided `xmatters.Site` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenSiteObject(diags *diag.Diagnostics, in *xmatters.Site) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(SiteObjectType.AttrTypes)
	}
	objReturn, d := types.ObjectValue(SiteObjectType.AttrTypes,
		map[string]attr.Value{
			"address1":    types.StringPointerValue(in.Address1),
			"address2":    types.StringPointerValue(in.Address2),
			"city":        types.StringPointerValue(in.City),
			"country":     types.StringPointerValue(in.Country),
			"id":          types.StringPointerValue(in.ID),
			"language":    types.StringPointerValue(in.Language),
			"latitude":    types.Float64PointerValue(in.Latitude),
			"longitude":   types.Float64PointerValue(in.Longitude),
			"name":        types.StringPointerValue(in.Name),
			"postal_code": types.StringPointerValue(in.PostalCode),
			"state":       types.StringPointerValue(in.State),
			"status":      types.StringPointerValue(in.Status),
			"timezone":    types.StringPointerValue(in.Timezone),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(SiteObjectType.AttrTypes)
	}
	return objReturn
}

// FlattenTimeframeObject returns a new `basetypes.ObjectValue` from the provided `xmatters.DeviceTimeframe` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenTimeframeObject(diags *diag.Diagnostics, in *xmatters.DeviceTimeframe) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(TimeframeObjectType.AttrTypes)
	}

	daysElems := make([]attr.Value, 0, len(in.Days))
	for _, d := range in.Days {
		daysElems = append(daysElems, customTypes.StringPointerValue(d))
	}
	daysSet, d := types.SetValue(customTypes.CustomStringType{}, daysElems)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(TimeframeObjectType.AttrTypes)
	}

	objectReturn, d := types.ObjectValue(
		TimeframeObjectType.AttrTypes,
		map[string]attr.Value{
			"days":                daysSet,
			"duration_in_minutes": types.Int32PointerValue(in.DurationInMinutes),
			"exclude_holidays":    types.BoolPointerValue(in.ExcludeHolidays),
			"name":                customTypes.StringPointerValue(in.Name),
			"start_time":          types.StringPointerValue(in.StartTime),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(TimeframeObjectType.AttrTypes)
	}
	return objectReturn
}

// FlattenDeviceObject returns a new `basetypes.ObjectValue` from the provided `xmatters.Device` object.
// If the provided object is nil, a null object with appropriate attribute types is returned.
func FlattenDeviceObject(diags *diag.Diagnostics, in *xmatters.Device) basetypes.ObjectValue {
	if in == nil {
		return types.ObjectNull(DeviceObjectType.AttrTypes)
	}
	objReturn, d := types.ObjectValue(DeviceObjectType.AttrTypes,
		map[string]attr.Value{
			"id":                 types.StringPointerValue(in.ID),
			"target_name":        types.StringPointerValue(in.TargetName),
			"country":            types.StringPointerValue(in.Country),
			"default_device":     types.BoolPointerValue(in.DefaultDevice),
			"delay":              types.Int32PointerValue(in.Delay),
			"device_type":        types.StringPointerValue(in.DeviceType),
			"email_address":      types.StringPointerValue(in.EmailAddress),
			"external_key":       types.StringPointerValue(in.ExternalKey),
			"externally_owned":   types.BoolPointerValue(in.ExternallyOwned),
			"name":               types.StringPointerValue(in.Name),
			"owner":              FlattenPersonReferenceID(in.Owner),
			"phone_number":       types.StringPointerValue(in.PhoneNumber),
			"pin":                types.StringPointerValue(in.PIN),
			"priority_threshold": types.StringPointerValue(in.PriorityThreshold),
			"sequence":           types.Int32PointerValue(in.Sequence),
			"status":             types.StringPointerValue(in.Status),
			"test_status":        types.StringPointerValue(in.TestStatus),
			"timeframes":         FlattenTimeframeSet(diags, in.Timeframes),
			"two_way_device":     types.BoolPointerValue(in.TwoWayDevice),
		},
	)
	if diags.Append(d...); d.HasError() {
		return types.ObjectNull(DeviceObjectType.AttrTypes)
	}
	return objReturn
}

// ExpandGroupCriteriaObject transforms a Terraform ObjectValue into an xMatters SearchCriteria object.
// If the provided object is nil, a nil SearchCriteria object is returned.
func ExpandGroupCriteriaObject(diags *diag.Diagnostics, in basetypes.ObjectValue) *xmatters.SearchCriteria {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	// Get the attributes from the object
	attrs := in.Attributes()
	// Extract operand
	operandAttr, ok := attrs["operand"]
	if !ok {
		return nil
	}
	operand := operandAttr.(basetypes.StringValue).ValueStringPointer()
	// Extract criterion set
	criterionAttr, ok := attrs["criterion"]
	if !ok {
		return nil
	}
	criterion := ExpandGroupCriterionSet(diags, criterionAttr.(basetypes.SetValue))
	return &xmatters.SearchCriteria{
		Operand:   operand,
		Criterion: criterion,
	}
}

// ------------------------------------------------------------
// List Flatteners
// ------------------------------------------------------------

// FlattenPersonList transforms a slice of xMatters Person objects into a Terraform List.
// If the provided object is nil, a null list with appropriate element types is returned.
func FlattenPersonList(diags *diag.Diagnostics, people []*xmatters.Person) types.List {
	// Create a slice of Person objects
	personElems := make([]attr.Value, 0, len(people))
	for _, p := range people {
		personObj := FlattenPersonObject(diags, p)
		if diags.HasError() {
			return types.ListNull(PersonObjectType)
		}
		personElems = append(personElems, personObj)
	}
	// Create and return a types.List of person objects
	peopleList, d := types.ListValue(PersonObjectType, personElems)
	if diags.Append(d...); d.HasError() {
		return types.ListNull(PersonObjectType)
	}
	return peopleList
}

// FlattenServiceList transforms a slice of xMatters Service objects into a Terraform List.
// If the provided object is nil, a null list with appropriate element types is returned.
func FlattenServiceList(diags *diag.Diagnostics, services []*xmatters.Service) types.List {
	// Create a slice of Service objects
	serviceElems := make([]attr.Value, 0, len(services))
	for _, s := range services {
		serviceObj := FlattenServiceObject(diags, s)
		if diags.HasError() {
			return types.ListNull(ServiceObjectType)
		}
		serviceElems = append(serviceElems, serviceObj)
	}
	// Create and return a types.List of service objects
	serviceList, d := types.ListValue(ServiceObjectType, serviceElems)
	if diags.Append(d...); d.HasError() {
		return types.ListNull(ServiceObjectType)
	}
	return serviceList
}

// FlattenGroupList transforms a slice of xMatters Group objects into a Terraform List.
// If the provided object is nil, a null list with appropriate element types is returned.
func FlattenGroupList(diags *diag.Diagnostics, groups []*xmatters.Group) types.List {
	// Create a slice of Group objects
	groupElems := make([]attr.Value, 0, len(groups))
	for _, g := range groups {
		groupObj := FlattenGroupObject(diags, g)
		if diags.HasError() {
			return types.ListNull(GroupObjectType)
		}
		groupElems = append(groupElems, groupObj)
	}
	// Create and return a types.List of group objects
	groupList, d := types.ListValue(GroupObjectType, groupElems)
	if diags.Append(d...); d.HasError() {
		return types.ListNull(GroupObjectType)
	}
	return groupList
}

// FlattenSiteList transforms a slice of xMatters Site objects into a Terraform List.
// If the provided object is nil, a null list with appropriate element types is returned.
func FlattenSiteList(diags *diag.Diagnostics, sites []*xmatters.Site) types.List {
	// Create a slice of Site objects
	siteElems := make([]attr.Value, 0, len(sites))
	for _, s := range sites {
		siteObj := FlattenSiteObject(diags, s)
		if diags.HasError() {
			return types.ListNull(SiteObjectType)
		}
		siteElems = append(siteElems, siteObj)
	}
	// Create and return a types.List of Site objects
	siteList, d := types.ListValue(SiteObjectType, siteElems)
	if diags.Append(d...); d.HasError() {
		return types.ListNull(SiteObjectType)
	}
	return siteList
}

// FlattenDeviceList transforms a slice of xMatters Device objects into a Terraform List.
// If the provided object is nil, a null list with appropriate element types is returned.
func FlattenDeviceList(diags *diag.Diagnostics, devices []*xmatters.Device) types.List {
	// Create a slice of Device objects
	deviceElems := make([]attr.Value, 0, len(devices))
	for _, d := range devices {
		deviceObj := FlattenDeviceObject(diags, d)
		if diags.HasError() {
			return types.ListNull(DeviceObjectType)
		}
		deviceElems = append(deviceElems, deviceObj)
	}
	// Create and return a types.List of device objects
	deviceList, d := types.ListValue(DeviceObjectType, deviceElems)
	if diags.Append(d...); d.HasError() {
		return types.ListNull(DeviceObjectType)
	}
	return deviceList
}

// ------------------------------------------------------------
// List Expanders
// ------------------------------------------------------------

// ExpandStringPointerSliceList transforms a Terraform List into a slice strings.
// If the provided list is null or unknown, a nil slice is returned.
func ExpandStringPointerSliceList(diags *diag.Diagnostics, in basetypes.ListValue) []*string {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	var strSlice []*string
	if diags.Append(in.ElementsAs(context.TODO(), &strSlice, false)...); diags.HasError() {
		return nil
	}
	return strSlice
}

// ExpandStringSliceList transforms a Terraform List into a slice strings.
// If the provided list is null or unknown, a nil slice is returned.
func ExpandStringSliceList(diags *diag.Diagnostics, in basetypes.ListValue) []string {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	var strSlice []string
	if diags.Append(in.ElementsAs(context.TODO(), &strSlice, false)...); diags.HasError() {
		return nil
	}
	return strSlice
}

// ExpandStringList transforms a Terraform List into a comma-separated string.
// If the provided list is null or unknown, an empty string is returned.
func ExpandStringList(diags *diag.Diagnostics, in basetypes.ListValue) string {
	if in.IsNull() || in.IsUnknown() {
		return ""
	}
	// Convert to slice of strings
	var strSlice []string
	if diags.Append(in.ElementsAs(context.Background(), &strSlice, false)...); diags.HasError() {
		return ""
	}
	// Join slice into a comma-separated string
	str := strings.Join(strSlice, ",")
	return str
}

// ExpandEncodedStringList transforms a Terraform List into a URL encoded comma-separated string.
// If the provided list is null or unknown, an empty string is returned.
func ExpandEncodedStringList(diags *diag.Diagnostics, in basetypes.ListValue) string {
	if in.IsNull() || in.IsUnknown() {
		return ""
	}
	// Convert to slice of strings
	var strSlice []string
	if diags.Append(in.ElementsAs(context.Background(), &strSlice, false)...); diags.HasError() {
		return ""
	}
	for i, s := range strSlice {
		strSlice[i] = url.QueryEscape(s)
	}
	// Join slice into a comma-separated string
	str := strings.Join(strSlice, ",")
	return str
}

// ------------------------------------------------------------
// Set Flatteners
// ------------------------------------------------------------

// FlattenServiceLinkSet transforms a slice of xMatters ServiceLink objects into a Terraform Set.
// If the provided object is nil, a null set with appropriate element types is returned.
func FlattenServiceLinkSet(diags *diag.Diagnostics, serviceLinks []*xmatters.ServiceLink) types.Set {
	// Create a slice of Service Link objects
	serviceLinkElems := make([]attr.Value, 0, len(serviceLinks))
	for _, sl := range serviceLinks {
		serviceLinkObj := FlattenServiceLinkObject(diags, sl)
		if diags.HasError() {
			return types.SetNull(ServiceLinkObjectType)
		}
		serviceLinkElems = append(serviceLinkElems, serviceLinkObj)
	}
	// Create and return a types.Set of service link objects
	serviceLinkList, d := types.SetValue(ServiceLinkObjectType, serviceLinkElems)
	if diags.Append(d...); diags.HasError() {
		return types.SetNull(ServiceLinkObjectType)
	}
	return serviceLinkList
}

// FlattenSupervisorSet transforms a slice of xMatters Person objects into a Terraform Set.
// If the provided slice of objects is nil, a null set with appropriate element types is returned.
func FlattenSupervisorSet(diags *diag.Diagnostics, people []*xmatters.Person) types.Set {
	// Create a slice of Person ID strings
	supervisorElems := make([]attr.Value, 0, len(people))
	for _, pr := range people {
		supervisorElems = append(supervisorElems, types.StringPointerValue(pr.ID))
	}
	// Create and return a types.List of supervisor ID strings
	supervisorList, d := types.SetValue(types.StringType, supervisorElems)
	if diags.Append(d...); diags.HasError() {
		return types.SetNull(types.StringType)
	}
	return supervisorList
}

// FlattenServiceSet transforms a slice of xMatters Service objects into a Terraform Set.
// If the provided object is nil, a null set with appropriate element types is returned.
func FlattenServiceSet(diags *diag.Diagnostics, services []*xmatters.Service) types.Set {
	// Create a slice of Service objects
	serviceElems := make([]attr.Value, 0, len(services))
	for _, s := range services {
		serviceObj := FlattenServiceObject(diags, s)
		if diags.HasError() {
			return types.SetNull(ServiceObjectType)
		}
		serviceElems = append(serviceElems, serviceObj)
	}
	// Create and return a types.List of service objects
	serviceList, d := types.SetValue(ServiceObjectType, serviceElems)
	if diags.Append(d...); d.HasError() {
		return types.SetNull(ServiceObjectType)
	}
	return serviceList
}

// FlattenRoleSet transforms an xMatters RolePagination object into a Terraform Set.
// If the provided slice of objects is nil, a null list with appropriate element types is returned.
func FlattenRoleSet(diags *diag.Diagnostics, roles []*xmatters.Role) types.Set {
	// Create a slice of Role name strings
	roleElems := make([]attr.Value, 0, len(roles))
	for _, r := range roles {
		roleElems = append(roleElems, customTypes.StringPointerValue(r.Name))
	}
	// Create and return a types.Set of role name strings
	roleSet, d := types.SetValue(customTypes.CustomStringType{}, roleElems)
	if diags.Append(d...); diags.HasError() {
		return types.SetNull(customTypes.CustomStringType{})
	}
	return roleSet
}

// FlattenTimeframeSet transforms a slice of xMatters DeviceTimeframe objects into a Terraform Set.
// If the provided slice of objects is nil, a null list with appropriate element types is returned.
func FlattenTimeframeSet(diags *diag.Diagnostics, timeframes []*xmatters.DeviceTimeframe) types.Set {
	// Create a slice of Timeframe objects
	timeframeElems := make([]attr.Value, 0, len(timeframes))
	for _, t := range timeframes {
		timeframeObj := FlattenTimeframeObject(diags, t)
		if diags.HasError() {
			return types.SetNull(TimeframeObjectType)
		}
		timeframeElems = append(timeframeElems, timeframeObj)
	}
	// Create and return a types.Set of Timeframe objects
	timeframeSet, d := types.SetValue(TimeframeObjectType, timeframeElems)
	if diags.Append(d...); d.HasError() {
		return types.SetNull(TimeframeObjectType)
	}
	return timeframeSet
}

// FlattenReferenceID transforms a slice of xMatters ReferenceById objects into a Terraform Set.
// If the provided object is nil, a null set with appropriate element types is returned.
func FlattenReferenceIDSet(diags *diag.Diagnostics, references []*xmatters.ReferenceById) types.Set {
	// Create a slice of Role name strings
	referenceElems := make([]attr.Value, 0, len(references))
	for _, o := range references {
		referenceElems = append(referenceElems, types.StringPointerValue(o.ID))
	}
	// Create and return a types.List of role name strings
	refIdSet, d := types.SetValue(types.StringType, referenceElems)
	if diags.Append(d...); diags.HasError() {
		return types.SetNull(types.StringType)
	}
	return refIdSet
}

// FlattenReferenceNameSet transforms a slice of xMatters ReferenceByName objects into a Terraform Set.
// If the provided object is nil, a null set with appropriate element types is returned.
func FlattenReferenceNameSet(diags *diag.Diagnostics, references []*xmatters.ReferenceByName) types.Set {
	// Create a slice of Role name strings
	referenceElems := make([]attr.Value, 0, len(references))
	for _, o := range references {
		referenceElems = append(referenceElems, customTypes.StringPointerValue(o.Name))
	}
	// Create and return a types.List of role name strings
	refNameSet, d := types.SetValue(customTypes.CustomStringType{}, referenceElems)
	if diags.Append(d...); diags.HasError() {
		return types.SetNull(customTypes.CustomStringType{})
	}
	return refNameSet
}

// FlattenGroupMemberSet transforms a slice of xMatters GroupMember objects into a Terraform Set.
// If the provided object is nil, a null set with appropriate element types is returned.
func FlattenGroupMemberSet(diags *diag.Diagnostics, in []*xmatters.GroupMember) types.Set {
	groupMemberElems := make([]attr.Value, 0, len(in))
	for _, m := range in {
		groupMemberObj := FlattenGroupMemberObject(diags, m)
		if diags.HasError() {
			return types.SetNull(GroupMemberObjectType)
		}
		groupMemberElems = append(groupMemberElems, groupMemberObj)
	}
	groupMemberSet, d := types.SetValue(GroupMemberObjectType, groupMemberElems)
	if diags.Append(d...); d.HasError() {
		return types.SetNull(GroupMemberObjectType)
	}
	return groupMemberSet
}

// FlattenGroupCriterionSet transforms a slice of xMatters SearchCriterion objects into a Terraform Set.
// If the provided object is nil, a null set with appropriate element types is returned.
func FlattenGroupCriterionSet(diags *diag.Diagnostics, in []*xmatters.SearchCriterion) types.Set {
	elems := make([]attr.Value, 0, len(in))
	for _, c := range in {
		elems = append(elems, FlattenGroupCriterionObject(diags, c))
	}
	set, d := types.SetValue(GroupCriterionObjectType, elems)
	if diags.Append(d...); d.HasError() {
		return types.SetNull(GroupCriterionObjectType)
	}
	return set
}

// ------------------------------------------------------------
// Set Expanders
// ------------------------------------------------------------

// ExpandServiceLinkSet transforms a Terraform Set into a slice of xMatters ServiceLink objects.
// If the provided set is null or unknown, a nil slice is returned.
func ExpandServiceLinkSet(diags *diag.Diagnostics, in basetypes.SetValue) []*xmatters.ServiceLink {
	if in.IsNull() || in.IsUnknown() {
		return []*xmatters.ServiceLink{}
	}
	var links []*xmatters.ServiceLink
	if diags.Append(in.ElementsAs(context.TODO(), &links, false)...); diags.HasError() {
		return nil
	}
	return links
}

// ExpandStringSliceSet transforms a Terraform List into a slice of xMatters Supervisor ID strings.
// If the provided set is null or unknown, a nil slice is returned.
func ExpandStringSliceSet(diags *diag.Diagnostics, in basetypes.SetValue) []*string {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	var stringSlice []*string
	if diags.Append(in.ElementsAs(context.TODO(), &stringSlice, false)...); diags.HasError() {
		return nil
	}
	return stringSlice
}

// ExpandTimeframeSet transforms a Terraform Set into a slice of xMatters DeviceTimeframe objects.
// If the provided set is null or unknown, a nil slice is returned.
func ExpandTimeframeSet(diags *diag.Diagnostics, in basetypes.SetValue) []*xmatters.DeviceTimeframe {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	var timeframes []*xmatters.DeviceTimeframe
	if diags.Append(in.ElementsAs(context.TODO(), &timeframes, false)...); diags.HasError() {
		return nil
	}
	return timeframes
}

// ExpandReferenceIDSet transforms a Terraform Set into a slice of xMatters ReferenceById objects.
// If the provided set is null or unknown, a nil slice is returned.
func ExpandReferenceIDSet(diags *diag.Diagnostics, in basetypes.SetValue) []*xmatters.ReferenceById {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	stringSlice := ExpandStringSliceSet(diags, in)
	if diags.HasError() {
		return nil
	}
	references := make([]*xmatters.ReferenceById, 0, len(stringSlice))
	for _, s := range stringSlice {
		references = append(references, &xmatters.ReferenceById{ID: s})
	}
	return references
}

// ExpandReferenceNameSet transforms a Terraform Set into a slice of xMatters ReferenceByName objects.
// If the provided set is null or unknown, a nil slice is returned.
func ExpandReferenceNameSet(diags *diag.Diagnostics, in basetypes.SetValue) []*xmatters.ReferenceByName {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	stringSlice := ExpandStringSliceSet(diags, in)
	if diags.HasError() {
		return nil
	}
	references := make([]*xmatters.ReferenceByName, 0, len(stringSlice))
	for _, s := range stringSlice {
		references = append(references, &xmatters.ReferenceByName{Name: s})
	}
	return references
}

// ExpandGroupMemberSet transforms a Terraform Set into a slice of xMatters GroupMember objects.
// If the provided set is null or unknown, a nil slice is returned.
func ExpandGroupMemberSet(diags *diag.Diagnostics, in basetypes.SetValue) []*xmatters.GroupMember {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	var members []*xmatters.GroupMember
	if diags.Append(in.ElementsAs(context.TODO(), &members, false)...); diags.HasError() {
		return nil
	}
	return members
}

// ExpandGroupCriterionSet transforms a Terraform Set into a slice of xMatters SearchCriterion objects.
// If the provided set is null or unknown, a nil slice is returned.
func ExpandGroupCriterionSet(diags *diag.Diagnostics, in basetypes.SetValue) []*xmatters.SearchCriterion {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	var criteria []*xmatters.SearchCriterion
	if diags.Append(in.ElementsAs(context.TODO(), &criteria, false)...); diags.HasError() {
		return nil
	}
	return criteria
}

// ------------------------------------------------------------
// Primitive Flatteners
// ------------------------------------------------------------

// FlattenGroupReferenceID transforms an xMatters GroupReference object into a Terraform StringValue.
// If the provided object is nil, a null string is returned.
func FlattenGroupReferenceID(in *xmatters.GroupReference) basetypes.StringValue {
	if in == nil {
		return basetypes.NewStringNull()
	}
	return types.StringPointerValue(in.ID)
}

// FlattenServiceReferenceID transforms an xMatters ServiceReference object into a Terraform StringValue.
// If the provided object is nil, a null string is returned.
func FlattenServiceReferenceID(in *xmatters.ServiceReference) basetypes.StringValue {
	if in == nil {
		return basetypes.NewStringNull()
	}
	return types.StringPointerValue(in.ID)
}

// FlattenPersonReferenceID transforms an xMatters PersonReference object into a Terraform StringValue.
// If the provided object is nil, a null string is returned.
func FlattenPersonReferenceID(in *xmatters.PersonReference) basetypes.StringValue {
	if in == nil {
		return basetypes.NewStringNull()
	}
	return types.StringPointerValue(in.ID)
}

// FlattenReferenceID transforms an xMatters ReferenceById object into a Terraform StringValue.
// If the provided object is nil, a null string is returned.
func FlattenReferenceID(in *xmatters.ReferenceById) basetypes.StringValue {
	if in == nil {
		return basetypes.NewStringNull()
	}
	return types.StringPointerValue(in.ID)
}

// ------------------------------------------------------------
// Primitive Expanders
// ------------------------------------------------------------

// ExpandGroupReferenceId transforms a Terraform StringValue into an xMatters GroupReference object.
// If the provided string is null or unknown, a nil object is returned.
func ExpandGroupReferenceId(in basetypes.StringValue) *xmatters.GroupReference {
	if in.IsNull() || in.IsUnknown() {
		return nil
	}
	return &xmatters.GroupReference{ID: in.ValueStringPointer()}
}

// ------------------------------------------------------------
