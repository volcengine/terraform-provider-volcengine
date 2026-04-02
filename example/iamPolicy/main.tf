resource "volcengine_iam_policy" "foo" {
  policy_name = "acc-test-k"
  description = "acc-modify-k"
  policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"iam:*\"],\"Resource\":[\"*\"]}]}"
}
