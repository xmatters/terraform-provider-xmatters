package groups

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/data-sources/services"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils/customTypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *GroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Returns a list of groups in your xMatters instance that match the provided criteria.",
		Attributes: map[string]schema.Attribute{
			"search": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: describe.GroupsSearch,
				Attributes:          GroupsDataSourceSearchSchema(),
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: describe.GroupsFilters,
				Attributes:          GroupsDataSourceFiltersSchema(),
			},
			"options": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: describe.GroupsOptions,
				Attributes:          GroupsDataSourceOptionsSchema(),
			},
			"groups": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: describe.GroupsList,
				NestedObject: schema.NestedAttributeObject{
					Attributes: GroupDataSourceSchema(),
				},
			},
		},
	}
}

// GroupDataSourceSchema returns the schema for the group and groups data sources.
func GroupDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
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
		"timezone": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupTimezone,
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
		"services": schema.SetNestedAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupServices,
			NestedObject: schema.NestedAttributeObject{
				Attributes: services.ServiceDataSourceSchema(),
			},
		},
		"group_type": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.GroupType,
		},
	}
}

// GroupsDataSourceFiltersSchema is a helper function to simplify the GroupsDataSource Schema implementation.
// It represents the Provider's optional search parameters to filter the list of groups returned.
func GroupsDataSourceSearchSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"terms": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.GroupsSearchTerms,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(2),
			},
		},
		"operand": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.GroupsSearchOperand,
			Validators: []validator.String{
				utils.OperandValidator{},
			},
		},
		"fields": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.GroupsSearchFields,
			ElementType:         types.StringType,
		},
	}
}

// GroupsDataSourceSearchSchema is a helper function to simplify the GroupsDataSource Schema implementation.
// It represents the Provider's optional search parameters to filter the list of groups returned.
func GroupsDataSourceFiltersSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"group_type": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.GroupsFilterGroupType,
		},
		"member_exists": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.GroupsFilterMemberExists,
		},
		"members": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.GroupsFilterMembers,
			ElementType:         types.StringType,
		},
		"sites": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.GroupsFilterSites,
			ElementType:         types.StringType,
		},
		"status": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.GroupsFilterStatus,
			Validators: []validator.String{
				utils.StatusValidator{},
			},
		},
		"supervisors": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.GroupsFilterSupervisors,
			ElementType:         types.StringType,
		},
	}
}

// GroupsDataSourceOptionsSchema is a helper function to simplify the GroupsDataSource Schema implementation.
// It represents the Provider's optional search parameters to filter the list of services returned.
func GroupsDataSourceOptionsSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"sort_by": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.GroupsOptionsSortBy,
		},
		"sort_order": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.GroupsOptionsSortOrder,
			Validators: []validator.String{
				stringvalidator.AlsoRequires(path.Expressions{
					path.MatchRoot("options").AtName("sort_by"),
				}...),
				utils.SortOrderValidator{},
			},
		},
	}
}
