
resource "volcengine_nlb_network_load_balancer" "foo" {
  load_balancer_name             = "nlb-test-tf"
  description                    = "nlb test by tf"
  vpc_id                         = "vpc-2d64s88ovqb5s58ozfe3uj5mx"
  security_group_ids             = ["sg-2d64s8elz2h3458ozfes73ytc"]
  network_type                   = "internet"
  region                         = "cn-guilin-boe"
  ip_address_version             = "ipv4"
  cross_zone_enabled             = true
  project_name                   = "default"
  modification_protection_status = "NonProtection"
  modification_protection_reason = "nlb test"
  # ipv4_bandwidth_package_id, ipv6_bandwidth_package_id initialized as an empty string, only for modify
  ipv4_bandwidth_package_id = ""
  ipv6_bandwidth_package_id = ""

  zone_mappings {
    zone_id      = "cn-guilin-a"
    subnet_id    = "subnet-2d64s9m0njojk58ozfeli4bik"
    ipv4_address = "your-ip"
  }

  zone_mappings {
    zone_id      = "cn-guilin-c"
    subnet_id    = "subnet-3rezsw47mbzeo5zsk2h774gwi"
    ipv4_address = "your-ip"
  }

  access_log_config {
   # Initialize the enabled to false, only for modify
    enabled    = false
    project_name = "928df680-0bb8-4403-b51b-166671f1107f"
    topic_name   = "fc4a94bc-70e5-4a42-82bb-8241499c31f5"
  }

  tags {
    key   = "k3"
    value = "v34"
  }
}

provider "volcengine" {
  enable_standard_endpoint = false
}
