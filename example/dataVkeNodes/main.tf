data "vestack_vke_nodes" "default"{
  ids = ["ncaa3e5mrsferqkomi190"]
  cluster_ids = ["c123", "c456"]
  statuses {
     phase = "Creating"
     conditions_type = "Progressing"
  }
  statuses {
       phase = "Creating123"
       conditions_type = "Progressing123"
  }
}