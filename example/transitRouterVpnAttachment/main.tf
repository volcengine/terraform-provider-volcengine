resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
}

resource "volcengine_customer_gateway" "foo" {
  ip_address = "192.0.1.3"
  customer_gateway_name = "acc-test"
  description = "acc-test"
}

resource "volcengine_vpn_connection" "foo" {
  vpn_connection_name = "acc-tf-test"
  description = "acc-tf-test"
  attach_type = "TransitRouter"
  customer_gateway_id = volcengine_customer_gateway.foo.id
  local_subnet = ["192.168.0.0/22"]
  remote_subnet = ["192.161.0.0/20"]
  dpd_action = "none"
  nat_traversal = true
  ike_config_psk = "acctest@!3"
  ike_config_version = "ikev1"
  ike_config_mode = "main"
  ike_config_enc_alg = "aes"
  ike_config_auth_alg = "md5"
  ike_config_dh_group = "group2"
  ike_config_lifetime = 9000
  ike_config_local_id = "acc_test"
  ike_config_remote_id = "acc_test"
  ipsec_config_enc_alg = "aes"
  ipsec_config_auth_alg = "sha256"
  ipsec_config_dh_group = "group2"
  ipsec_config_lifetime = 9000
  log_enabled = false
}

resource "volcengine_transit_router_vpn_attachment" "foo" {
  zone_id = "cn-beijing-a"
  transit_router_attachment_name = "tf-test-acc"
  description = "tf-test-acc-desc"
  transit_router_id = volcengine_transit_router.foo.id
  vpn_connection_id = volcengine_vpn_connection.foo.id
  tags {
    key = "k1"
    value = "v1"
  }
}
