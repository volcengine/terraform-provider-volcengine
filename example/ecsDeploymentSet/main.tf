resource "volcengine_ecs_deployment_set" "foo" {
  deployment_set_name = "acc-test-ecs-ds"
  description = "acc-test"
  granularity = "switch"
  strategy = "Availability"
}