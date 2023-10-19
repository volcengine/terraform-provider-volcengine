resource "volcengine_iam_policy" "foo" {
  policy_name = "acc-test-policy"
  description = "acc-test"
  policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}

resource "volcengine_iam_user_group" "foo" {
  user_group_name = "acc-test-group"
  description = "acc-test"
  display_name = "acc-test"
}

resource "volcengine_iam_user_group_policy_attachment" "foo" {
  policy_name = volcengine_iam_policy.foo.policy_name
  policy_type = "Custom"
  user_group_name = volcengine_iam_user_group.foo.user_group_name
}

data "volcengine_iam_user_group_policy_attachments" "foo" {
  user_group_name = volcengine_iam_user_group_policy_attachment.foo.user_group_name
}