resource "volcengine_escloud_instance" "foo" {
  instance_configuration {
    version            = "V6_7"
    zone_number        = 1
    enable_https       = true
    admin_user_name    = "admin"
    admin_password     = "xxxx"
    charge_type        = "PostPaid"
    configuration_code = "es.standard"
    enable_pure_master = true
    instance_name      = "from-tf4"
    node_specs_assigns {
      type               = "Master"
      number             = 3
      resource_spec_name = "es.x4.medium"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Hot"
      number             = 2
      resource_spec_name = "es.x4.large"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Kibana"
      number             = 1
      resource_spec_name = "kibana.x2.small"
    }
    subnet_id = "subnet-2bz9vxrixqigw2dx0eextz50p"
    project_name = "default"
    force_restart_after_scale = false
    project_name = "default"
  }
}