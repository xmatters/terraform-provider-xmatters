package serviceDependency

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
func (r *ServiceDependencyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.ServiceDependencyResourceDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: describe.ServiceDependencyID,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"service_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: describe.ServiceDependencyServiceID,
				Validators: []validator.String{
					utils.UUIDValidator{},
				},
			},
			"dependent_service_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: describe.ServiceDependencyDependentServiceID,
				Validators: []validator.String{
					utils.UUIDValidator{},
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
