resource "volcengine_vke_node" "foo" {
  cluster_id = "ccj08tcur544aafnqu450"
  instance_id = "i-yck8mmkjxpa8j5x0wfey"
  keep_instance_name = true
  additional_container_storage_enabled = false
  container_storage_path = ""
  kubernetes_config {
    labels {
      key   = "ni"
      value = "bb"
    }
    labels {
      key   = "cccc"
      value = "dddd"
    }
    taints {
      key = "nini"
      value = "dddd"
      effect = "NoSchedule"
    }
    taints {
      key = "ninininini"
      value = "111"
      effect = "NoSchedule"
    }
    cordon = true
  }
}