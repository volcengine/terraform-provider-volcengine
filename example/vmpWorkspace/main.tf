resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-vmp-workspace"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test"
  username                  = "admin123"
  password                  = "Pass123456"
  project_name = "default"
  tags {
    key = "k1"
    value = "v1"
  }
}
