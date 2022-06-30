resource "volcengine_server_group" "foo" {
  load_balancer_id = "clb-273z7d4r8tvk07fap8tsniyfe"
  server_group_name = "demo-demo11"
  description = "hello demo11"
}