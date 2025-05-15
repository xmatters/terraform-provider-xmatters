package person

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *PersonDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.PersonDataSourceDescription,
		Attributes:          PersonDataSourceSchema(),
	}
}

// PersonDataSourceSchema returns the schema for the person and persons data sources.
func PersonDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"person_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: describe.PersonIDSearch,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(3),
			},
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonID,
		},
		"target_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonTargetName,
		},
		"first_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonFirstName,
		},
		"last_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonLastName,
		},
		"roles": schema.SetAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonRoles,
			ElementType:         types.StringType,
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonStatus,
		},
		"web_login": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonWebLogin,
		},
		"site": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonSite,
		},
		"timezone": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonTimezone,
		},
		"language": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonLanguage,
		},
		"supervisors": schema.SetAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonSupervisors,
			ElementType:         types.StringType,
		},
		"phone_login": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonPhoneLogin,
		},
		"license_type": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonLicenseType,
		},
		"external_key": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonExternalKey,
		},
		"externally_owned": schema.BoolAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonExternallyOwned,
		},
		"last_login": schema.StringAttribute{
			Computed: true,
		},
	}
}
