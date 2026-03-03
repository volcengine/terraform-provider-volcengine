resource "volcengine_ecs_key_pair" "foo" {
  key_pair_name = "acc-test-key-name"
  description ="acc-test"
  project_name = "default"
  tags {
    key = "tfk1"
    value = "tfv1"
  }
}