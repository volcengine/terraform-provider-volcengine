data "volcengine_iam_entities_policies" "default" {
  policy_name = "AdministratorAccess"
  policy_type = "System"
}
