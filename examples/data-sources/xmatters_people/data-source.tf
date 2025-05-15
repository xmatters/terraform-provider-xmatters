// No filter is applied here returning all people
data "xmatters_people" "get_people_list" {
}

// Search for people by terms and fields
data "xmatters_people" "get_people_list_by_search" {
  search = {
    terms   = "Craig James"
    operand = "AND"
    fields = [
      "FIRST_NAME",
      "LAST_NAME"
    ]
  }
}

// Filters applied to limit returned devices 
data "xmatters_people" "get_people_list_by_filter" {
  filters = {
    created_from        = "2024-06-14T16:17:31.450Z"
    created_to          = "2025-01-14T22:25:43.290Z"
    devices_exists      = true
    devices_status      = "ACTIVE"
    devices_test_status = "TESTED"
  }
  options = {
    sort_by    = "FIRST_LAST_NAME"
    sort_order = "ASCENDING"
  }
}
