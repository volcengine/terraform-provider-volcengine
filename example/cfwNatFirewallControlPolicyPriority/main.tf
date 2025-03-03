resource "volcengine_cfw_address_book" "foo" {
  group_name   = "acc-test-address-book"
  description  = "acc-test"
  group_type   = "ip"
  address_list = ["192.168.1.1", "192.168.2.2"]
}

resource "volcengine_cfw_nat_firewall_control_policy" "foo" {
  direction         = "in"
  nat_firewall_id   = "nfw-ydmkayvjsw2vsavx****"
  action            = "accept"
  destination_type  = "group"
  destination       = volcengine_cfw_address_book.foo.id
  proto             = "TCP"
  source_type       = "net"
  source            = "0.0.0.0/0"
  description       = "acc-test-control-policy"
  dest_port_type    = "port"
  dest_port         = "300"
  repeat_type       = "Weekly"
  repeat_start_time = "01:00"
  repeat_end_time   = "11:00"
  repeat_days       = [2, 5, 4]
  start_time        = 1736092800
  end_time          = 1738339140
  priority          = 1
  status            = true
}

resource "volcengine_cfw_nat_firewall_control_policy_priority" "foo" {
  direction       = "in"
  nat_firewall_id = "nfw-ydmkayvjsw2vsavx****"
  rule_id         = volcengine_cfw_nat_firewall_control_policy.foo.rule_id
  new_prio        = 2
}
