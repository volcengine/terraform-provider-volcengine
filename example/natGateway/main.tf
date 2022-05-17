resource "vestack_nat_gateway" "foo" {
  vpc_id = "vpc-2740cxyk9im0w7fap8u013dfe"
  subnet_id = "subnet-2740cym8mv9q87fap8u3hfx4i"
  spec = "Medium"
  nat_gateway_name = "tf-auto-demo-1"
  description = "This nat gateway auto-created by terraform. "
}