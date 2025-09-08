// Basic:
resource "xmatters_device" "basic_device" {
  // Required attributes
  device_type        = "EMAIL"
  name               = "Work Email"
  owner              = "481086d8-357a-4279-b7d5-d7dce48fcd12"
  priority_threshold = "MEDIUM"
  test_status        = "TESTED"
  timeframes = [
    {
      days                = ["MO", "WE", "FR"]
      duration_in_minutes = 5
      exclude_holidays    = true
      name                = "50/50 Shift"
      start_time          = "08:00"
    },
  ]
  // Required when device_type is EMAIL
  email_address = "mmcbride@your-company.org"
}

// Advanced:
resource "xmatters_device" "advanced_device" {
  // Required attributes
  device_type        = "EMAIL"
  name               = "Work Email"
  owner              = "481086d8-357a-4279-b7d5-d7dce48fcd12"
  priority_threshold = "MEDIUM"
  test_status        = "TESTED"
  timeframes = [
    {
      days                = ["MO", "WE", "FR"]
      duration_in_minutes = 5
      exclude_holidays    = true
      name                = "50/50 Shift"
      start_time          = "08:00"
    },
  ]
  // Required when device_type is EMAIL
  email_address = "mmcbride@your-company.org"
  // Optional attributes
  default_device   = true
  delay            = 15
  external_key     = "mmcbride-email"
  externally_owned = true
  status           = "ACTIVE"
}
