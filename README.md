# xMatters Terraform Provider

This Terraform provider enables users to interact with the xMatters platform via its REST API. The provider allows management of various xMatters objects such as users, groups, devices, and services using Terraform configuration files.

## Features

- **Resource Management**: Easily create, update, and delete xMatters resources using Terraform configuration files.
- **Idempotent Operations**: Perform operations safely and consistently with idempotent actions, ensuring predictable outcomes.
- **State Management**: Maintain state consistency and track changes to xMatters resources within Terraform state files.
- **API Compatibility**: Leverage the full potential of the Everbridge xMatters REST API to manage a wide range of objects and configurations.
- **Automation**: Automate provisioning and management of xMatters resources alongside your infrastructure provisioning process.

## Prerequisites

- [Terraform](https://www.terraform.io/downloads.html) >= 1.2 installed locally.
  - The [latest version](https://developer.hashicorp.com/terraform/downloads?product_intent=terraform) is recommended for optimal compatibility with the xMatters provider.
- Access to an xMatters instance with appropriate permissions to manage resources.
- API credentials configured with necessary permissions to interact with the xMatters REST API.
- [Go](https://golang.org/doc/install) >= 1.22.

## Usage

### Getting Started

1. **Install the Provider**: Add the xMatters Terraform provider to your Terraform configuration.
2. **Configure Provider**: Configure authentication and endpoint details for your xMatters instance.
3. **Define Resources**: Define xMatters resources and their configurations within your Terraform configuration files.
4. **Apply Changes**: Execute `terraform init`, `terraform plan`, and `terraform apply` commands to provision or modify xMatters resources as per your configuration.

Refer to the [documentation](./docs) for detailed usage instructions and resource reference.

[xMatters REST API documentation](https://help.xmatters.com/xmapi) is also available for non-Terraform or service specific information.

### Authentication

The xMatters REST API supports two methods for authentication to an xMatters instance through API Tokens or HTTP Basic authentication.

The desired authentication method must be defined in the provider construction with the `auth.type` field. The value must be one of `BASIC` or `API_TOKEN` and is a required field.

Token authentication authorizes requests by passing an OAuth token. This avoids storage of personal credentials in applications or files. Tokens can be revoked at any time. For more information about obtaining and working with OAuth access tokens, see [OAuth](#oauth). HTTP Basic authentication is used to authorize requests by passing the Username and Password of an xMatters account, or a personal xMatters API Key and Secret.

1. OAuth using `auth.token`
2. Basic authentication using `auth.username` with `auth.password`

Local user environment variables can also be used. For more on available authentication methods, check the [docs](./docs/index.md)

## Example Usage

### Example Provider Authentication with OAuth 2.0 Token

```hcl
terraform {
  required_providers {
    xmatters = {
      source  = "xmatters/xmatters"
      version = "~> 0.1"
    }
  }
}

provider "xmatters" {
  base_url = "https://example.xmatters.com"
  auth = {
    type  = "API_TOKEN"
    token = var.xmatters_token
  }
}
```

### Example Provider Authentication with Basic Auth

```hcl
terraform {
  required_providers {
    xmatters = {
      source  = "xmatters/xmatters"
    }
  }
}

provider "xmatters" {
  base_url = "https://example.xmatters.com"
  auth = {
    type     = "BASIC"
    username = var.xmatters_username
    password = var.xmatters_password
  }
}
```

## Resources

- [Documentation](./docs): Comprehensive documentation covering provider usage, configuration, and resource reference.
- [Issues](https://support.xmatters.com/): Report issues, or suggest features.

## OAuth

To obtain an OAuth 2.0 access token and a refresh token, provide your client ID and the credentials of an account that you want to use to authorize requests. The authorization token provides the same xMatters permissions as the user account used to create the token.

Further instruction on acquiring your personal access token can be found by logging into your xMatters instance and visiting the `Workflows` tab, under the `Authentication` section. Here you can see your `client_id` as well as explicit steps to create or refresh an OAuth token for use in the provider.

Once you have obtained the token, you can use it in the provider configuration.

# Building The Provider (For Developers)

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine as described in the prerequisites.

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of tests, run `TF_ACC=1 go test -v ./...`.
