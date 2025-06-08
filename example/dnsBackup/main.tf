resource "volcengine_dns_zone" "foo" {
  zone_name = "xxxx.com"
  tags {
    key   = "xx"
    value = "xx"
  }
  project_name = "default"
  remark       = "xxx"
}

resource "volcengine_dns_backup" "foo" {
  zid = volcengine_dns_zone.foo.id
}
