resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = "cn-beijing-a"
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_vpn_gateway" "foo" {
  vpc_id = volcengine_vpc.foo.id
  subnet_id = volcengine_subnet.foo.id
  bandwidth = 20
  vpn_gateway_name = "acc-test"
  description = "acc-test"
  period = 2
  project_name = "default"
}

resource "volcengine_customer_gateway" "foo" {
  ip_address = "192.0.1.3"
  customer_gateway_name = "acc-test"
  description = "acc-test"
  project_name = "default"
}

resource "volcengine_vpn_connection" "foo" {
  vpn_connection_name = "acc-tf-test"
  description = "acc-tf-test"
  vpn_gateway_id = volcengine_vpn_gateway.foo.id
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
  project_name = "default"
  log_enabled = false
}

resource "volcengine_vpn_gateway_route" "foo" {
  vpn_gateway_id = volcengine_vpn_gateway.foo.id
  destination_cidr_block = "192.168.0.0/20"
  next_hop_id = volcengine_vpn_connection.foo.id
}