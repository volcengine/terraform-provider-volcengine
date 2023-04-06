data "volcengine_cr_vpc_endpoints" "default" {
  registry = "enterprise-1"
  statuses = ["Enabled", "Enabling", "Disabling", "Failed"]
}