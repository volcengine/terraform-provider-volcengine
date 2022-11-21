resource "volcengine_vke_cluster" "foo" {
  name = "terraform-test-15"
  description = "created by terraform"
  delete_protection_enabled = false
  cluster_config {
    subnet_ids = ["subnet-2bzud0pbor8qo2dx0ee884y6h"]
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
    pod_network_mode = "Flannel"
    flannel_config {
      pod_cidrs = ["172.27.224.0/19"]
      max_pods_per_node = 64
    }
    vpc_cni_config {
      subnet_ids = ["subnet-2bzud0pbor8qo2dx0ee884y6h"]
    }
  }
  services_config {
    service_cidrsv4 = ["192.168.0.0/16"]
  }
  tags {
    key = "k1"
    value = "v1"
  }
}