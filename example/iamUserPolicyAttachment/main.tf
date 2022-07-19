resource "volcengine_iam_user" "user" {
  user_name = "TfTest"
  description = "test"
}

resource "volcengine_iam_policy" "policy" {
  policy_name = "TerraformResourceTest1"
  description = "created by terraform 1"
  policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}

resource "volcengine_iam_user_policy_attachment" "foo" {
  user_name = volcengine_iam_user.user.user_name
  policy_name = volcengine_iam_policy.policy.policy_name
  policy_type = volcengine_iam_policy.policy.policy_type
}