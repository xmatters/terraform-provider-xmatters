// No filter is applied here returning all services
data "xmatters_services" "get_service_list" {}

// Limit the returned services by ownership
data "xmatters_services" "get_service_list_by_owner" {
  owner = "DBA_Admins"
}

// Search for services by terms and fields
data "xmatters_services" "get_service_list_with_search" {
  owner = "DBA_Admins"
  search = {
    terms   = "database postgres"
    operand = "OR"
    fields  = ["NAME"]
  }
}
