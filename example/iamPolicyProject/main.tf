resource "volcengine_iam_policy_project" "foo" {
  principal_type = "User"
  principal_name = "jonny"
  policy_type = "Custom"
  policy_name = "restart-oas-ecs"
  project_names = ["default"]
}
