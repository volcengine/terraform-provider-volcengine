resource "volcengine_iam_role" "foo" {
  role_name             = "tf-test"
  display_name          = "tf-test-modify"
  description           = "tf-test"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"
  max_session_duration  = 4800
  tags {
    key   = "key-1"
    value = "value-1"
  }
}
