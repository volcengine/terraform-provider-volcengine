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
  permission_set_id    = volcengine_cloud_identity_permission_set.foo.id
  target_id            = "210005****"
  principal_type       = "User"
  principal_id         = volcengine_cloud_identity_user.foo.id
  deprovision_strategy = "None"
}

# It is not recommended to use this resource to provision the permission_set.
# When the `volcengine_cloud_identity_permission_set` is updated, you can use this resource to provision the permission set.
# When deleting this resource, resource `volcengine_cloud_identity_permission_set_assignment` must be deleted first, and the `deprovision_strategy` of `volcengine_cloud_identity_permission_set_assignment` should be set as `None`.
resource "volcengine_cloud_identity_permission_set_provisioning" "foo" {
  permission_set_id   = volcengine_cloud_identity_permission_set.foo.id
  target_id           = "210005****"
  provisioning_status = "Provisioned"
}
