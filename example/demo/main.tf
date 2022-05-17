resource "vestack_demo" "default" {
  name = "vif-name"
  network_interfaces  {
    subnet_id = "vnet-id"
    security_group_ids = ["sg-id-1","sg-id-2"]
  }
}