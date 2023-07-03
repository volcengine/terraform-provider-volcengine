resource "volcengine_ecs_launch_template" "foo" {
  launch_template_name = "tf-test-zzm"
  version_description = "testterraformaa"
  instance_type_id = "ecs.g1.large"
  image_id = "image-ycb26d1ryzl8j1fcxa9m"
  instance_name = "instance-test"
  instance_charge_type = "PostPaid"
  host_name = "instance-host-name"
  volumes {
    volume_type = "ESSD_PL0"
    size = 20
    //delete_with_instance = false
  }

  network_interfaces {
    subnet_id = "subnet-3tispp1nai4e8idddd"
    security_group_ids = ["sg-xxxxxxxx"]
  }
  eip_bandwidth = 3
  eip_isp = "BGP"
  eip_billing_type = "PostPaidByBandwidth"
  user_data = "IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="
}