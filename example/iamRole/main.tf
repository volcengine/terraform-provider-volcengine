resource "volcengine_iam_role" "foo" {
  role_name = "acc-test-role"
  display_name = "acc-test"
  description = "acc-test"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"
  max_session_duration = 3600
}
