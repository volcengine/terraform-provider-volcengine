resource "volcengine_iam_security_config" "foo" {
  user_name                 = "jonny"
  safe_auth_type            = "email"
  safe_auth_exempt_duration = 11
}
