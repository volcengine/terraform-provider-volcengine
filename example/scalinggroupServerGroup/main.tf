resource "vestack_scalinggroup_server_group" "foo" {
  scaling_group_id = "scg-ybru8pazhgl8j1di4tyd"
  server_group_attributes {
    port = 2
    server_group_id = "rsp-12binhi72jmyo17q7y2jtabud"
    weight = 50
  }
  server_group_attributes {
    port = 100
    server_group_id = "rsp-2feq85j3pz4zk59gp67vuz5ss"
    weight = 50
  }
}