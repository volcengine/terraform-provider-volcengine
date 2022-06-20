resource "vestack_tos_bucket" "default" {
  bucket_name = "test-xym-tf"
  tos_storage_class ="IA"
  tos_acl = "public-read"
}