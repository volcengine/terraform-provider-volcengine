data "volcengine_cfw_vpc_firewall_acl_rules" "foo" {
  vpc_firewall_id = "vfw-ydmjakzksgf7u99j6sby"
  action          = ["accept", "deny"]
}
