terraform {
  required_providers {
    xmatters = {
      source  = "github.com/xmatters/terraform-provider-xmatters"
      version = "~> 0.1.0"
    }
  }
}

#Configure the provider for Basic Auth
provider "xmatters" {
  base_url = "https://example.xmatters.com"
  auth = {
    auth_type = "BASIC"
    username  = "mmcbride"
    password  = "SuperSecretPassword!"
  }
}

#Configure provider for OAuth token
provider "xmatters" {
  base_url = "https://example.xmatters.com"
  auth = {
    auth_type = "API_TOKEN"
    token     = var.oauth_token
  }
}
