resource "volcengine_direct_connect_connection" "foo"{
  direct_connect_connection_name="tf-test-connection"
  description="tf-test"
  direct_connect_access_point_id="ap-cn-beijing-a"
  line_operator="ChinaOther"
  port_type="10GBase"
  port_spec="10G"
  bandwidth=1000
  peer_location="XX路XX号XX楼XX机房"
  customer_name="tf-a"
  customer_contact_phone="12345678911"
  customer_contact_email="email@aaa.com"
  tags{
    key="k1"
    value="v1"
  }
}