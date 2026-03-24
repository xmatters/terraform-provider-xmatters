package utils

import "github.com/hashicorp/terraform-plugin-framework/types/basetypes"

// PreserveExplicitEmptyString keeps an explicitly-configured empty string stable in state.
// This is useful when an API defaults omitted/empty input to a non-empty value.
func PreserveExplicitEmptyString(config basetypes.StringValue, state *basetypes.StringValue) {
	if !config.IsNull() && !config.IsUnknown() && config.ValueString() == "" {
		*state = config
	}
}

// PreserveExplicitEmptySet keeps an explicitly-configured empty set stable in state.
// This is useful when an API defaults omitted/empty input to one or more values.
func PreserveExplicitEmptySet(config basetypes.SetValue, state *basetypes.SetValue) {
	if !config.IsNull() && !config.IsUnknown() && len(config.Elements()) == 0 {
		*state = config
	}
}
