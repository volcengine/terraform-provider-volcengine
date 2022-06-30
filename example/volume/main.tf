resource "volcengine_volume" "foo" {
  volume_name = "terraform-test"
  zone_id = "cn-lingqiu-a"
  volume_type = "PTSSD"
  kind = "data"
  size = 40
}