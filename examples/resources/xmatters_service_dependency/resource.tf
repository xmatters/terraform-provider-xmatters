// Create a new service dependency in xMatters using the `xmatters_service_dependency` resource
resource "xmatters_service_dependency" "create_modify_service_dependency" {
  service_id           = "c0b5351e-3848-45c9-b1d3-c3bdad36517a"
  dependent_service_id = "ac523c52-4873-4883-8e51-760eb47cf6f9"
}
