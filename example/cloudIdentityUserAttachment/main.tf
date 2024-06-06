resource "volcengine_cloud_identity_group" "foo" {
  group_name   = "acc-test-group"
  display_name = "tf-test-group"
  join_type    = "Manual"
  description  = "tf"
}

resource "volcengine_cloud_identity_user" "foo" {
  user_name    = "acc-test-user"
  display_name = "tf-test-user"
  description  = "tf"
  email        = "88@qq.com"
  phone        = "181"
}

resource "volcengine_cloud_identity_user_attachment" "foo" {
  user_id  = volcengine_cloud_identity_user.foo.id
  group_id = volcengine_cloud_identity_group.foo.id
}
