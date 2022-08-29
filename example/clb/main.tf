resource "volcengine_clb" "foo" {
  type = "public"
  subnet_id = "subnet-273xjcb6wohs07fap8sz3ihhs"
  load_balancer_spec = "small_1"
  description = "Demo"
  load_balancer_name = "terraform-auto-create"
}