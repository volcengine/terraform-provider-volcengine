data "volcengine_organization_service_control_policies" "foo" {
  policy_type = "Custom"
  query       = "test"
}