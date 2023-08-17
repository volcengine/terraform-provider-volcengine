resource "volcengine_iam_policy" "foo" {
  policy_name = "acc-test-policy"
  description = "acc-test"
  policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}