resource "volcengine_iam_policy" "foo" {
  policy_name = "acc-test"
  description = "acc-modify"
  policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}
