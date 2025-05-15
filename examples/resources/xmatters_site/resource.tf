// Basic:
resource "xmatters_site" "basic_site" {
  name     = "West Coast Headquarters"
  country  = "USA"
  language = "en"
  timezone = "US/Los Angeles"
}

// Advanced:
resource "xmatters_site" "advanced_site" {
  address1         = "12345 Alcosta Blvd"
  address2         = "Suite 1234"
  city             = "San Ramon"
  country          = "USA"
  external_key     = "a22f2798-fbf6-4270-bdd0-e9c476c846bc"
  externally_owned = true
  language         = "en"
  latitude         = 39.77
  longitude        = -121.95
  name             = "West Coast Headquarters"
  postal_code      = "94583"
  state            = "CA"
  status           = "ACTIVE"
  timezone         = "US/Los Angeles"
}
