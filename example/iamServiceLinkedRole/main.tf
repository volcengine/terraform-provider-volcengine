resource "volcengine_iam_service_linked_role" "foo" {
  service_name = "ecs"
  tags {
      key   = "key-2"
      value = "value-3"
  }
}
