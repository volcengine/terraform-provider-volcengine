# Only create a backup key in the replica region;
# Next, managing key requires the use of resource "volcengine_kms_key".
resource "volcengine_kms_replicate_key" "replica" {
  keyring_name   = "test"
  key_name       = "mrk-Tf-Test-1"
  replica_region = "cn-shanghai"
  description    = "replica description"
  tags {
    key = "tfk1"
    value = "tfv1"
  }
}