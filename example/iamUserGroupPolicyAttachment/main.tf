resource "volcengine_iam_user_group_policy_attachment" "foo" {
  policy_name = "test"
  policy_type = "Custom"
  user_group_name = "tf-test"
}
