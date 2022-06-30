resource "volcengine_login_profile" "foo" {
  user_name = "tf-test1"
  password = "*****"
  login_allowed = true
  password_reset_required = true
}