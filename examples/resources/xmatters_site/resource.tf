// Basic:
resource "xmatters_site" "basic_site" {
  name     = "West Coast Headquarters"
  country  = "USA"
  language = "en"
  timezone = "US/Pacific"
}

// Advanced:
resource "xmatters_site" "advanced_site" {
  address1    = "12345 Alcosta Blvd"
  address2    = "Suite 1234"
  city        = "San Ramon"
  country     = "USA"
  language    = "en"
  latitude    = 39.77
  longitude   = -121.95
  name        = "West Coast Headquarters"
  postal_code = "94583"
  state       = "CA"
  status      = "ACTIVE"
  timezone    = "US/Pacific"
}
