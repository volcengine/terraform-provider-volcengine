resource "volcengine_iam_user" "foo" {
  user_name    = "jonny-g"
  description  = "test"
  display_name = "name"
  mobile_phone = "17700000000"
  email        = "modify@example.com"
  tags {
    key   = "key1"
    value = "value1"
  }
  tags {
    key   = "key2"
    value = "value2"
  }
}
