provider "volcengine" {
  region = "cn-beijing"
}
resource "volcengine_kms_key_primary_region" "primary" {
  keyring_name   = "test"
  key_name       = "mrk-Tf-test"
  # Note: The primary region is switched from cn-beijing to cn-shanghai, so cn-beijing will become a replica key 
  # and will no longer support operations such as key rotation and switching the primary region.
  # To continue operating on the key, need to switch the region to cn-shanghai.
  primary_region = "cn-shanghai"
}