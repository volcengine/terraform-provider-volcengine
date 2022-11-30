resource "volcengine_scaling_instance_attachment" "foo" {
  scaling_group_id = "scg-yc23rtcea88hcchybf8g"
  instance_id = "i-yc23soxj50gsnz7rxnjp"
  delete_type = "Remove"
  entrusted = true
  detach_option = "none"
}