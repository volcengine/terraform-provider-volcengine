data "volcengine_private_zones" "foo" {
  zid            = 770000
  zone_name      = "volces.com"
  search_mode    = "EXACT"
  recursion_mode = true
  line_mode      = 3
}
