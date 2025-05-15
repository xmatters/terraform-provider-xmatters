package customTypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// -----------------------------------------------------------------------------
// Custom Sequence Type
// -----------------------------------------------------------------------------

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.Int32Typable = CustomSequenceType{}

type CustomSequenceType struct {
	basetypes.Int32Type
	// ... potentially other fields ...
}

func (t CustomSequenceType) Equal(o attr.Type) bool {
	other, ok := o.(CustomSequenceType)

	if !ok {
		return false
	}

	return t.Int32Type.Equal(other.Int32Type)
}

func (t CustomSequenceType) Int32() string {
	return "CustomSequenceType"
}

func (t CustomSequenceType) ValueFromInt32(ctx context.Context, in basetypes.Int32Value) (basetypes.Int32Valuable, diag.Diagnostics) {
	// CustomSequenceValue defined in the value type section
	value := CustomSequenceValue{
		Int32Value: in,
	}

	return value, nil
}

func (t CustomSequenceType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.Int32Type.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	int32Value, ok := attrValue.(basetypes.Int32Value)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	int32Valuable, diags := t.ValueFromInt32(ctx, int32Value)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting Int32Value to Int32Valuable: %v", diags)
	}

	return int32Valuable, nil
}

func (t CustomSequenceType) ValueType(ctx context.Context) attr.Value {
	// CustomSequenceValue defined in the value type section
	return CustomSequenceValue{}
}

// -----------------------------------------------------------------------------
// Custom Sequence Value
// -----------------------------------------------------------------------------

// Ensure the implementation satisfies the expected interfaces
var _ basetypes.Int32Valuable = CustomSequenceValue{}

type CustomSequenceValue struct {
	basetypes.Int32Value
	// ... potentially other fields ...
}

func (v CustomSequenceValue) Equal(o attr.Value) bool {
	other, ok := o.(CustomSequenceValue)

	if !ok {
		return false
	}

	return v.Int32Value.Equal(other.Int32Value)
}

func (v CustomSequenceValue) Type(ctx context.Context) attr.Type {
	// CustomSequenceType defined in the schema type section
	return CustomSequenceType{}
}

func DeviceSequenceValue(value int32) CustomSequenceValue {
	return CustomSequenceValue{
		Int32Value: basetypes.NewInt32Value(value),
	}
}

func DeviceSequencePointerValue(value *int32) CustomSequenceValue {
	return CustomSequenceValue{
		Int32Value: basetypes.NewInt32PointerValue(value),
	}
}

// -----------------------------------------------------------------------------
// Custom Sequence Semantics
// -----------------------------------------------------------------------------

// CustomSequenceValue defined in the value type section
// Ensure the implementation satisfies the expected interfaces
var _ basetypes.Int32ValuableWithSemanticEquals = CustomSequenceValue{}

func (v CustomSequenceValue) Int32SemanticEquals(ctx context.Context, newValuable basetypes.Int32Valuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	// The framework should always pass the correct value type, but always check
	configuredValue, ok := newValuable.(CustomSequenceValue)

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

	// If configured value is greater than or equal to the new value, return true
	if v.ValueInt32() <= configuredValue.ValueInt32() {
		return true, diags
	}
	return false, diags
}
