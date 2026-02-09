resource "volcengine_tls_host_group" "foo" {
  host_group_name   = "tfgroup-test-x"
  host_group_type   = "Label"
  host_identifier   = "hostlable"
  auto_update       = true
  update_start_time = "00:00"
  update_end_time   = "02:00"
  service_logging   = false
  iam_project_name  = "default"
}

resource "volcengine_tls_host_group" "foo_ip" {
  host_group_name   = "tfgroup-ip-x"
  host_group_type   = "IP"
  host_ip_list      = ["192.168.0.1", "192.168.0.2", "192.168.0.3"]
  auto_update       = true
  update_start_time = "00:00"
  update_end_time   = "02:00"
  service_logging   = false
  iam_project_name  = "default"
}