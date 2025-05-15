// No filter is applied here returning all sites
data "xmatters_sites" "get_site_list" {}

// Filters applied to limit returned sites
data "xmatters_sites" "get_sites_with_filters" {
  filters = {
    country  = "CANADA"
    geocoded = true
    status   = "ACTIVE"
  }
}

// Search for sites by terms and fields
data "xmatters_sites" "get_sites_with_search" {
  search = {
    terms   = "Headquarters"
    operand = "AND"
    fields  = "NAME"
  }
}
