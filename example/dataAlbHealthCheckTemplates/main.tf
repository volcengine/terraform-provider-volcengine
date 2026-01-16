data "volcengine_alb_health_check_templates" "foo" {
  ids = ["hctpl-1iidd1tobnim874adhf708uwf"]
  tags {
    key = "key1"
    value = "value2"
  }
}