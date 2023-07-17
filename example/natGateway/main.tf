resource "volcengine_nat_gateway" "foo" {
  vpc_id = "vpc-im67wjcikxkw8gbssx8ufpj8"
  subnet_id = "subnet-im67x70vxla88gbssz1hy1z2"
  spec = "Medium"
  nat_gateway_name = "tf-auto-demo-1"
  billing_type = "PostPaid"
  description = "This nat gateway auto-created by terraform. "
  project_name = "default"
}