resource "volcengine_acl" "foo" {
  acl_name = "tf-test-3"
  description = "tf-test"
}

resource "volcengine_acl_entry" "foo" {
  acl_id = volcengine_acl.foo.id
  description = "tf acl entry desc demo"
  entry = "192.2.2.1/32"
}