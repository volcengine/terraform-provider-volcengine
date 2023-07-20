resource "volcengine_ecs_launch_template" "foo" {
  description = "acc-test-desc"
  eip_bandwidth = 1
  eip_billing_type = "PostPaidByBandwidth"
  eip_isp = "ChinaMobile"
  host_name = "acc-xx"
  hpc_cluster_id = "acc-xx"
  image_id = "acc-xx"
  instance_charge_type = "acc-xx"
  instance_name = "acc-xx"
  instance_type_id = "acc-xx"
  key_pair_name = "acc-xx"
  launch_template_name = "acc-test-template2"
}

data "volcengine_ecs_launch_templates" "foo"{
  ids = [volcengine_ecs_launch_template.foo.id]
}