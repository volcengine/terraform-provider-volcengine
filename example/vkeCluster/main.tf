resource "volcengine_vke_cluster" "foo" {
  name = "terraform-test-15"
  description = "created by terraform"
  delete_protection_enabled = false
  cluster_config {
    subnet_ids = ["subnet-rrqvkt2nq1hcv0x57ccqf3x"]
    api_server_public_access_enabled = true
    api_server_public_access_config {
      public_access_network_config {
        billing_type = "PostPaidByBandwidth"
        bandwidth = 1
      }
    }
    resource_public_access_default_enabled = true
  }
  pods_config {
    pod_network_mode = "VpcCniShared"
#    flannel_config {
#      pod_cidrs = ["172.27.224.0/19"]
#      max_pods_per_node = 64
#    }
    vpc_cni_config {
      subnet_ids = ["subnet-rrqvkt2nq1hcv0x57ccqf3x", "subnet-miklcqh75vcw5smt1amo4ik5", "subnet-13g0x0ytpm0hs3n6nu5j591lv"]
    }
  }
  services_config {
    service_cidrsv4 = ["172.30.0.0/18"]
  }
  tags {
    key = "k1"
    value = "v1"
  }
  logging_config {
    //log_project_id = "3189316d-a1ee-4892-a8fc-9a566489d590"
    log_setups {
      enabled = false
      log_ttl = 30
      log_type = "Audit"
    }
  }
}