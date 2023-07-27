resource "volcengine_vke_node" "foo" {
  cluster_id = "ccj08tcur544aafnqu450"
  instance_id = "i-yck8mmkjxpa8j5x0wfey"
  keep_instance_name = true
  additional_container_storage_enabled = false
  container_storage_path = ""
  kubernetes_config {
    labels {
      key   = "tf-key1"
      value = "tf-value1"
    }
    labels {
      key   = "tf-key2"
      value = "tf-value2"
    }
    taints {
      key = "tf-key3"
      value = "tf-value3"
      effect = "NoSchedule"
    }
    taints {
      key = "tf-key4"
      value = "tf-value4"
      effect = "NoSchedule"
    }
    cordon = true
  }
}