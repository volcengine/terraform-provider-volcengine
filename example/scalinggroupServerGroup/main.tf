resource "vestack_scalinggroup_server_group" "foo" {
  scaling_group_id = "scg-ybpystn1rqgso04q8wsj"
  server_group_attributes {
    port = 8081
    server_group_id = "rsp-12binhi72jmyo17q7y2jtabud"
    weight = 5
  }
}