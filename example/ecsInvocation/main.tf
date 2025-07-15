resource "volcengine_ecs_invocation" "foo" {
  command_id             = "cmd-ychkepkhtim0tr3b****"
  instance_ids           = ["i-ychmz92487l8j00o****"]
  invocation_name        = "tf-test"
  invocation_description = "tf"
  username               = "root"
  timeout                = 90
  working_dir            = "/home"
  repeat_mode            = "Rate"
  frequency              = "5m"
  launch_time            = "2023-06-20T09:48:00Z"
  recurrence_end_time    = "2023-06-20T09:59:00Z"
  project_name           = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  parameters {
    name  = "test_str"
    value = "tf"
  }
}