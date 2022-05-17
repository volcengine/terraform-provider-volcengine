resource "vestack_security_group_rule" "g1test3" {
  direction = "egress"
  security_group_id = "sg-273ycgql3ig3k7fap8t3dyvqx"
  protocol = "tcp"
  port_start = "8000"
  port_end = "9003"
  cidr_ip = "10.0.0.0/8"
}