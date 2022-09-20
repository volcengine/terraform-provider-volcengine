resource "volcengine_ecs_deployment_set" "default" {
  deployment_set_name = "tf-test"
  description ="test1"
  granularity = "host"
}