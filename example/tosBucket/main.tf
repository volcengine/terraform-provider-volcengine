resource "volcengine_tos_bucket" "default" {
  bucket_name          = "tf-acc-test-bucket-0123-3"
#  storage_class        = "IA"
  public_acl           = "private"
  az_redundancy        = "multi-az"
  enable_version       = true
  bucket_acl_delivered = true
  account_acl {
    account_id = "1"
    permission = "READ"
  }
  account_acl {
    account_id = "2001"
    permission = "WRITE_ACP"
  }
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
