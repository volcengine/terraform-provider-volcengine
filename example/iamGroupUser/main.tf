data "volcengine_iam_group_users" "foo" {
  user_name = "test_user"
}

output "groups" {
  value = data.volcengine_iam_group_users.foo.user_groups
}
