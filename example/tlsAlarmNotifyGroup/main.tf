resource "volcengine_tls_alarm_notify_group" "foo" {
  iam_project_name = "yyy"
  alarm_notify_group_name = "tf-test"
  notify_type = ["Trigger"]
  receivers {
    receiver_type = "User"
    receiver_names = ["vke-qs"]
    receiver_channels = ["Email", "Sms"]
    start_time = "23:00:00"
    end_time = "23:59:59"
  }
}