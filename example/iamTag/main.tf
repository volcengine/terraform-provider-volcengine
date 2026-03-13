resource "volcengine_iam_tag" "foo" {
  resource_type  = "User"
  resource_names = ["jonny"]
  tags {
    key   = "key1"
    value = "value1"
  }
  tags {
      key   = "key2"
      value = "value2"
    }
}
