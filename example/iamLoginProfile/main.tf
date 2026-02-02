resource "volcengine_iam_login_profile" "foo" {
  user_name                 = "jonny"
  password                  = "Password@123"
  login_allowed             = true
  password_reset_required   = true
  safe_auth_flag            = true
  safe_auth_type            = "phone"
  safe_auth_exempt_required = 1
  safe_auth_exempt_unit     = 1
  safe_auth_exempt_duration = 1
}
