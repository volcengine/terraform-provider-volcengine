resource "volcengine_cloudfs_file_system" "foo" {
  fs_name    = "tffile"
  zone_id    = "cn-beijing-b"
  cache_plan = "T2"
  mode       = "HDFS_MODE"
  read_only  = true

  subnet_id          = "subnet-13fca1crr5d6o3n6nu46cyb5m"
  security_group_id  = "sg-rrv1klfg5s00v0x578mx14m"
  cache_capacity_tib = 10
  vpc_route_enabled  = true

  tos_bucket = "tfacc"
  tos_prefix = "pre/"
}


resource "volcengine_cloudfs_file_system" "foo1" {
  fs_name    = "tffileu"
  zone_id    = "cn-beijing-b"
  cache_plan = "T2"
  mode       = "ACC_MODE"
  read_only  = true

  subnet_id          = "subnet-13fca1crr5d6o3n6nu46cyb5m"
  security_group_id  = "sg-rrv1klfg5s00v0x578mx14m"
  cache_capacity_tib = 15
  vpc_route_enabled  = false

  tos_bucket = "tfacc"
}