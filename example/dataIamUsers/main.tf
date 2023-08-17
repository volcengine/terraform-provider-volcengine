resource "volcengine_iam_user" "foo" {
  user_name = "acc-test-user"
  description = "acc test"
  display_name = "name"
}
data "volcengine_iam_users" "foo"{
  user_names = [volcengine_iam_user.foo.user_name]
}