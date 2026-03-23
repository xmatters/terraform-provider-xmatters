package group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *GroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Returns a group in your xMatters instance that matches the provided criteria.",
		Attributes:          GroupDataSourceSchema(),
	}
}

// GroupDataSourceSchema returns the schema for the group and groups data sources.
func GroupDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"group_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: describe.GroupIDSearch,
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupID,
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupTargetName,
		},
		"description": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupDescription,
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupStatus,
		},
		"external_key": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupExternalKey,
		},
		"externally_owned": schema.BoolAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupExternallyOwned,
		},
		"allow_duplicates": schema.BoolAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupAllowDuplicates,
		},
		"site": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupSite,
		},
		"observed_by_all": schema.BoolAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupObservedByAll,
		},
		"observers": schema.SetAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupObservers,
			ElementType:         customTypes.CustomStringType{},
		},
		"supervisors": schema.SetAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupSupervisors,
			ElementType:         types.StringType,
		},
		"group_type": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupType,
		},
		"criteria": schema.SingleNestedAttribute{
			Computed:   true,
			Attributes: GroupCriteriaSchema(),
		},
	}
}

// GroupCriteriaSchema is a helper function to simplify the GroupDataSource Schema implementation.
func GroupCriteriaSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"operand": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupCriteriaOperand,
		},
		"criterion": schema.SetNestedAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupCriteriaCriterion,
			NestedObject: schema.NestedAttributeObject{
				Attributes: GroupCriterionSchema(),
			},
		},
	}
}

// GroupCriterionSchema is a helper function to simplify the GroupCriteriaSchema Schema implementation.
func GroupCriterionSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"criterion_type": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupCriterionType,
		},
		"field": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupCriterionField,
		},
		"operand": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupCriterionOperand,
		},
		"value": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupCriterionValue,
		},
	}
}
