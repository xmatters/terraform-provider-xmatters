package provider

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
)

// Schema satisfies the provider.Provider interface for the xMatters Provider.
func (p *XMattersProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: describe.ProviderDescription,
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: describe.ProviderBaseURL,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^https:\/\/.+\.xmatters\.com(?:\.au)?$`), "Invalid xMatters URL"),
				},
			},
			"auth": schema.SingleNestedAttribute{
				Required:            true,
				MarkdownDescription: describe.ProviderAuth,
				Attributes: map[string]schema.Attribute{
					"auth_type": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: describe.ProviderAuthType,
						Validators: []validator.String{
							stringvalidator.OneOf("BASIC", "API_TOKEN"),
						},
					},
					"username": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: describe.ProviderUsername,
						Validators: []validator.String{
							utils.BasicAuthValidator{},
							stringvalidator.AlsoRequires(path.Expressions{
								path.MatchRoot("auth").AtName("password"),
							}...),
						},
					},
					"password": schema.StringAttribute{
						Optional:            true,
						Sensitive:           true,
						MarkdownDescription: describe.ProviderPassword,
						Validators: []validator.String{
							utils.BasicAuthValidator{},
							stringvalidator.AlsoRequires(path.Expressions{
								path.MatchRoot("auth").AtName("username"),
							}...),
						},
					},
					"token": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: describe.ProviderToken,
						Validators: []validator.String{
							utils.TokenValidator{},
						},
					},
				},
			},
		},
	}
}
