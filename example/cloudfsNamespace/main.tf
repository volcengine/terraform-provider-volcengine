resource "volcengine_cloudfs_namespace" "foo" {
  fs_name = "tf-test-fs"
  tos_bucket = "tf-test"
  read_only = true
}
