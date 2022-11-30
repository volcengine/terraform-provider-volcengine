// 创建步骤：terraform init -> terraform plan -> terraform apply
// 删除步骤: terraform state rm volcengine_scaling_configuration.foo1 -> terraform destroy

// 创建伸缩组
resource "volcengine_scaling_group" "foo" {
  scaling_group_name = "zzm-tf-test"
  subnet_ids = ["subnet-2fegl9waotzi859gp67relkhv"]
  multi_az_policy = "BALANCE"
  desire_instance_number = 0
  min_instance_number = 0
  max_instance_number = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown = 10
}

// 创建伸缩配置
resource "volcengine_scaling_configuration" "foo1" {
  scaling_configuration_name = "terraform-test"
  scaling_group_id = volcengine_scaling_group.foo.scaling_group_id
  image_id = "image-ybx2d38wdfl8j1pupx7b"
  instance_types = ["ecs.g1.2xlarge"]
  instance_name = "tf-test"
  instance_description = ""
  host_name = ""
  password = ""
  key_pair_name = "zktest"
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
  security_group_ids = ["sg-12b8llnkn1la817q7y1be4kop"]
  eip_bandwidth = 0
  eip_isp = "ChinaMobile"
  eip_billing_type = "PostPaidByBandwidth"
}

// 绑定伸缩配置
resource "volcengine_scaling_configuration_attachment" "foo2" {
  depends_on = [volcengine_scaling_configuration.foo1]
  scaling_configuration_id = volcengine_scaling_configuration.foo1.scaling_configuration_id
}

// 启用伸缩组
resource "volcengine_scaling_group_enabler" "foo3" {
  depends_on = [volcengine_scaling_configuration_attachment.foo2]
  scaling_group_id = volcengine_scaling_group.foo.scaling_group_id
}