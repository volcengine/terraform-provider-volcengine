resource "volcengine_vke_node" "foo" {
  cluster_id = "ccj08tcur544aafnqu450"
  instance_id = "i-ycklnc6atg2udanwod4g"
  keep_instance_name = true
  additional_container_storage_enabled = false
  container_storage_path = ""
  #node_pool_id = "pcj3ifqd5ue15fodk669g"
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