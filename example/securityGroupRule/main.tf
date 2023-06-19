resource "volcengine_security_group_rule" "g1test3" {
  direction         = "egress"
  security_group_id = "sg-2d6722jpp55og58ozfd1sqtdb"
  protocol          = "tcp"
  port_start        = 8000
  port_end          = 9003
  cidr_ip           = "10.0.0.0/8"
  description       = "tft1234"
}

resource "volcengine_security_group_rule" "g1test2" {
  direction         = "egress"
  security_group_id = "sg-2d6722jpp55og58ozfd1sqtdb"
  protocol          = "tcp"
  port_start        = 8000
  port_end          = 9003
  cidr_ip           = "10.0.0.0/24"
}

resource "volcengine_security_group_rule" "g1test1" {
  direction         = "egress"
  security_group_id = "sg-2d6722jpp55og58ozfd1sqtdb"
  protocol          = "tcp"
  port_start        = 8000
  port_end          = 9003
  cidr_ip           = "10.0.0.0/24"
  priority          = 2
}


resource "volcengine_security_group_rule" "g1test0" {
  direction         = "ingress"
  security_group_id = "sg-2d6722jpp55og58ozfd1sqtdb"
  protocol          = "tcp"
  port_start        = 80
  port_end          = 80
  cidr_ip           = "10.0.0.0/24"
  priority          = 2
  policy            = "drop"
  description       = "tft"
}

resource "volcengine_security_group_rule" "g1test06" {
  direction         = "ingress"
  security_group_id = "sg-2d6722jpp55og58ozfd1sqtdb"
  protocol          = "tcp"
  port_start        = 8000
  port_end          = 9003
  source_group_id   = "sg-3rfe5j4xdnklc5zsk2hcw5c6q"
  priority          = 2
  policy            = "drop"
  description       = "tft"
}