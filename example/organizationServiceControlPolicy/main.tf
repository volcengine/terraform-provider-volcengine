resource "volcengine_organization_service_control_policy" "foo" {
  policy_name = "tfpolicy11"
  description = "tftest1"
  statement   = "{\"Statement\":[{\"Effect\":\"Deny\",\"Action\":[\"ecs:RunInstances\"],\"Resource\":[\"*\"]}]}"
}

resource "volcengine_organization_service_control_policy" "foo2" {
  policy_name = "tfpolicy21"
  statement   = "{\"Statement\":[{\"Effect\":\"Deny\",\"Action\":[\"ecs:DeleteInstance\"],\"Resource\":[\"*\"]}]}"
}