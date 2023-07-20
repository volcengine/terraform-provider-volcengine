resource "volcengine_ecs_key_pair" "foo" {
  key_pair_name = "acc-test-key-name"
  description ="acc-test"
}
data "volcengine_ecs_key_pairs" "foo"{
  key_pair_name = volcengine_ecs_key_pair.foo.key_pair_name
}