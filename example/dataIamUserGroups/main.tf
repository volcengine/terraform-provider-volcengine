resource "volcengine_iam_user_group" "foo" {
  user_group_name = "acc-test-group"
  description = "acc-test"
  display_name = "acc-test"
}

data "volcengine_iam_user_groups" "foo" {
  query = "acc-test"
}