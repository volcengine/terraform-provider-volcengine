# endpoint cannot be created,please import by command `terraform import volcengine_cr_endpoint.default endpoint:registryId`

resource "volcengine_cr_endpoint" "default"{
     registry = "tf-1"
     enabled = true
}