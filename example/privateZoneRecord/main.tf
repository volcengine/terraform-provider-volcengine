resource "volcengine_private_zone_record" "foo" {
  zid    = 245****
  host   = "www"
  type   = "A"
  value  = "10.1.1.158"
  weight = 8
  ttl    = 700
  remark = "tf-test"
  enable = true
}

data "volcengine_private_zone_record_sets" "foo" {
  zid         = volcengine_private_zone_record.foo.zid
  host        = volcengine_private_zone_record.foo.host
  search_mode = "EXACT"
}

resource "volcengine_private_zone_record_weight_enabler" "foo" {
  zid            = volcengine_private_zone_record.foo.zid
  record_set_id  = [for set in data.volcengine_private_zone_record_sets.foo.record_sets : set.record_set_id if set.type == volcengine_private_zone_record.foo.type][0]
  weight_enabled = true
}
