package utils

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
