data "volcengine_clb_rules" "default"{
  listener_id = "lsn-273ywvnmiu70g7fap8u2xzg9d"
  ids = ["rule-273z9jo9v3mrk7fap8sq8v5x7"]
}

output "data" {
  value = data.volcengine_clb_rules.default
}