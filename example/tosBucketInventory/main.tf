resource "volcengine_tos_bucket_inventory" "foo" {
  bucket_name              = "terraform-demo"
  inventory_id             = "acc-test-inventory"
  is_enabled               = true
  included_object_versions = "All"
  schedule {
    frequency = "Weekly"
  }
  filter {
    prefix = "test-tf"
  }
  optional_fields {
    field = ["Size", "StorageClass", "CRC64"]
  }
  destination {
    tos_bucket_destination {
      format     = "CSV"
      account_id = "21000*****"
      bucket     = "terraform-demo"
      prefix     = "tf-test-prefix"
      role       = "TosArchiveTOSInventory"
    }
  }
}
