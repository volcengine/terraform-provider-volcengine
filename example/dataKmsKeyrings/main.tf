data "volcengine_kms_keyrings" "default" {
  keyring_name = ["tf-test-1", "tf-test-2", "tf-test-3"]
  description = ["tf-1", "tf-2"]
}