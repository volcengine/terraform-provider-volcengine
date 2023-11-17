resource "volcengine_alb_server_group_server" "foo" {
  server_group_id = "rsp-1g7317vrcx3pc2zbhq4c3i6a2"
  instance_id     = "i-ycony2kef4ygp2f8cgmk"
  type            = "ecs"
  weight          = 30
  ip              = "172.16.0.3"
  port            = 5679
  description     = "test add server group server ecs1"
}