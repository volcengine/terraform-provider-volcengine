resource "volcengine_iam_tag" "foo" {
  resource_type  = "User"
  resource_names = ["jonny"]
  tags {
    key   = "key4"
    value = "value4"
  }
  tags {
      key   = "key3"
      value = "value3"
    }
}
