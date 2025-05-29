resource "volcengine_kms_secret" "foo" {
  secret_name = "tf-test1"
  secret_type = "Generic"
  description = "tf-test"
  secret_value = "{\"dasdasd\":\"dasdasd\"}"
}