resource "volcengine_iam_user_policy_attachment" "foo" {
  user_name = "jonny"
  policy_name = "AdministratorAccess"
  policy_type = "System"
}
