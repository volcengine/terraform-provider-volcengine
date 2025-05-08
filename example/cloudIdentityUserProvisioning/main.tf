resource "volcengine_cloud_identity_user" "foo" {
  user_name    = "acc-test-user"
  display_name = "tf-test-user"
  description  = "tf"
  email        = "88@qq.com"
  phone        = "181"
}

resource "volcengine_cloud_identity_user_provisioning" "foo" {
  principal_type           = "User"
  principal_id             = volcengine_cloud_identity_user.foo.id
  target_id                = "210026****"
  description              = "tf"
  identity_source_strategy = "Ignore"
  duplication_strategy     = "KeepBoth"
  duplication_suffix       = "tf_suffix"
  deletion_strategy        = "Delete"
  policy_name              = ["AdministratorAccess"]
}
