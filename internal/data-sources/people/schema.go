package people

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *PeopleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.PeopleDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"search": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: describe.PeopleSearch,
				Attributes:          PeopleDataSourceSearchSchema(),
			},
			"filters": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: describe.PeopleFilters,
				Attributes:          PeopleDataSourceFiltersSchema(),
			},
			"options": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: describe.PeopleOptions,
				Attributes:          PeopleDataSourceOptionsSchema(),
			},
			"people": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: describe.PeopleList,
				NestedObject: schema.NestedAttributeObject{
					Attributes: PersonDataSourceSchema(),
				},
			},
		},
	}
}

// PersonDataSourceSchema is a helper function to simplify the PeopleDataSource Schema implementation.
// It represents the Provider's implmentation of the xMatters Person object.
func PersonDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonID,
		},
		"external_key": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonExternalKey,
		},
		"externally_owned": schema.BoolAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonExternallyOwned,
		},
		"first_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonFirstName,
		},
		"language": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonLanguage,
		},
		"last_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonLastName,
		},
		"license_type": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonLicenseType,
		},
		"phone_login": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonPhoneLogin,
		},
		"roles": schema.SetAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonRoles,
			ElementType:         types.StringType,
		},
		"site": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonSite,
		},
		"status": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonStatus,
		},
		"supervisors": schema.SetAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonSupervisors,
			ElementType:         types.StringType,
		},
		"target_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonTargetName,
		},
		"timezone": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonTimezone,
		},
		"last_login": schema.StringAttribute{
			Computed: true,
		},
		"web_login": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: describe.PersonWebLogin,
		},
	}
}

// PeopleDataSourceSearchSchema is a helper function to simplify the PeopleDataSource Schema implementation.
// It represents the Provider's optional search parameters to filter the list of services returned.
func PeopleDataSourceSearchSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"terms": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleSearchTerms,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(2),
			},
		},
		"operand": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleSearchOperand,
			Validators: []validator.String{
				utils.OperandValidator{},
			},
		},
		"fields": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleSearchFields,
			ElementType:         types.StringType,
		},
	}
}

// PeopleDataSourceFiltersSchema is a helper function to simplify the PeopleDataSource Schema implementation.
// It represents the Provider's optional search parameters to filter the list of services returned.
func PeopleDataSourceFiltersSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"created_from": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersCreatedFrom,
			Validators: []validator.String{
				utils.TimestampValidator{},
			},
		},
		"created_to": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersCreatedTo,
			Validators: []validator.String{
				utils.TimestampValidator{},
			},
		},
		"created_before": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersCreatedBefore,
			Validators: []validator.String{
				utils.TimestampValidator{},
			},
		},
		"created_after": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersCreatedAfter,
			Validators: []validator.String{
				utils.TimestampValidator{},
			},
		},
		"devices_exists": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersDevices,
		},
		"devices_email_exists": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersDevicesEmail,
		},
		"devices_failsafe_exists": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersDevicesFailsafe,
		},
		"devices_mobile_exists": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersDevicesMobile,
		},
		"devices_sms_exists": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersDevicesSMS,
		},
		"devices_voice_exists": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersDevicesVoice,
		},
		"devices_status": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersDevicesStatus,
			Validators: []validator.String{
				utils.StatusValidator{},
			},
		},
		"devices_test_status": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersDevicesTestStatus,
		},
		"email_address": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersEmailAddress,
			Validators: []validator.String{
				utils.EmailValidator{},
			},
		},
		"first_name": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersFirstName,
		},
		"last_name": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersLastName,
		},
		"groups": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersGroups,
			ElementType:         types.StringType,
		},
		"groups_exists": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersGroupsExists,
		},
		"license_type": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersLicenseType,
		},
		"phone_number": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersPhoneNumber,
			Validators: []validator.String{
				utils.PhoneNumberValidator{},
			},
		},
		"roles": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersRoles,
			ElementType:         types.StringType,
		},
		"site": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersSite,
		},
		"status": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersStatus,
			Validators: []validator.String{
				utils.StatusValidator{},
			},
		},
		"supervisors": schema.ListAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersSupervisors,
			ElementType:         types.StringType,
		},
		"supervisors_exists": schema.BoolAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersSupervisorsExists,
		},
		"target_name": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersTargetName,
		},
		"web_login": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleFiltersWebLogin,
			Validators: []validator.String{
				stringvalidator.LengthAtMost(100),
			},
		},
	}
}

// PeopleDataSourceOptionsSchema is a helper function to simplify the PeopleDataSource Schema implementation.
// It represents the Provider's optional search parameters to filter the list of services returned.
func PeopleDataSourceOptionsSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"sort_by": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleOptionsSortBy,
		},
		"sort_order": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: describe.PeopleOptionsSortOrder,
			Validators: []validator.String{
				stringvalidator.AlsoRequires(path.Expressions{
					path.MatchRoot("options").AtName("sort_by"),
				}...),
				utils.SortOrderValidator{},
			},
		},
	}
}
