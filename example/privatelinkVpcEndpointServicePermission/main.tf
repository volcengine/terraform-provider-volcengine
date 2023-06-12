resource "volcengine_privatelink_vpc_endpoint_service_permission" "foo" {
  service_id        = "epsvc-3rel73uf2ewao5zsk2j2l58ro"
  permit_account_id = "210000000"
}

resource "volcengine_privatelink_vpc_endpoint_service_permission" "foo1" {
  service_id        = "epsvc-3rel73uf2ewao5zsk2j2l58ro"
  permit_account_id = "210000001"
}