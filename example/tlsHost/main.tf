
resource "volcengine_tls_host_group" "foo" {
  host_group_name   = "tfgroup-ip-tf"
  host_group_type   = "IP"
  host_ip_list      = ["192.168.0.1", "192.168.0.2", "192.168.0.3"]
  auto_update       = true
  update_start_time = "00:00"
  update_end_time   = "02:00"
  service_logging   = false
  iam_project_name  = "default"
}

# 删除指定 IP
resource "volcengine_tls_host" "delete_foo" {
  host_group_id = volcengine_tls_host_group.foo.id
  ip            = "192.168.0.1"
}

# 删除异常机器
resource "volcengine_tls_host" "delete_abnormal" {
  host_group_id = volcengine_tls_host_group.foo.id
}
