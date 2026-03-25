package utils

import (
	"context"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/language"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// UUID Input Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

// UUIDValidator validates that the string is a valid UUID.
type UUIDValidator struct{}

func (v UUIDValidator) Description(ctx context.Context) string {
	return "The value must be a valid UUID."
}

func (v UUIDValidator) MarkdownDescription(ctx context.Context) string {
	return "The value must be a valid UUID."
}

func (v UUIDValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	value := req.ConfigValue.ValueString()
	if strings.TrimSpace(value) != value {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid UUID",
			fmt.Sprintf("The value provided is not a valid UUID: %s. Leading or trailing whitespace is not allowed.", value),
		)
		return
	}
	if _, err := uuid.Parse(value); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid UUID",
			fmt.Sprintf("The value provided is not a valid UUID: %s. Err: %s", value, err.Error()),
		)
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Token Input Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

type TokenValidator struct{}

func (v TokenValidator) Description(ctx context.Context) string {
	return "Validates that 'token' is set if 'type' is 'API_TOKEN'."
}

func (v TokenValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v TokenValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	var typeValue types.String
	req.Config.GetAttribute(ctx, path.Root("auth").AtName("auth_type"), &typeValue)

	if typeValue.ValueString() == "API_TOKEN" {
		if req.ConfigValue.IsNull() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing 'token'",
				"Missing 'token' attribute when 'type' is 'API_TOKEN'.",
			)
		}
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Basic Auth Input Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

type BasicAuthValidator struct{}

func (v BasicAuthValidator) Description(ctx context.Context) string {
	return "Validates that 'username' and 'password' are set if 'type' is 'BASIC'."
}

func (v BasicAuthValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v BasicAuthValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	var typeValue types.String
	req.Config.GetAttribute(ctx, path.Root("auth").AtName("auth_type"), &typeValue)

	if typeValue.ValueString() == "BASIC" {
		if req.ConfigValue.IsNull() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				fmt.Sprintf("Missing '%s'", req.Path),
				fmt.Sprintf("Missing '%s' attribute when 'type' is 'BASIC'.", req.Path),
			)
		}
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Phone Login/PIN Input Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

type PhoneValidator struct{}

func (v PhoneValidator) Description(ctx context.Context) string {
	return "Validates that the phone number is in the correct format."
}

func (v PhoneValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v PhoneValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	if len(fmt.Sprint(req.ConfigValue.ValueInt64())) > 30 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			fmt.Sprintf("Invalid '%s'", req.Path),
			fmt.Sprintf("'%s' can have a maximum of 30 characters.", req.Path),
		)
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Phone Number Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

type PhoneNumberValidator struct{}

func (v PhoneNumberValidator) Description(ctx context.Context) string {
	return "Validates that the phone number is in the correct format."
}

func (v PhoneNumberValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v PhoneNumberValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	re := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	if !re.MatchString(req.ConfigValue.ValueString()) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			fmt.Sprintf("Invalid '%s'", req.Path),
			fmt.Sprintf("'%s' must be a valid phone number.", req.Path),
		)
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Timezone Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

type TimezoneValidator struct{}

func (v TimezoneValidator) Description(ctx context.Context) string {
	return "Validates that the timezone is in the correct format."
}

func (v TimezoneValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v TimezoneValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	_, err := time.LoadLocation(req.ConfigValue.ValueString())
	if err != nil || req.ConfigValue.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			fmt.Sprintf("Invalid '%s'", req.Path),
			fmt.Sprintf("'%s' must be a valid timezone.", req.Path),
		)
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Timestamp Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

type TimestampValidator struct{}

func (v TimestampValidator) Description(ctx context.Context) string {
	return "Validates that the timestamp is in the correct format."
}

