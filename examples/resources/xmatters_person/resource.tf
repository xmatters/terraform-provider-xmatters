// Basic:
resource "xmatters_person" "basic_create" {
  target_name  = "mmcbride"
  web_login    = "mmcbride"
  first_name   = "Margot"
  last_name    = "McBride"
  roles        = "Full Access User"
  timezone     = "US/Pacific"
  language     = "en"
  license_type = "FULL_USER"
  site         = "1d7505ce-8d50-411d-9d6e-8cfbc365b5fe"
  supervisors  = []
}

// Advanced:
resource "xmatters_person" "advanced_person" {
  target_name = "mmcbride"
  first_name  = "Margot"
  last_name   = "McBride"
  roles = [
    "Developer",
    "Full Access User"
  ]
  status    = "ACTIVE"
  web_login = "mmcbride@your-company.org"
  site      = "1d7505ce-8d50-411d-9d6e-8cfbc365b5fe"
  timezone  = "US/Pacific"
  language  = "en"
  supervisors = [
    "481086d8-357a-4279-b7d5-d7dce48fcd12",
    "545686d8-3491-4a12-ddb7-a33239e82bc7"
  ]
  password             = "PassWord123!"
  force_password_reset = true
  phone_login          = "123456789"
  phone_pin            = "9728"
  license_type         = "FULL_USER"
  external_key         = "c5aa0fff0a0a0aa7009a39da035ea396"
  externally_owned     = true
}
