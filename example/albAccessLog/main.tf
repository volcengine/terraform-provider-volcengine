# Enable ALB Access Log (TOS Bucket)
resource "volcengine_alb_access_log" "default" {
  load_balancer_id = "alb-bdchexlt87pc8dv40nbr6mu7"
  bucket_name      = "tos-buket"
}