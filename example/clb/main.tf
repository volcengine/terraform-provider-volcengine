resource "volcengine_clb" "public_clb" {
  type = "public"
  subnet_id = "subnet-mj92ij84m5fk5smt1arvwrtw"
  load_balancer_spec = "small_1"
  description = "Demo"
  load_balancer_name = "terraform-auto-create"
  project_name = "yyy"
  eip_billing_config {
    isp = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth = 1
  }
}

resource "volcengine_clb" "private_clb" {
  type = "private"
  subnet_id = "subnet-mj92ij84m5fk5smt1arvwrtw"
  load_balancer_spec = "small_1"
  description = "Demo"
  load_balancer_name = "terraform-auto-create"
  project_name = "default"
}

resource "volcengine_eip_address" "eip" {
  billing_type = "PostPaidByBandwidth"
  bandwidth = 1
  isp = "BGP"
  name = "tf-eip"
  description = "tf-test"
  project_name = "default"
}

resource "volcengine_eip_associate" "associate" {
  allocation_id = volcengine_eip_address.eip.id
  instance_id = volcengine_clb.private_clb.id
  instance_type = "ClbInstance"
}