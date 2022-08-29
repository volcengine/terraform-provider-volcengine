resource "volcengine_server_group_server" "foo" {
  server_group_id = "rsp-274xltv2sjoxs7fap8tlv3q3s"
  instance_id = "i-ybp1scasbe72q1vq35wv"
  type = "ecs"
  weight = 100
  port = 80
  description = "This is a server"
}