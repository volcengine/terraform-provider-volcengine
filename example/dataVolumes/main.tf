data "volcengine_zones" "foo"{
}

resource "volcengine_volume" "foo" {
  volume_name = "acc-test-volume-${count.index}"
  volume_type = "ESSD_PL0"
  description = "acc-test"
  kind = "data"
  size = 60
  zone_id = data.volcengine_zones.foo.zones[0].id
  volume_charge_type = "PostPaid"
  project_name = "default"
  count = 3
}

data "volcengine_volumes" "foo"{
  ids = volcengine_volume.foo[*].id
}
