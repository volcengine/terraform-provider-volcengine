resource "volcengine_kms_keyring" "foo" {
  keyring_name   = "tf-test"
  description = "tf-test"
  project_name = "default"
}