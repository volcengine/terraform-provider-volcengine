resource "volcengine_tls_alarm" "foo" {
  alarm_name = "test-terraform-tf"
  project_id = "88d31abb-62c7-40f5-998e-889747c2a116"
  status = false
  trigger_period = 2
  #alarm_period = 60
  alarm_notify_group = ["bf3ecf26-2081-4e27-ae18-f44dbe5c6138"]
  user_define_msg = "test for terraform"
  
  query_request {
    number = 1
    topic_id = "a690a9b8-72c1-40a3-b8c6-f89a81d3748e"
    start_time_offset = -15
    end_time_offset = 0
    query = "Failed | select count(*) as errNum"
    time_span_type = "Relative"
    truncated_time = "Minute"
    end_time_offset_unit = "Minute"
    start_time_offset_unit = "Minute"
  }
  
  request_cycle {
    type = "Period"
    time = 20
    // cron_tab = "0 18 * * *" # If type is Cron
  }
  
  alarm_period_detail {
    sms = 20
    phone = 20
    email = 20
    general_webhook = 20
  }
  
  # Condition and Severity are ignored if TriggerConditions is used
  # condition = "$1.errNum>0"
  # severity = "critical"

  trigger_conditions {
    condition = "$1.errNum>0"
    severity = "critical"
    count_condition = "__count__ > 0"
    no_data = false
  }


  send_resolved = true
}