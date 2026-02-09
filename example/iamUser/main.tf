resource "volcengine_iam_user" "foo" {
  user_name    = "jonny"
  description  = "test"
  display_name = "name"
  mobile_phone = "+8618800000000"
  email        = "test@example.com"
  tags {
    key   = "key1"
    value = "value1"
  }
}
