resource "vestack_scaling_instance_attach" "foo" {
  scaling_group_id = "scg-ybqm0b6kcigh9zu9ce6t"
  instance_ids = ["i-ybrv0nb671l8j1om72u7", "i-ybrv0nb66zl8j1tv0n9p", "i-ybrv0nb670l8j1iwc8dz"]
}