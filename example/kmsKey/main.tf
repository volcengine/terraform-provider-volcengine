resource "volcengine_kms_keyring" "foo" {
  keyring_name   = "tf-test"
  description = "tf-test"
  project_name = "default"
}

resource "volcengine_kms_key" "foo" {
  keyring_name   = volcengine_kms_keyring.foo.keyring_name
  key_name = "mrk-tf-key-mod"
  description = "tf test key-mod"
  tags {
    key = "tfkey3"
    value = "tfvalue3"
  }
}