resource "volcengine_iam_role" "foo" {
  role_name             = "tf-test-lh"
  display_name          = "tf-test-lh"
  description           = "tf-test-lh"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"IAM\":[\"trn:iam::2000000001:root\"]}}]}"
  max_session_duration  = 4800
  tags {
    key   = "key-1"
    value = "value-1"
  }
}
