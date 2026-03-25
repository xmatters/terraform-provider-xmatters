package test

import (
	"context"
	"testing"

	tfresource "github.com/hashicorp/terraform-plugin-framework/resource"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	groupresource "github.com/xmatters/terraform-provider-xmatters/internal/resources/group"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
)

func TestGroupResourceSchemaCriteriaIsConfigurable(t *testing.T) {
	r := &groupresource.GroupResource{}
	resp := &tfresource.SchemaResponse{}

	r.Schema(context.Background(), tfresource.SchemaRequest{}, resp)

	rawCriteriaAttr, ok := resp.Schema.Attributes["criteria"]
	if !ok {
		t.Fatalf("expected 'criteria' attribute in group resource schema")
	}

	criteriaAttr, ok := rawCriteriaAttr.(resourceschema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("expected 'criteria' to be schema.SingleNestedAttribute, got %T", rawCriteriaAttr)
	}

	if !criteriaAttr.Optional {
		t.Fatalf("expected 'criteria' to be optional")
	}
}

func TestGroupCriteriaSchemaNestedAttributesAreConfigurable(t *testing.T) {
	criteriaAttrs := groupresource.GroupCriteriaSchema()

	rawOperandAttr, ok := criteriaAttrs["operand"]
	if !ok {
		t.Fatalf("expected 'operand' attribute in criteria schema")
	}
	operandAttr, ok := rawOperandAttr.(resourceschema.StringAttribute)
	if !ok {
		t.Fatalf("expected 'operand' to be schema.StringAttribute, got %T", rawOperandAttr)
	}
	if !operandAttr.Optional {
		t.Fatalf("expected 'criteria.operand' to be optional")
	}
	if len(operandAttr.Validators) == 0 {
		t.Fatalf("expected 'criteria.operand' to have validators")
	}

	rawCriterionAttr, ok := criteriaAttrs["criterion"]
	if !ok {
		t.Fatalf("expected 'criterion' attribute in criteria schema")
	}
	criterionAttr, ok := rawCriterionAttr.(resourceschema.SetNestedAttribute)
	if !ok {
		t.Fatalf("expected 'criterion' to be schema.SetNestedAttribute, got %T", rawCriterionAttr)
	}
	if !criterionAttr.Optional {
		t.Fatalf("expected 'criteria.criterion' to be optional")
	}

	criterionObjectAttrs := groupresource.GroupCriterionSchema()
	for _, field := range []string{"criterion_type", "field", "operand", "value"} {
		rawFieldAttr, ok := criterionObjectAttrs[field]
		if !ok {
			t.Fatalf("expected '%s' attribute in criteria.criterion schema", field)
		}
		fieldAttr, ok := rawFieldAttr.(resourceschema.StringAttribute)
		if !ok {
			t.Fatalf("expected '%s' to be schema.StringAttribute, got %T", field, rawFieldAttr)
		}
		if !fieldAttr.Optional {
			t.Fatalf("expected 'criteria.criterion.%s' to be optional", field)
		}
	}

	rawCriterionTypeAttr := criterionObjectAttrs["criterion_type"].(resourceschema.StringAttribute)
	if len(rawCriterionTypeAttr.Validators) == 0 {
		t.Fatalf("expected 'criteria.criterion.criterion_type' to have validators")
	}

	rawFieldAttr := criterionObjectAttrs["field"].(resourceschema.StringAttribute)
	if len(rawFieldAttr.Validators) == 0 {
		t.Fatalf("expected 'criteria.criterion.field' to have validators")
	}

	rawOperandCriterionAttr := criterionObjectAttrs["operand"].(resourceschema.StringAttribute)
	if _, ok := rawOperandCriterionAttr.CustomType.(customTypes.CustomStringType); !ok {
		t.Fatalf("expected 'criteria.criterion.operand' custom type to be CustomStringType, got %T", rawOperandCriterionAttr.CustomType)
	}

	rawValueAttr := criterionObjectAttrs["value"].(resourceschema.StringAttribute)
	if _, ok := rawValueAttr.CustomType.(customTypes.CustomStringType); !ok {
		t.Fatalf("expected 'criteria.criterion.value' custom type to be CustomStringType, got %T", rawValueAttr.CustomType)
	}
}

func TestGroupResourceSchemaSupervisorsIsRequired(t *testing.T) {
	r := &groupresource.GroupResource{}
	resp := &tfresource.SchemaResponse{}

	r.Schema(context.Background(), tfresource.SchemaRequest{}, resp)

	rawSupervisorsAttr, ok := resp.Schema.Attributes["supervisors"]
	if !ok {
		t.Fatalf("expected 'supervisors' attribute in group resource schema")
	}

	supervisorsAttr, ok := rawSupervisorsAttr.(resourceschema.SetAttribute)
	if !ok {
		t.Fatalf("expected 'supervisors' to be schema.SetAttribute, got %T", rawSupervisorsAttr)
	}

	if !supervisorsAttr.Required {
		t.Fatalf("expected 'supervisors' to be required")
	}

	if supervisorsAttr.Optional {
		t.Fatalf("expected 'supervisors' to not be optional")
	}

	if supervisorsAttr.Computed {
		t.Fatalf("expected 'supervisors' to not be computed")
	}

	if len(supervisorsAttr.Validators) < 2 {
		t.Fatalf("expected 'supervisors' validators to include min length and UUID validation")
	}

	if supervisorsAttr.ElementType != types.StringType {
		t.Fatalf("expected 'supervisors' element type to be types.StringType, got %T", supervisorsAttr.ElementType)
	}
}

func TestGroupResourceSchemaGroupTypeHasAllowedValueValidator(t *testing.T) {
	r := &groupresource.GroupResource{}
	resp := &tfresource.SchemaResponse{}

	r.Schema(context.Background(), tfresource.SchemaRequest{}, resp)

	rawGroupTypeAttr, ok := resp.Schema.Attributes["group_type"]
	if !ok {
		t.Fatalf("expected 'group_type' attribute in group resource schema")
	}

	groupTypeAttr, ok := rawGroupTypeAttr.(resourceschema.StringAttribute)
	if !ok {
		t.Fatalf("expected 'group_type' to be schema.StringAttribute, got %T", rawGroupTypeAttr)
	}

	if groupTypeAttr.CustomType != nil {
		t.Fatalf("expected 'group_type' to use regular string type, got custom type %T", groupTypeAttr.CustomType)
	}

	if len(groupTypeAttr.Validators) == 0 {
		t.Fatalf("expected 'group_type' to have allowed-value validators")
	}
}
