resource "volcengine_tls_alarm" "foo" {
  alarm_name = "test"
  project_id = "cc44f8b6-0328-4622-b043-023fca735cd4"
  //status = true
  //trigger_period = 1
  //alarm_period = 10
  alarm_notify_group = ["3019107f-28a2-4208-a2b6-c33fcb97ac3a"]
  user_define_msg = "test for terraform"
  query_request {
    number = 1
    topic_id = "af1a2240-ba62-4f18-b421-bde2f9684e57"
    start_time_offset = -15
    query = "Failed | select count(*) as errNum"
    end_time_offset = 0
  }
  request_cycle {
    type = "Period"
    time = 11
  }
  alarm_period_detail {
    sms = 10
    phone = 10
    email = 2
    general_webhook = 3
  }
  condition = "$1.errNum>0"
}