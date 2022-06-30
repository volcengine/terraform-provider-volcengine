resource "volcengine_server_group_server" "foo" {
  server_group_id = "rsp-273zn4ewlhkw07fap8tig9ujz"
  instance_id = "i-72q1zvko6i5lnawvg940"
  type = "ecs"
  weight = 100
  ip = "192.168.100.99"
  port = 80
  description = "This is a server"
}