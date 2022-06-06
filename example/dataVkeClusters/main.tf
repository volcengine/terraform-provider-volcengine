data "vestack_vke_clusters" "default"{
  pods_config_pod_network_mode = "VpcCniShared"
  statuses {
    phase = "Creating"
    conditions_type = "Progressing"
  }
}