resource "volcengine_cloud_identity_permission_set" "foo" {
  name             = "acc-test-permission_set-${count.index}"
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

  count = 2
}

data "volcengine_cloud_identity_permission_sets" "foo" {
  ids = volcengine_cloud_identity_permission_set.foo[*].id
}
