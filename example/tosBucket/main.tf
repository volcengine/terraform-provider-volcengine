resource "vestack_tos_bucket" "default" {
  bucket_name = "test-xym-tf"
  storage_class ="IA"
  public_acl = "private"
  enable_version = true
  account_acl {
    account_id = "1"
    permission = "READ"
  }
  account_acl {
    account_id = "2001"
    permission = "WRITE_ACP"
  }
}