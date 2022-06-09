resource "vestack_scaling_configuration" "foo" {
  active = true
  enable = true
  substitute = "scc-ybpz1ky544m0tre9ev3j"
  scaling_configuration_name = "tf-test"
  scaling_group_id = "scg-ybpystn1rqgso04q8wsj"
  image_id = "image-ybpbrfay1gl8j1srwwyz"
  instance_types = ["ecs.g2i.large"]
  instance_name = "tf-test"
  instance_description = "InstanceDescription"
  host_name = "HostName"
  password = ""
  key_pair_name = "tf-test"
  security_enhancement_strategy = "Active"
  volumes {
    volume_type = "ESSD_PL0"
    size = 20
    delete_with_instance = true
  }
  security_group_ids = ["sg-2ff4fhdtlo8ao59gp67iiq9o3"]
  eip_bandwidth = 10
  eip_isp = "BGP"
  eip_billing_type = "PostPaidByTraffic"
  user_data = "IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="
}

