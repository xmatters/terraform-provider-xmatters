package utils

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
)

func TestPreserveExplicitEmptyString(t *testing.T) {
	config := types.StringValue("")
	state := types.StringValue("ON_CALL")

	PreserveExplicitEmptyString(config, &state)

	if state.ValueString() != "" {
		t.Fatalf("expected state string to remain explicit empty value, got %q", state.ValueString())
	}
}

func TestPreserveExplicitEmptySet(t *testing.T) {
	config := types.SetValueMust(types.StringType, []attr.Value{})
	state := types.SetValueMust(types.StringType, []attr.Value{
		types.StringValue("4b6fe547-24da-c5eb-cdea-1fa664f8053e"),
	})

	PreserveExplicitEmptySet(config, &state)

	if len(state.Elements()) != 0 {
		t.Fatalf("expected state set to remain explicit empty value, got %d elements", len(state.Elements()))
	}
}

func TestPreserveExplicitEmptyHelpers_DoNotOverrideNonEmpty(t *testing.T) {
	stringConfig := types.StringValue("DYNAMIC")
	stringState := types.StringValue("ON_CALL")
	PreserveExplicitEmptyString(stringConfig, &stringState)
	if stringState.ValueString() != "ON_CALL" {
		t.Fatalf("expected non-empty string state to remain unchanged, got %q", stringState.ValueString())
	}

	setConfig := types.SetValueMust(types.StringType, []attr.Value{
		types.StringValue("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
	})
	setState := types.SetValueMust(types.StringType, []attr.Value{
		types.StringValue("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
	})
	PreserveExplicitEmptySet(setConfig, &setState)
	if len(setState.Elements()) != 1 {
		t.Fatalf("expected non-empty set state to remain unchanged, got %d elements", len(setState.Elements()))
	}
}

func TestPreserveExplicitEmptyString_WithCustomStringValue(t *testing.T) {
	config := customTypes.StringValue("")
	state := customTypes.StringPointerValue(nil)

	PreserveExplicitEmptyString(config.StringValue, &state.StringValue)

	if state.IsNull() {
		t.Fatalf("expected state custom string to be explicit empty string, got null")
	}
	if state.ValueString() != "" {
		t.Fatalf("expected state custom string to remain explicit empty value, got %q", state.ValueString())
	}
}

func TestPreservePriorStringWhenAPIOmitted(t *testing.T) {
	prior := types.StringValue("1234")
	state := types.StringNull()

	PreservePriorStringWhenAPIOmitted(true, prior, &state)

	if state.IsNull() || state.ValueString() != "1234" {
		t.Fatalf("expected state string to preserve prior value when API omits field, got %q", state.ValueString())
	}
}

func TestPreservePriorSetWhenAPIOmitted(t *testing.T) {
	prior := types.SetValueMust(customTypes.CustomStringType{}, []attr.Value{
		customTypes.StringValue("Company Supervisor"),
	})
	state := types.SetValueMust(customTypes.CustomStringType{}, []attr.Value{})

	PreservePriorSetWhenAPIOmitted(true, prior, &state)

	if len(state.Elements()) != 1 {
		t.Fatalf("expected state set to preserve prior value when API omits field, got %d elements", len(state.Elements()))
	}
}

func TestPreservePriorHelpers_DoNotOverrideWhenNotOmitted(t *testing.T) {
	stringPrior := types.StringValue("1234")
	stringState := types.StringValue("5678")
	PreservePriorStringWhenAPIOmitted(false, stringPrior, &stringState)
	if stringState.ValueString() != "5678" {
		t.Fatalf("expected string state to remain unchanged when API did not omit field, got %q", stringState.ValueString())
	}

	setPrior := types.SetValueMust(customTypes.CustomStringType{}, []attr.Value{
		customTypes.StringValue("Company Supervisor"),
	})
	setState := types.SetValueMust(customTypes.CustomStringType{}, []attr.Value{
		customTypes.StringValue("Group Supervisor"),
	})
	PreservePriorSetWhenAPIOmitted(false, setPrior, &setState)
	if len(setState.Elements()) != 1 {
		t.Fatalf("expected set state to remain unchanged when API did not omit field, got %d elements", len(setState.Elements()))
	}
}
