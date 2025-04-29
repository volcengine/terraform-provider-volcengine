resource "volcengine_cr_registry" "foo" {
  name               = "acc-test-cr"
  delete_immediately = false
  password           = "1qaz!QAZ"
  project            = "default"
}