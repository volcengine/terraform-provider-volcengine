resource "volcengine_ecs_launch_template" "foo" {
  description = "acc-test-desc"
  eip_bandwidth = 1
  eip_billing_type = "PostPaidByBandwidth"
  eip_isp = "ChinaMobile"
  host_name = "tf-host-name"
  hpc_cluster_id = "hpcCluster-l8u24ovdmoab6opf"
  image_id = "image-ycjwwciuzy5pkh54xx8f"
  instance_charge_type = "PostPaid"
  instance_name = "tf-acc-name"
  instance_type_id = "ecs.g1.large"
  key_pair_name = "tf-key-pair"
  launch_template_name = "tf-acc-template"
}

data "volcengine_ecs_launch_templates" "foo"{
  ids = [volcengine_ecs_launch_template.foo.id]
}