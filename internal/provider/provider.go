package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	deviceDS "github.com/xmatters/terraform-provider-xmatters/internal/data-sources/device"
	devicesDS "github.com/xmatters/terraform-provider-xmatters/internal/data-sources/devices"
	groupDS "github.com/xmatters/terraform-provider-xmatters/internal/data-sources/group"
	groupsDS "github.com/xmatters/terraform-provider-xmatters/internal/data-sources/groups"
	peopleDS "github.com/xmatters/terraform-provider-xmatters/internal/data-sources/people"
	personDS "github.com/xmatters/terraform-provider-xmatters/internal/data-sources/person"
	serviceDS "github.com/xmatters/terraform-provider-xmatters/internal/data-sources/service"
	servicesDS "github.com/xmatters/terraform-provider-xmatters/internal/data-sources/services"
	siteDS "github.com/xmatters/terraform-provider-xmatters/internal/data-sources/site"
	sitesDS "github.com/xmatters/terraform-provider-xmatters/internal/data-sources/sites"
	userQuotasDS "github.com/xmatters/terraform-provider-xmatters/internal/data-sources/userQuotas"
	"github.com/xmatters/terraform-provider-xmatters/internal/describe"
	deviceR "github.com/xmatters/terraform-provider-xmatters/internal/resources/device"
	groupR "github.com/xmatters/terraform-provider-xmatters/internal/resources/group"
	groupRosterR "github.com/xmatters/terraform-provider-xmatters/internal/resources/groupRoster"
	personR "github.com/xmatters/terraform-provider-xmatters/internal/resources/person"
	serviceR "github.com/xmatters/terraform-provider-xmatters/internal/resources/service"
	serviceDependency "github.com/xmatters/terraform-provider-xmatters/internal/resources/serviceDependency"
	siteR "github.com/xmatters/terraform-provider-xmatters/internal/resources/site"
)

// Ensure XMattersProvider satisfies various provider interfaces.
var _ provider.Provider = &XMattersProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &XMattersProvider{
			Version: version,
		}
	}
}

// XMattersProvider defines the provider implementation.
type XMattersProvider struct {
	// Version is set to the provider Version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	Version string
}

// Metadata satisfies the provider.Provider interface for the xMatters Provider.
func (p *XMattersProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "xmatters"
	resp.Version = p.Version
}

// Configure satisfies the provider.Provider interface for the xMatters Provider.
func (p *XMattersProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data XMattersProviderModel
	var config Config
	var err error

	// Retrieve provider data from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known, non-null value.
	if data.BaseURL.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("base_url"),
			"Unknown Base URL",
			"The provider cannot create the XMatters API client as there is an unknown configuration value for the Base URL.",
		)
	}
	config.BaseURL, err = GetBaseURLInput(data.BaseURL)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Retrieve Base URL",
			"An unexpected error occurred when retrieving the Base URL from the provider configuration. ",
		)
	}

	if data.Auth.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("auth"),
			"Unknown Auth Configuration",
			"The provider cannot create the XMatters API client as there is an unknown configuration value for the Auth Configuration.",
		)
	}
	config.Auth, err = GetAuthInput(data.Auth)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Retrieve Auth Configuration",
			"An unexpected error occurred when retrieving the Auth Configuration from the provider configuration. ",
		)
	}

	// If there are any errors, return early
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating xMatters client")

	// Create a new xMatters client using the configuration values
	client, err := config.Client()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create xMatters API Client",
			"An unexpected error occurred when creating the xMatters API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"xMatters Client Error: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
	tflog.Info(ctx, "Configured xMatters client", map[string]any{"success": true})
}

// Resources satisfies the provider.Provider interface for the xMatters Provider.
func (p *XMattersProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		serviceR.NewServiceResource,
		serviceDependency.NewServiceDependencyResource,
		personR.NewPersonResource,
		groupR.NewGroupResource,
		groupRosterR.NewGroupRosterResource,
		siteR.NewSiteResource,
		deviceR.NewDeviceResource,
	}
}

// DataSources satisfies the provider.Provider interface for the xMatters Provider.
func (p *XMattersProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		personDS.NewPersonDataSource,
		peopleDS.NewPeopleDataSource,
		serviceDS.NewServiceDataSource,
		servicesDS.NewServicesDataSource,
		userQuotasDS.NewUserQuotasDataSource,
		groupDS.NewGroupDataSource,
		groupsDS.NewGroupsDataSource,
		siteDS.NewSiteDataSource,
		sitesDS.NewSitesDataSource,
		deviceDS.NewDeviceDataSource,
		devicesDS.NewDevicesDataSource,
	}
}

// GetBaseURLInput returns the value of the input string if it is not null, otherwise it returns the value of the given environment variable.
func GetBaseURLInput(input basetypes.StringValue) (string, error) {
	var out string
	if !input.IsNull() {
		out = input.ValueString()
	} else {
		if value, ok := os.LookupEnv(describe.APIBaseURL); ok {
			out = value
		} else {
			return "", fmt.Errorf("base_url is not set, uanble to retrieve from local ENV")
		}
	}
	return out, nil
}

// GetOptions returns a slice of `xmatters.Option` from the provider configuration to be used while instantiating the client.
func GetAuthInput(input types.Object) (AuthConfig, error) {
	auth := AuthConfig{}
	if !input.IsNull() {
		for name, value := range input.Attributes() {
			tfVal, err := value.ToTerraformValue(context.Background())
			if err != nil {
				tflog.Warn(context.Background(), "Error converting Terraform Value to Go Value", map[string]any{"error": err.Error()})
			}
			switch name {
			case "auth_type":
				var authType *string
				err := tfVal.As(&authType)
				if err != nil {
					tflog.Warn(context.Background(), "Error converting Terraform Value to String", map[string]any{"error": err.Error()})
				}
				auth.AuthType = *authType
			case "username":
				var user *string
				err := tfVal.As(&user)
				if err != nil {
					tflog.Warn(context.Background(), "Error converting Terraform Value to String", map[string]any{"error": err.Error()})
				}
				if user == nil && auth.AuthType == "BASIC" {
					if value, ok := os.LookupEnv(describe.APIUserEnvVarKey); ok {
						user = &value
					} else {
						return auth, fmt.Errorf("username is not set, uanble to retrieve from local ENV")
					}
				}
				auth.Username = user
			case "password":
				var pass *string
				err := tfVal.As(&pass)
				if err != nil {
					tflog.Warn(context.Background(), "Error converting Terraform Value to String", map[string]any{"error": err.Error()})
				}
				if pass == nil && auth.AuthType == "BASIC" {
					if value, ok := os.LookupEnv(describe.APIPassEnvVarKey); ok {
						pass = &value
					} else {
						return auth, fmt.Errorf("password is not set, uanble to retrieve from local ENV")
					}
				}
				auth.Password = pass
			case "token":
				var token *string
				err := tfVal.As(&token)
				if err != nil {
					tflog.Warn(context.Background(), "Error converting Terraform Value to String", map[string]any{"error": err.Error()})
				}
				if token == nil && auth.AuthType == "API_TOKEN" {
					if value, ok := os.LookupEnv(describe.APIToken); ok {
						token = &value
					} else {
						return auth, fmt.Errorf("token is not set, uanble to retrieve from local ENV")
					}
				}
				auth.Token = token
			}
		}
	}
	return auth, nil
}
