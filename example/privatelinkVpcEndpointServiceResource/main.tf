resource "volcengine_privatelink_vpc_endpoint_service_resource" "foo" {
  service_id  = "epsvc-3rel73uf2ewao5zsk2j2l58ro"
  resource_id = "clb-3reii8qfbp7gg5zsk2hsrbe3c"
}

resource "volcengine_privatelink_vpc_endpoint_service_resource" "foo1" {
  service_id  = "epsvc-3rel73uf2ewao5zsk2j2l58ro"
  resource_id = "clb-2d6sfye98rzls58ozfducee1o"
}

resource "volcengine_privatelink_vpc_endpoint_service_resource" "foo2" {
  service_id  = "epsvc-3rel73uf2ewao5zsk2j2l58ro"
  resource_id = "clb-3refkvae02gow5zsk2ilaev5y"
}