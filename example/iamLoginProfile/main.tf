resource "volcengine_iam_user" "foo" {
  user_name = "acc-test-user"
  description = "acc-test"
  display_name = "name"
}

resource "volcengine_iam_login_profile" "foo" {
  user_name = volcengine_iam_user.foo.user_name
  password = "93f0cb0614Aab12"
  login_allowed = true
  password_reset_required = false
}
