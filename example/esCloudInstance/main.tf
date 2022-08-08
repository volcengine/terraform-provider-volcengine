
resource "volcengine_escloud_instance" "foo" {
  instance_configuration {
    version            = "V7_10"
    region_id          = "cn-north-4"
    zone_id            = "cn-langfang-a"
    zone_number        = 1
    enable_https       = true
    admin_user_name    = "admin"
    admin_password     = "1qaz!QAZ"
    charge_type        = "PostPaid"
    configuration_code = "es.standard"
    enable_pure_master = false
    instance_name      = "from-tf2"
    node_specs_assigns {
      type               = "Master"
      number             = 3
      resource_spec_name = "es.x4.medium"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Hot"
      number             = 0
      resource_spec_name = "es.x4.medium"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Kibana"
      number             = 1
      resource_spec_name = "kibana.x2.small"
      storage_spec_name  = ""
      storage_size       = 0
    }
    subnet {
      subnet_id = "subnet-1g0d5yqrsxszk8ibuxxzile2l"
      subnet_name = "subnet-1g0d5yqrsxszk8ibuxxzile2l"
    }
    vpc {
      vpc_id= "vpc-3cj17x7u9bzeo6c6rrtzfpaeb"
      vpc_name = "test-1231新建"
    }
  }
}
