resource "volcengine_image_import" "foo" {
  platform    = "CentOS"
  url         = "https://*****_system.qcow2"
  image_name  = "acc-test-image"
  description = "acc-test"
  boot_mode    = "UEFI"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
