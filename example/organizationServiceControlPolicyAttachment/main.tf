resource "volcengine_organization_service_control_policy" "foo" {
  policy_name = "tfpolicy11"
  description = "tftest1"
  statement   = "{\"Statement\":[{\"Effect\":\"Deny\",\"Action\":[\"ecs:RunInstances\"],\"Resource\":[\"*\"]}]}"
}

resource "volcengine_organization_service_control_policy_attachment" "foo" {
  policy_id   = volcengine_organization_service_control_policy.foo.id
  target_id   = "21*********94"
  target_type = "Account"
}

resource "volcengine_organization_service_control_policy_attachment" "foo1" {
  policy_id   = volcengine_organization_service_control_policy.foo.id
  target_id   = "73*********9"
  target_type = "OU"
}