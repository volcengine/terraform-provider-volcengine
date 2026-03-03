resource "volcengine_vpc_gateway_endpoint" "foo" {
  vpc_id        = "vpc-1elnagq9r6neo1jcpwjx*****"
  service_id    = "gwepsvc-3rfeh9mwev56o5zsk2il*****"
  endpoint_name = "acc-test-gateway-ep"
  description   = "acc-test"
  project_name  = "default"
  vpc_policy    = "{\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":\"*\",\"Action\":\"*\",\"Resource\":\"*\",\"Condition\":null}]}"
  tags {
    key   = "tfk1"
    value = "tfv1"
  }

  tags {
    key   = "tfk2"
    value = "tfv2"
  }
}