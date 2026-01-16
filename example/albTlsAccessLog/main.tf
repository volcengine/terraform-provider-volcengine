# Enable ALB TLS Access Log (TLS Topic)
resource "volcengine_alb_tls_access_log" "default" {
  load_balancer_id = "alb-bdchexlt87pc8dv40nbr6mu7"
  topic_id         = "a63a5016-3a68-4723-a754-235a09653ce8"
  project_id       = "3746fa99-3eda-42ab-b2c2-a0bf5d6b26ac"
}