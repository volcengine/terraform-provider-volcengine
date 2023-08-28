data "volcengine_cloudfs_ns_quotas" "default" {
  fs_names = ["tffile", "tftest2"]
}