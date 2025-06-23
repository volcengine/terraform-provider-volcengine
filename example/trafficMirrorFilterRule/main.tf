resource "volcengine_traffic_mirror_filter" "foo" {
  traffic_mirror_filter_name = "acc-test-traffic-mirror-filter"
  description                = "acc-test"
  project_name               = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_traffic_mirror_filter_rule" "foo-ingress" {
  traffic_mirror_filter_id = volcengine_traffic_mirror_filter.foo.id
  traffic_direction        = "ingress"
  description              = "acc-test"
  priority                 = 1
  policy                   = "reject"
  protocol                 = "all"
  source_cidr_block        = "10.0.1.0/24"
  source_port_range        = "-1/-1"
  destination_cidr_block   = "10.0.0.0/24"
  destination_port_range   = "-1/-1"
}

resource "volcengine_traffic_mirror_filter_rule" "foo-egress" {
  traffic_mirror_filter_id = volcengine_traffic_mirror_filter.foo.id
  traffic_direction        = "egress"
  description              = "acc-test"
  priority                 = 2
  policy                   = "reject"
  protocol                 = "tcp"
  source_cidr_block        = "10.0.1.0/24"
  source_port_range        = "80/80"
  destination_cidr_block   = "10.0.0.0/24"
  destination_port_range   = "88/90"
}