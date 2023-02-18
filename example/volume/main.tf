resource "volcengine_volume" "foo" {
  volume_name = "terraform-test"
  zone_id = "cn-beijing-b"
  volume_type = "ESSD_PL0"
  kind = "data"
  size = 40
  volume_charge_type = "PrePaid"
}

resource "volcengine_volume_attach" "foo" {
  volume_id = volcengine_volume.foo.id
  instance_id = "i-yc8pfhbafwijutv6s1fv"
}

resource "volcengine_volume" "foo2" {
  volume_name = "terraform-test3"
  zone_id = "cn-beijing-b"
  volume_type = "ESSD_PL0"
  kind = "data"
  size = 40
  volume_charge_type = "PrePaid"
  instance_id = "i-yc8pfhbafwijutv6s1fv"
}