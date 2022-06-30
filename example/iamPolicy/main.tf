resource "vestack_iam_policy" "foo" {
  policy_name = "TerraformResourceTest1"
  description = "created by terraform 1"
  policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}