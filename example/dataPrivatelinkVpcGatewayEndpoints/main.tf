data "volcengine_vpc_gateway_endpoints" "default" {
  ids = ["gwep-273yuq6q7bgn47fap8squ****"]
}

data "volcengine_vpc_gateway_endpoints" "foo" {
  vpc_id       = "vpc-bp15zkdt37pq72zv****"
  name_regex   = "^acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  output_file = "vpc_gateway_endpoints_output"
}
