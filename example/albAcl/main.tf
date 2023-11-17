resource "volcengine_alb_acl" "foo" {
  acl_name    = "tf-test-1"
  description = "tftest"
  acl_entries {
    entry       = "172.20.1.0/24"
    description = "e1"
  }

  acl_entries {
    entry       = "172.20.3.0/24"
    description = "e2"
  }
}