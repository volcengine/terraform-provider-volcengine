resource "volcengine_nas_permission_group" "foo" {
  permission_group_name = "acc-test"
  description = "acctest"
  permission_rules {
    cidr_ip = "*"
    rw_mode = "RW"
    use_mode = "All_squash"
  }
  permission_rules {
    cidr_ip = "192.168.0.0"
    rw_mode = "RO"
    use_mode = "All_squash"
  }
}

data "volcengine_nas_permission_groups" "default" {
  filters {
    key = "PermissionGroupId"
    value = volcengine_nas_permission_group.foo.id
  }
}