resource "volcengine_direct_connect_gateway" "foo"{
     direct_connect_gateway_name="tf-test-gateway"
     description="tf-test"
     tags{
          key="k1"
          value="v1"
     }
}