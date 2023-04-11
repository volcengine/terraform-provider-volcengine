resource "volcengine_vpn_connection" "foo" {
  vpn_connection_name = "tf-test"
  description = "tf-test"
  vpn_gateway_id = "vgw-2feq19gnyc9hc59gp68914u6o"
  customer_gateway_id = "cgw-12ayj1s157gn417q7y29bixqy"
  local_subnet = ["192.168.0.0/22"]
  remote_subnet = ["192.161.0.0/20"]
  dpd_action = "none"
  nat_traversal = true
  ike_config_psk = "tftest@!3"
  ike_config_version = "ikev1"
  ike_config_mode = "main"
  ike_config_enc_alg = "aes"
  ike_config_auth_alg = "md5"
  ike_config_dh_group = "group2"
  ike_config_lifetime = 100
  ike_config_local_id = "tf_test"
  ike_config_remote_id = "tf_test"
  ipsec_config_enc_alg = "aes"
  ipsec_config_auth_alg = "sha256"
  ipsec_config_dh_group = "group2"
  ipsec_config_lifetime = 100
  project_name = "default"
}