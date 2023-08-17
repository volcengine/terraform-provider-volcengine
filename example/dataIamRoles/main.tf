resource "volcengine_iam_role" "foo1" {
  role_name = "acc-test-role1"
  display_name = "acc-test1"
  description = "acc-test1"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"
  max_session_duration = 3600
}

resource "volcengine_iam_role" "foo2" {
  role_name = "acc-test-role2"
  display_name = "acc-test2"
  description = "acc-test2"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"ecs\"]}}]}"
  max_session_duration = 3600
}

data "volcengine_iam_roles" "foo"{
  role_name = "${volcengine_iam_role.foo1.role_name},${volcengine_iam_role.foo2.role_name}"
}
