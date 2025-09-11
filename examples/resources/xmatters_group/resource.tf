// Basic:
resource "xmatters_group" "basic_group" {
  name = "DevOps Team"
}

// Advanced:
resource "xmatters_group" "advanced_group" {
  name                = "SRE Team"
  status              = "ACTIVE"
  description         = "Site Reliability Engineering group"
  group_type          = "ONCALL"
  allow_duplicates    = true
  site                = "481086d8-357a-4279-b7d5-d7dce48fcd12"
  observed_by_all     = false
  observers           = ["observer1", "observer2"]
  use_default_devices = false
  supervisors         = ["481086d8-357a-4279-b7d5-d7dce48fcd12", "12345678-1234-1234-1234-123456789012"]
  external_key        = "sre-group"
  externally_owned    = true
}

