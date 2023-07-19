data "volcengine_images" "foo" {
  os_type = "Linux"
  visibility = "public"
  instance_type_id = "ecs.g1.large"
}