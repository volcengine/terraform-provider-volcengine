resource "volcengine_iam_role" "foo" {
  role_name = "TerraformTestRole"
  display_name = "terraform role"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"
  description = "created by terraform"
  max_session_duration = 43200
}