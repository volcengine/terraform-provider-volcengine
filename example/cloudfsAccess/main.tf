resource "volcengine_cloudfs_access" "foo1" {
  fs_name = "tftest2"

  subnet_id         = "subnet-13fca1crr5d6o3n6nu46cyb5m"
  security_group_id = "sg-rrv1klfg5s00v0x578mx14m"
  vpc_route_enabled = false
}