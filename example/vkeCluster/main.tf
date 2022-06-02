resource "vestack_vke_cluster" "foo" {
  name = "terraform-test-1"
  description = "created by terraform"
  delete_protection_enabled = false
  cluster_config {
    subnet_ids = ["subnet-3recgzi7hfim85zsk2i8l9ve7"]
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
      pod_cidrs = ["172.29.128.0/18"]
      max_pods_per_node = 64
    }
    vpc_cni_config {
      subnet_ids = ["subnet-2d68w6sl5lf5s58ozfeo83ggo"]
    }
  }
  services_config {
    service_cidrsv4 = ["192.168.248.0/21"]
  }
}