package customTypes

import (
	"context"
	"fmt"
	"strings"

	"github.com/biter777/countries"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// -----------------------------------------------------------------------------
// Custom Country Type
// -----------------------------------------------------------------------------

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringTypable = CustomCountryType{}

type CustomCountryType struct {
	basetypes.StringType
	// ... potentially other fields ...
}

func (t CustomCountryType) Equal(o attr.Type) bool {
	other, ok := o.(CustomCountryType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t CustomCountryType) String() string {
	return "CustomCountryType"
}

func (t CustomCountryType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	// CustomCountryValue defined in the value type section
	value := CustomCountryValue{
		StringValue: in,
	}

	return value, nil
}

func (t CustomCountryType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (t CustomCountryType) ValueType(ctx context.Context) attr.Value {
	// CustomCountryValue defined in the value type section
	return CustomCountryValue{}
}

// -----------------------------------------------------------------------------
// Custom Country Value
// -----------------------------------------------------------------------------

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringValuable = CustomCountryValue{}

type CustomCountryValue struct {
	basetypes.StringValue
	// ... potentially other fields ...
}

func (v CustomCountryValue) Equal(o attr.Value) bool {
	other, ok := o.(CustomCountryValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v CustomCountryValue) Type(ctx context.Context) attr.Type {
	// CustomCountryType defined in the schema type section
	return CustomCountryType{}
}

func CountryValue(value string) CustomCountryValue {
	return CustomCountryValue{
		StringValue: basetypes.NewStringValue(value),
	}
}

func CountryPointerValue(value *string) CustomCountryValue {
	return CustomCountryValue{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}

// -----------------------------------------------------------------------------
// Custom Country Semantics
// -----------------------------------------------------------------------------

// CustomCountryValue defined in the value type section
// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringValuableWithSemanticEquals = CustomCountryValue{}

func (v CustomCountryValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	// The framework should always pass the correct value type, but always check
	newValue, ok := newValuable.(CustomCountryValue)

	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	returnCode := strings.TrimSpace(v.ValueString())
	inputCountry := strings.TrimSpace(newValue.ValueString())

	// If the strings are equivalent, keep the prior value
	if inputCountry == returnCode || CodeMatch(returnCode, inputCountry) {
		return true, diags
	}
	return false, diags
}

func CodeMatch(countryCode string, input string) bool {
	country := countries.ByName(countryCode)
	if strings.Contains(country.String(), input) || country.Alpha2() == input {
		return true
	}

	return false
}
