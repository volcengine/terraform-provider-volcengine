# Enable health check log collection
resource "volcengine_alb_health_log" "example" {
  load_balancer_id = "alb-bdchexlt87pc8dv40nbr6mu7"
  topic_id         = "cd507e58-64d2-48e3-9e98-f384430d773a"
  project_id       = "29018d87-858b-4d24-bb8e-5ac958fa5ca5"
}