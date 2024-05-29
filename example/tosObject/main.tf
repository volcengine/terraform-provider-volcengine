resource "volcengine_tos_object" "default" {
  bucket_name = "tf-acc-test-bucket"
  object_name = "tf-acc-test-object"
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
  tags {
    key = "k1"
    value = "v1"
  }
#  lifecycle {
#    ignore_changes = ["file_path"]
#  }
}