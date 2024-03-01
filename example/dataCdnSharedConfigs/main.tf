data "volcengine_cdn_shared_configs" "foo"{
    config_name = "tf-test"
    config_type = "allow_ip_access_rule"
    project_name = "default"
}