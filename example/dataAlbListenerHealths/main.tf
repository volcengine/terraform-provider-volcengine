data "volcengine_alb_listener_healths" "example" {
  listener_ids = ["lsn-xoetdjk3dzwg54ov5ewpam7c", "lsn-bdcxfof3fy808dv40ofappua"]
  only_un_healthy = true
  project_name = "default"
}