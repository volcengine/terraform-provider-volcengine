resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "*******"
}

data "volcengine_vmp_workspaces" "foo"{
  ids = [volcengine_vmp_workspace.foo.id]
}