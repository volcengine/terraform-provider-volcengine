resource "vestack_tos_object" "default" {
  bucket_name = "test-xym-1"
  object_name = "demo_xym"
  file_path = "/Users/bytedance/Work/Go/build/test.txt"
#  storage_class ="IA"
  public_acl = "private"
  encryption = "AES256"
  #content_type = "text/plain"
  account_acl {
    account_id = "1"
    permission = "READ"
  }
  account_acl {
    account_id = "2001"
    permission = "WRITE_ACP"
  }
#  lifecycle {
#    ignore_changes = ["file_path"]
#  }
}