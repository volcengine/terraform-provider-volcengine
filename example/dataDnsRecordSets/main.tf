data "volcengine_dns_zones" "foo" {
  key = "xxx"
  search_mode = "xx"
}

data "volcengine_dns_record_sets" "foo" {
  zid = data.volcengine_dns_zones.foo.zones[0].zid
}