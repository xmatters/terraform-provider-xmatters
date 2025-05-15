package describe

const (
	ProviderDescription = "xMatters is a digital service reliability platform that helps users reduce the business impact of service issues and major incidents. The xMatters provider is used to configure infrastructure in xMatters using the <a href=\"https://help.xmatters.com/xmapi/index.html\" target=\"_blank\"> xMatters REST API</a>. Documentation is available for the following supported Data Sources and Resources."
	ProviderAuth        = "Authentication settings for your xMatters instance."
	ProviderBaseURL     = "Base URL of the xMatters instance. For example, https://<company>.xmatters.com"
	ProviderAuthType    = "Authentication method used to connect to the xMatters instance. Available options are BASIC and API_TOKEN."
	ProviderUsername    = "xMatters username or API key of the authenticating user. Basic authentication accepts either either a username or API key."
	ProviderPassword    = "xMatters password or API secret of the authenticating user. Basic authentication accepts either either a password or API secret."
	ProviderToken       = "Required only for OAuth 2.0 Token authentication."
	// Provider env var definitions
	APIBaseURL       = "XMATTERS_BASE_URL"
	APIUserEnvVarKey = "XMATTERS_USERNAME"
	APIPassEnvVarKey = "XMATTERS_PASSWORD"
	APIToken         = "XMATTERS_TOKEN"
)
