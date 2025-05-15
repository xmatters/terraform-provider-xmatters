package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/xmatters/xmatters-go"
)

// Config is a struct that holds the configuration for the xMatters provider.
type Config struct {
	BaseURL string
	Auth    AuthConfig
}

// AuthConfig is a struct that holds the authentication configuration for the xMatters provider.
type AuthConfig struct {
	AuthType string  `tfsdk:"auth_type" mapstructure:"auth_type"`
	Username *string `tfsdk:"username" mapstructure:"username"`
	Password *string `tfsdk:"password" mapsctructure:"password"`
	Token    *string `tfsdk:"token" mapstructure:"token"`
}

// Client returns a new xMatters API client based on the configuration.
func (c *Config) Client() (*xmatters.XMattersAPI, error) {
	var err error
	var client *xmatters.XMattersAPI
	options := append([]xmatters.Option{}, xmatters.WithBaseURL(c.BaseURL))

	switch c.Auth.AuthType {
	case "BASIC":
		client, err = xmatters.NewWithBasicAuth(&c.BaseURL, c.Auth.Username, c.Auth.Password, options...)
		if err != nil {
			return nil, fmt.Errorf("error creating new xMatters client: %w", err)
		}
	case "API_TOKEN":
		client, err = xmatters.NewWithToken(&c.BaseURL, c.Auth.Token, options...)
		if err != nil {
			return nil, fmt.Errorf("error creating new xMatters client: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported authentication type: %s", c.Auth.AuthType)
	}

	tflog.Info(context.Background(), fmt.Sprintf("xMatters Client configured with options: %v", options))

	return client, nil
}
