package utils_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	providerutils "github.com/xmatters/terraform-provider-xmatters/internal/utils"
)

func testGroupTypeConfig(t *testing.T, groupType string) tfsdk.Config {
	t.Helper()

	testSchema := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"group_type": schema.StringAttribute{
				Optional: true,
			},
		},
	}

	raw := tftypes.NewValue(
		testSchema.Type().TerraformType(context.Background()),
		map[string]tftypes.Value{
			"group_type": tftypes.NewValue(tftypes.String, groupType),
		},
	)

	return tfsdk.Config{
		Raw:    raw,
		Schema: testSchema,
	}
}

func knownCriteriaObject() types.Object {
	return types.ObjectValueMust(
		providerutils.GroupCriteriaObjectType.AttrTypes,
		map[string]attr.Value{
			"operand": types.StringValue("AND"),
			"criterion": types.SetValueMust(
				providerutils.GroupCriterionObjectType,
				[]attr.Value{},
			),
		},
	)
}

func TestGroupTypeValidator_AllowsCriteriaWhenGroupTypeDynamic(t *testing.T) {
	v := providerutils.GroupTypeValidator{RequiredValue: "DYNAMIC"}
	req := validator.ObjectRequest{
		Path:        path.Root("criteria"),
		Config:      testGroupTypeConfig(t, "DYNAMIC"),
		ConfigValue: knownCriteriaObject(),
	}
	resp := &validator.ObjectResponse{}

	v.ValidateObject(context.Background(), req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("expected no validation errors, got: %#v", resp.Diagnostics)
	}
}

func TestGroupTypeValidator_RejectsCriteriaWhenGroupTypeNotDynamic(t *testing.T) {
	v := providerutils.GroupTypeValidator{RequiredValue: "DYNAMIC"}
	req := validator.ObjectRequest{
		Path:        path.Root("criteria"),
		Config:      testGroupTypeConfig(t, "ON_CALL"),
		ConfigValue: knownCriteriaObject(),
	}
	resp := &validator.ObjectResponse{}

	v.ValidateObject(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected validation error for criteria with non-dynamic group_type")
	}
}

func TestGroupTypeValidator_RequiresCriteriaWhenGroupTypeDynamic(t *testing.T) {
	v := providerutils.GroupTypeValidator{RequiredValue: "DYNAMIC"}
	req := validator.ObjectRequest{
		Path:        path.Root("criteria"),
		Config:      testGroupTypeConfig(t, "DYNAMIC"),
		ConfigValue: types.ObjectNull(providerutils.GroupCriteriaObjectType.AttrTypes),
	}
	resp := &validator.ObjectResponse{}

	v.ValidateObject(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected validation error when group_type is dynamic and criteria is missing")
	}
}

func TestGroupTypeValidator_AllowsMissingCriteriaWhenGroupTypeNotDynamic(t *testing.T) {
	v := providerutils.GroupTypeValidator{RequiredValue: "DYNAMIC"}
	req := validator.ObjectRequest{
		Path:        path.Root("criteria"),
		Config:      testGroupTypeConfig(t, "BROADCAST"),
		ConfigValue: types.ObjectNull(providerutils.GroupCriteriaObjectType.AttrTypes),
	}
	resp := &validator.ObjectResponse{}

	v.ValidateObject(context.Background(), req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("expected no validation errors when criteria is missing for non-dynamic group_type, got: %#v", resp.Diagnostics)
	}
}
