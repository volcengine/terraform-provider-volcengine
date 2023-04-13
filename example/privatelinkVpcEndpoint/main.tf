resource "volcengine_privatelink_vpc_endpoint" "endpoint" {
  vpc_id = "vpc-2d5z8cl807y8058ozfds8****"
  security_group_ids = ["sg-2d5z8cr53k45c58ozfdum****"]
  service_id = "epsvc-2byz5nzgiansw2dx0eehh****"
  endpoint_name = "tf-test-ep"
  description = "tf-test"
}

resource "volcengine_privatelink_vpc_endpoint_zone" "zone" {
  endpoint_id = volcengine_privatelink_vpc_endpoint.endpoint.id
  subnet_id = "subnet-2bz47q19zhx4w2dx0eevn****"
  private_ip_address = "172.16.0.252"
}