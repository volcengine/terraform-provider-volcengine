resource "volcengine_cloud_identity_user" "foo" {
  user_name    = "acc-test-user-${count.index}"
  display_name = "tf-test-user-${count.index}"
  description  = "tf"
  email        = "88@qq.com"
  phone        = "181"

  count = 2
}

data "volcengine_cloud_identity_users" "foo" {
  user_name = "acc-test-user"
  source    = "Manual"
}
