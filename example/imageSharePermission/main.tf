resource "volcengine_image" "foo" {
  image_name         = "acc-test-image"
  description        = "acc-test"
  instance_id        = "i-ydi2q1s7wgqc6ild****"
  create_whole_image = false
  project_name       = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_image_share_permission" "foo" {
  image_id   = volcengine_image.foo.id
  account_id = "21000*****"
}
