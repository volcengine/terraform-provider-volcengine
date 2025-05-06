resource "volcengine_dns_record" "foo" {
  zid   = 58846
  host  = "a.com"
  type  = "A"
  value = "1.1.1.2"
}
