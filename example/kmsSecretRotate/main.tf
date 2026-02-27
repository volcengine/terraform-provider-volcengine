resource "volcengine_kms_secret_rotate" "default" {
  secret_name  = "ecs-secret-test"
  version_name = "v1"
}