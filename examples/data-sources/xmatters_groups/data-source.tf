// No filter is applied here returning all groups
data "xmatters_groups" "get_all_groups" {}

// Search for groups by terms and fields
data "xmatters_groups" "get_groups_by_search" {
  search = {
    terms   = "Developers"
    operand = "AND"
    fields  = "NAME"
  }
}

// Filters applied to limit returned groups
data "xmatters_groups" "get_groups_by_filter" {
  filters = {
    group_type    = "ONCALL"
    member_exists = true
    members = [
      "mmcbride"
    ]
    sites = [
      "HQ",
      "Remote"
    ]
    status = "ACTIVE"
    supervisors = [
      "481086d8-357a-4279-b7d5-d7dce48fcd12",
      "12345678-1234-1234-1234-123456789012"
    ]
  }
}

// Options to sort the returned list of groups
data "xmatters_groups" "get_groups_sorted" {
  sort = {
    sort_by    = "NAME"
    sort_order = "ASCENDING"
  }
}
