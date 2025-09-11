package groupRoster

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
func (r *GroupRosterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Returns a groupRoster in your xMatters instance that matches the provided criteria.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: describe.GroupRosterResourceID,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					utils.UUIDValidator{},
				},
			},
			"group": schema.ObjectAttribute{
				Computed:            true,
				MarkdownDescription: describe.GroupRosterResourceGroup,
				AttributeTypes:      utils.GroupReferenceObjectType.AttrTypes,
			},
			"members": schema.SetNestedAttribute{
				Required:            true,
				MarkdownDescription: describe.GroupRosterResourceMembers,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: describe.GroupRosterResourceMemberID,
							Validators: []validator.String{
								utils.UUIDValidator{},
							},
						},
						"member_type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: describe.GroupRosterResourceMemberType,
						},
					},
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
