resource "volcengine_cen_attach_instance" "foo" {
  cen_id = "cen-12ar8uclj68sg17q7y20v9gil"
  instance_id = "vpc-2fe5dpn0av2m859gp68rhk2dc"
  instance_type = "VPC"
  instance_region_id = "cn-beijing"
}

resource "volcengine_cen_attach_instance" "foo1" {
  cen_id = "cen-12ar8uclj68sg17q7y20v9gil"
  instance_id = "vpc-in66ktl5t24g8gbssz0sqva1"
  instance_type = "VPC"
  instance_region_id = "cn-beijing"
}