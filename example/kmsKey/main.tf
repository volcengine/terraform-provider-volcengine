resource "volcengine_kms_key" "foo" {
  keyring_name   = "tf-test-16"
  key_name = "mrk-tf-key-mod"
  description = "tf test key-mod"
  tags {
    key = "tfkey3"
    value = "tfvalue3"
  }
#  multi_region = true
}