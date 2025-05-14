resource "volcengine_vefaas_release" "foo" {
  function_id = "f0zvcxxx"
  revision_number = 0
  target_traffic_weight = 100
  lifecycle {
    ignore_changes = [revision_number]
  }
}
