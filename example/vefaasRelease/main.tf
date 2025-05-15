resource "volcengine_vefaas_release" "foo" {
  function_id = "9p5emxxxx"
  revision_number = 0
  target_traffic_weight = 30
  lifecycle {
    ignore_changes = [revision_number]
  }
}
