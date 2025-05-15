package userQuotas

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *UserQuotasDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.UserQuotasDataSourceDescription,
		Attributes:          UserQuotasDataSourceSchema(),
	}
}

// UserQuotasDataSourceSchema returns the schema for the user_license_quotas data sources.
func UserQuotasDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"stakeholder_users_enabled": schema.BoolAttribute{
			Computed:            true,
			MarkdownDescription: describe.UserQuotasStakeholderUsersEnabled,
		},
		"stakeholder_users": schema.SingleNestedAttribute{
			Computed:            true,
			MarkdownDescription: describe.UserQuotasStakeholderUsers,
			Attributes:          QuotaDetailsSchema(),
		},
		"full_users": schema.SingleNestedAttribute{
			Computed:            true,
			MarkdownDescription: describe.UserQuotasFullUsers,
			Attributes:          QuotaDetailsSchema(),
		},
	}
}

func QuotaDetailsSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"total": schema.Int64Attribute{
			Computed:            true,
			MarkdownDescription: describe.UserQuotasQuotaTotal,
		},
		"active": schema.Int64Attribute{
			Computed:            true,
			MarkdownDescription: describe.UserQuotasQuotaActive,
		},
		"unused": schema.Int64Attribute{
			Computed:            true,
			MarkdownDescription: describe.UserQuotasQuotaUnused,
		},
	}
}
