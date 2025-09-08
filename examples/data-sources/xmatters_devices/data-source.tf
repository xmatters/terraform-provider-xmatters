// No filter is applied here, returning all devices
data "xmatters_devices" "get_device_list" {}

// Filters applied to limit returned devices
data "xmatters_devices" "get_device_list_with_filters" {
  filters = {
    device_status = "ACTIVE"
    device_type   = ["EMAIL"]
  }
}
