resource "volcengine_kms_keyring" "foo" {
  keyring_name   = "tf-test-16"
  description = "tf-test"
  project_name = "default"
}