func (v TimestampValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v TimestampValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	if _, err := time.Parse(time.RFC3339, req.ConfigValue.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			fmt.Sprintf("Invalid '%s'", req.Path),
			fmt.Sprintf("'%s' must be a timestamp in ISO 8601 format.", req.Path),
		)
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Language Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

type LanguageValidator struct{}

func (v LanguageValidator) Description(ctx context.Context) string {
	return "Validates that the language is in the correct format."
}

func (v LanguageValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v LanguageValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	if _, err := language.Parse(req.ConfigValue.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			fmt.Sprintf("Invalid '%s'", req.Path),
			fmt.Sprintf("'%s' must be a valid language code.", req.Path),
		)
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Email Address Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

type EmailValidator struct{}

func (v EmailValidator) Description(ctx context.Context) string {
	return "Validates that the email address is in the correct format."
}

func (v EmailValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v EmailValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	if _, err := mail.ParseAddress(req.ConfigValue.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			fmt.Sprintf("Invalid '%s'", req.Path),
			fmt.Sprintf("'%s' must be a valid email address.", req.Path),
		)
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Latitude Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

type LatitudeValidator struct{}

func (v LatitudeValidator) Description(ctx context.Context) string {
	return "Validates that the latitude is in the correct format."
}

func (v LatitudeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v LatitudeValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	if req.ConfigValue.ValueFloat64() < -90.0 || req.ConfigValue.ValueFloat64() > 90.0 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			fmt.Sprintf("Invalid '%s'", req.Path),
			fmt.Sprintf("'%s' must be a valid latitude between -90.0 and 90.0", req.Path),
		)
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Longitude Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

type LongitudeValidator struct{}

func (v LongitudeValidator) Description(ctx context.Context) string {
	return "Validates that the longitude is in the correct format."
}

func (v LongitudeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v LongitudeValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	if req.ConfigValue.ValueFloat64() < -180.0 || req.ConfigValue.ValueFloat64() > 180.0 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			fmt.Sprintf("Invalid '%s'", req.Path),
			fmt.Sprintf("'%s' must be a valid longitude between -180.0 and 180.0", req.Path),
		)
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Custom Properties Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------

type CustomPropertiesValidator struct{}

func (v CustomPropertiesValidator) Description(ctx context.Context) string {
	return "Validates that only one of the six available object types have been configured."
}

func (v CustomPropertiesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v CustomPropertiesValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	setAttributes := 0
	var mapKey string
	for key, value := range req.ConfigValue.Attributes() {
		mapKey = key
		if !value.IsNull() && !value.IsUnknown() {
			setAttributes++
		}
	}
	if setAttributes != 1 {
		resp.Diagnostics.AddAttributeError(
			req.Path.AtMapKey(mapKey),
			"Invalid Custom Property",
			fmt.Sprintf("Each custom property must have exactly one attribute set, but %d attributes were set.", setAttributes),
		)
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Status Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------
type StatusValidator struct{}

func (v StatusValidator) Description(ctx context.Context) string {
	return "Validates that the status is a valid xMatters status."
}

func (v StatusValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v StatusValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	validStatuses := []string{"ACTIVE", "INACTIVE"}
	for _, status := range validStatuses {
		if req.ConfigValue.ValueString() == status {
			return
		}
	}
	resp.Diagnostics.AddAttributeError(
		req.Path,
		fmt.Sprintf("Invalid '%s'", req.Path),
		fmt.Sprintf("'%s' must be one of the following values: %v.", req.Path, validStatuses),
	)
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Operand Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------
type OperandValidator struct{}

func (v OperandValidator) Description(ctx context.Context) string {
	return "Validates that the operand is a valid xMatters operand."
}

func (v OperandValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v OperandValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	validOperands := []string{"AND", "OR"}
	for _, operand := range validOperands {
		if req.ConfigValue.ValueString() == operand {
			return
		}
	}
	resp.Diagnostics.AddAttributeError(
		req.Path,
		fmt.Sprintf("Invalid '%s'", req.Path),
		fmt.Sprintf("'%s' must be one of the following values: %v.", req.Path, validOperands),
	)
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Sort Order Validation
// -----------------------------------------------------------------------------------------------------------------------------------------------------------
type SortOrderValidator struct{}

func (v SortOrderValidator) Description(ctx context.Context) string {
	return "Validates that the sort order is a valid xMatters sort order."
}

func (v SortOrderValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v SortOrderValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	validSortOrders := []string{"ASCENDING", "DESCENDING"}
	// Check if the value is one of the valid sort orders
	for _, sortOrder := range validSortOrders {
		if req.ConfigValue.ValueString() == sortOrder {
			return
		}
	}
	resp.Diagnostics.AddAttributeError(
		req.Path,
		fmt.Sprintf("Invalid '%s'", req.Path),
		fmt.Sprintf("'%s' must be one of the following values: %v.", req.Path, validSortOrders),
	)
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
// Group Type Validator
// -----------------------------------------------------------------------------------------------------------------------------------------------------------
type GroupTypeValidator struct {
	RequiredValue string
}

func (v GroupTypeValidator) Description(ctx context.Context) string {
	return fmt.Sprintf(
		"The criteria block can only be set when 'group_type' is '%s', and is required when 'group_type' is '%s'.",
		v.RequiredValue,
		v.RequiredValue,
	)
}

func (v GroupTypeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v GroupTypeValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	// Get the group_type attribute from the config
	var groupTypeValue types.String
	req.Config.GetAttribute(ctx, path.Root("group_type"), &groupTypeValue)

	// Skip validation until group_type is known.
	if groupTypeValue.IsUnknown() {
		return
	}

	// If criteria is set, group_type must match RequiredValue.
	if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() && groupTypeValue.ValueString() != v.RequiredValue {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid criteria",
			fmt.Sprintf("The criteria block can only be set when 'group_type' is '%s'.", v.RequiredValue),
		)
		return
	}

	// If group_type is RequiredValue, criteria must be set.
	if req.ConfigValue.IsNull() && groupTypeValue.ValueString() == v.RequiredValue {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Missing criteria",
			fmt.Sprintf("The criteria block is required when 'group_type' is '%s'.", v.RequiredValue),
		)
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------------------
