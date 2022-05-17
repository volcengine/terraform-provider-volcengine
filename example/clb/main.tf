resource "vestack_clb" "foo" {
  type = "public"
  subnet_id = "subnet-2744i7u9alnnk7fap8tkq8aft"
  load_balancer_spec = "small_1"
  region_id = "cn-north-3"
  description = "Demo"
}