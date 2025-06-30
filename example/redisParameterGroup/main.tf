resource "volcengine_redis_parameter_group" "foo" {
  name = "tf-test"
  engine_version = "5.0"
  description = "tf-test-description"
  param_values {
    name = "active-defrag-cycle-max"
    value = "30"
  }
  param_values {
    name = "active-defrag-cycle-min"
    value = "15"
  }
}