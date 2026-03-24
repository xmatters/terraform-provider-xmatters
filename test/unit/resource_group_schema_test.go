package test

import (
	"context"
	"testing"

	tfresource "github.com/hashicorp/terraform-plugin-framework/resource"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	groupresource "github.com/xmatters/terraform-provider-xmatters/internal/resources/group"
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
}
