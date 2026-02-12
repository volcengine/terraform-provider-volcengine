resource "volcengine_kms_secret" "foo" {
  secret_name = "tf-test1"
  secret_type = "Generic"
  description = "tf-test"
  secret_value = "{\"dasdasd\":\"dasdasd\"}"
  version_name = "v1.0"
}

resource "volcengine_kms_secret" "foo_ecs" {
  secret_name = "tf-test2"
  version_name = "v2.0"
  secret_type = "ECS"
  description = "tf-test ecs"
  secret_value = "{\"UserName\":\"root\",\"Password\":\"********\"}"
  extended_config = "{\"InstanceId\":\"i-yeehzz2tc0ygp2******\",\"SecretSubType\":\"Password\",\"CustomData\":{\"desc\":\"test\"}}"
  project_name = "default"
  encryption_key = "trn:kms:cn-beijing:21000******:keyrings/Tf-test/keys/Test-key1"
  automatic_rotation = false
  force_delete = false
  pending_window_in_days = 7
}