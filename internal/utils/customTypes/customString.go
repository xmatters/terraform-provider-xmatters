package customTypes

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// -----------------------------------------------------------------------------
// Custom String Type
// -----------------------------------------------------------------------------

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringTypable = CustomStringType{}

type CustomStringType struct {
	basetypes.StringType
	// ... potentially other fields ...
}

func (t CustomStringType) Equal(o attr.Type) bool {
	other, ok := o.(CustomStringType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t CustomStringType) String() string {
	return "CustomStringType"
}

func (t CustomStringType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	// CustomStringValue defined in the value type section
	value := CustomStringValue{
		StringValue: in,
	}

	return value, nil
}

func (t CustomStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

func (t CustomStringType) ValueType(ctx context.Context) attr.Value {
	// CustomStringValue defined in the value type section
	return CustomStringValue{}
}

// -----------------------------------------------------------------------------
// Custom String Value
// -----------------------------------------------------------------------------

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringValuable = CustomStringValue{}

type CustomStringValue struct {
	basetypes.StringValue
	// ... potentially other fields ...
}

func (v CustomStringValue) Equal(o attr.Value) bool {
	other, ok := o.(CustomStringValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v CustomStringValue) Type(ctx context.Context) attr.Type {
	// CustomStringType defined in the schema type section
	return CustomStringType{}
}

func StringValue(value string) CustomStringValue {
	return CustomStringValue{
		StringValue: basetypes.NewStringValue(value),
	}
}

func StringPointerValue(value *string) CustomStringValue {
	return CustomStringValue{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}

// -----------------------------------------------------------------------------
// Custom String Semantics
// -----------------------------------------------------------------------------

// CustomStringValue defined in the value type section
// Ensure the implementation satisfies the expected interfaces
var _ basetypes.StringValuableWithSemanticEquals = CustomStringValue{}

func (v CustomStringValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	// The framework should always pass the correct value type, but always check
	newValue, ok := newValuable.(CustomStringValue)

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

	priorString := strings.TrimSpace(v.ValueString())

	newString := strings.TrimSpace(newValue.ValueString())

	// If the strings are equivalent, keep the prior value
	if priorString == newString {
		return true, diags
	}
	return false, diags
}
