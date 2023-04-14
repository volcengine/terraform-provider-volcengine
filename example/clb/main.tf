resource "volcengine_clb" "foo" {
  type = "public"
  subnet_id = "subnet-mj92ij84m5fk5smt1arvwrtw"
  load_balancer_spec = "small_1"
  description = "Demo"
  load_balancer_name = "terraform-auto-create"
  project_name = "yyy"
}