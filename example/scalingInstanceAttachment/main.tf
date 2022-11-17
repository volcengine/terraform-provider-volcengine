resource "volcengine_scaling_instance_attachment" "foo" {
  scaling_group_id = "scg-yc03gjtwlcl8j104srzi"
  instance_id = "i-yc03fwlmfym0treofvda"
  delete_type = "Detach"
  entrusted = false
  detach_option = "none"
}