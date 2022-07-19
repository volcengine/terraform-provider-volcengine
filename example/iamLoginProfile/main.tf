resource "volcengine_iam_login_profile" "foo" {
  user_name = "tf-test"
  password = "******"
  login_allowed = true
  password_reset_required = false
}