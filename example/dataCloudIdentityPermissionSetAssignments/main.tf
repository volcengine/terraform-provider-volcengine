resource "volcengine_cloud_identity_permission_set" "foo" {
  name             = "acc-test-permission_set"
  description      = "tf"
  session_duration = 5000
  permission_policies {
    permission_policy_type = "System"
    permission_policy_name = "AdministratorAccess"
    inline_policy_document = ""
  }
  permission_policies {
    permission_policy_type = "System"
    permission_policy_name = "ReadOnlyAccess"
    inline_policy_document = ""
  }
  permission_policies {
    permission_policy_type = "Inline"
    inline_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
  }
}

resource "volcengine_cloud_identity_user" "foo" {
  user_name    = "acc-test-user"
  display_name = "tf-test-user"
  description  = "tf"
  email        = "88@qq.com"
  phone        = "181"
}

resource "volcengine_cloud_identity_permission_set_assignment" "foo" {
  permission_set_id = volcengine_cloud_identity_permission_set.foo.id
  target_id         = "210026****"
  principal_type    = "User"
  principal_id      = volcengine_cloud_identity_user.foo.id
}

data "volcengine_cloud_identity_permission_set_assignments" "foo" {
  permission_set_id = volcengine_cloud_identity_permission_set_assignment.foo.permission_set_id
}
