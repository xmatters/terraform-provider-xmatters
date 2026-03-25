package utils_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	providerutils "github.com/xmatters/terraform-provider-xmatters/internal/utils"
)

func TestUUIDValidator_AcceptsValidUUID(t *testing.T) {
	v := providerutils.UUIDValidator{}
	req := validator.StringRequest{
		Path:        path.Root("supervisors").AtListIndex(0),
		ConfigValue: types.StringValue("2e75facc-b780-4f6a-80ec-2298561013bb"),
	}
	resp := &validator.StringResponse{}

	v.ValidateString(context.Background(), req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("expected no validation errors for valid UUID, got: %#v", resp.Diagnostics)
	}
}

func TestUUIDValidator_RejectsUUIDWithLeadingOrTrailingWhitespace(t *testing.T) {
	tests := []string{
		" 2e75facc-b780-4f6a-80ec-2298561013bb",
		"2e75facc-b780-4f6a-80ec-2298561013bb ",
		" 2e75facc-b780-4f6a-80ec-2298561013bb ",
	}

	for _, input := range tests {
		req := validator.StringRequest{
			Path:        path.Root("supervisors").AtListIndex(0),
			ConfigValue: types.StringValue(input),
		}
		resp := &validator.StringResponse{}

		providerutils.UUIDValidator{}.ValidateString(context.Background(), req, resp)

		if !resp.Diagnostics.HasError() {
			t.Fatalf("expected validation error for UUID with whitespace: %q", input)
		}
	}
}
