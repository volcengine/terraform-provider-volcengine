data "volcengine_kms_keyrings" "default" {
  filters = "[{\"Key\":\"KeyringName\",\"Values\":[\"tf-test-11\"]}]"
}