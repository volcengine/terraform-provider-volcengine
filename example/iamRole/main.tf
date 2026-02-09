resource "volcengine_iam_role" "foo" {
  role_name             = "tf-test-role"
  display_name          = "tf-test-modify"
  description           = "tf-test-modify"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"
  max_session_duration  = 3600
  tags {
    key   = "key-modify"
    value = "value-modify"
  }
}
