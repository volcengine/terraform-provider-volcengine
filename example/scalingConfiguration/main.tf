resource "volcengine_scaling_configuration" "foo" {
  scaling_configuration_name = "tf-test"
  scaling_group_id = "scg-ycinx27x25gh9y31p0fy"
  image_id = "image-ycgud4t4hxgso0e27bdl"
  instance_types = ["ecs.g2i.large"]
  instance_name = "tf-test"
  instance_description = ""
  host_name = ""
  password = ""
  key_pair_name = "renhuaxi"
  security_enhancement_strategy = "InActive"
  volumes {
    volume_type = "ESSD_PL0"
    size = 20
    delete_with_instance = false
  }
  volumes {
    volume_type = "ESSD_PL0"
    size = 20
    delete_with_instance = true
  }
  security_group_ids = ["sg-2fepz3c793g1s59gp67y21r34"]
  eip_bandwidth = 10
  eip_isp = "ChinaMobile"
  eip_billing_type = "PostPaidByBandwidth"
  user_data = "IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="
  tags {
    key = "xx"
    value = "xx"
  }
  tags {
    key = "da"
    value = "da"
  }
  project_name = "default"
  hpc_cluster_id = ""
  spot_strategy = "NoSpot"
}

