resource "volcengine_cen" "foo" {
  cen_name     = "acc-test-cen"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  count = 2
}

data "volcengine_cens" "foo"{
  ids = volcengine_cen.foo[*].id
}
