resource "vestack_vpn_connection" "foo" {
  vpn_connection_name = "tf-test"
  description = "tf-test"
  vpn_gateway_id = "vgw-2d689v3lxs0zk58ozfebct3fc"
  customer_gateway_id = "cgw-2d68c4zglycjk58ozfe96norh"
  local_subnet = ["192.168.0.0/20"]
  remote_subnet = ["192.161.0.0/20"]
  dpd_action = "none"
  nat_traversal = true
  ike_config_psk = "tftest@!3"
  ike_config_version = "ikev1"
  ike_config_mode = "main"
  ike_config_enc_alg = "aes"
  ike_config_auth_alg = "md5"
  ike_config_dh_group = "group1"
  ike_config_lifetime = 200
  ike_config_local_id = "tf_test"
  ike_config_remote_id = "tf_tes"
  ipsec_config_enc_alg = "aes"
  ipsec_config_auth_alg = "sha1"
  ipsec_config_dh_group = "group2"
  ipsec_config_lifetime = 200
}