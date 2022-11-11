resource "volcengine_scaling_configuration" "foo" {
  scaling_configuration_name = "tf-test"
  scaling_group_id = "scg-ybru8pazhgl8j1di4tyd"
  image_id = "image-ybpbrfay1gl8j1srwwyz"
  instance_types = ["ecs.g1.4xlarge"]
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
  security_group_ids = ["sg-2ff4fhdtlo8ao59gp67iiq9o3"]
  eip_bandwidth = 0
  eip_isp = "ChinaMobile"
  eip_billing_type = "PostPaidByBandwidth"
  user_data = "IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="
}

