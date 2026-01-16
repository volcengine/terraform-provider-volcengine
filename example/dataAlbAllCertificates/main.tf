# Query all certificates (both regular and CA certificates)
data "volcengine_alb_all_certificates" "default" {
  # Optional filters
  ids = ["cert-1pf4a8k8tokcg845wfariphc2", "cert-xoekc6lpu9s054ov5eohm3bj"]
  project_name = "default"
  tags {
    key    = "key1"
    value = "value2"
  }
}