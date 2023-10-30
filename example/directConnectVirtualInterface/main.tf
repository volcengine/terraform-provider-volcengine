resource "volcengine_direct_connect_virtual_interface" "foo"{
     virtual_interface_name="tf-test-vi"
     description="tf-test"
     direct_connect_connection_id="dcc-rtkzeotzst1cu3numzi****"
     direct_connect_gateway_id="dcg-638x4bjvjawwn3gd5xw****"
     vlan_id=2
     local_ip="**.**.**.**/**"
     peer_ip="**.**.**.**/**"
     route_type="Static"
     enable_bfd=false
     tags{
          key="k1"
          value="v1"
     }
}