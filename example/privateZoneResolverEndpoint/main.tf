resource "volcengine_private_zone_resolver_endpoint" "foo" {
  name              = "tf-test"
  vpc_id            = "vpc-13f9uuuqfdjb43n6nu5p1****"
  vpc_region        = "cn-beijing"
  security_group_id = "sg-mj2nsckay29s5smt1b0d****"
  ip_configs {
    az_id     = "cn-beijing-a"
    subnet_id = "subnet-mj2o4co2m2v45smt1bx1****"
    ip        = "172.16.0.2"
  }
  ip_configs {
    az_id     = "cn-beijing-a"
    subnet_id = "subnet-mj2o4co2m2v45smt1bx1****"
    ip        = "172.16.0.3"
  }
  ip_configs {
    az_id     = "cn-beijing-a"
    subnet_id = "subnet-mj2o4co2m2v45smt1bx1****"
    ip        = "172.16.0.4"
  }
  ip_configs {
    az_id     = "cn-beijing-a"
    subnet_id = "subnet-mj2o4co2m2v45smt1bx1****"
    ip        = "172.16.0.5"
  }
}