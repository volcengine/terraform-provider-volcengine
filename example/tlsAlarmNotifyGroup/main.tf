resource "volcengine_tls_alarm_notify_group" "foo" {
  iam_project_name = "default"
  alarm_notify_group_name = "tf-test-modify-b"
  notify_type = ["Recovery"]
  receivers {
    receiver_type = "User"
    receiver_names = ["jonny"]
    receiver_channels = ["Email"]
    start_time = "23:00:00"
    end_time = "23:59:59"
    general_webhook_url = "https://www.volcengine.com/docs/6470/112220?lang=zh"
    general_webhook_body = "test"
    general_webhook_headers {
      key = "test"
      value = "test"
    }
    general_webhook_method = "POST"
  }

  notice_rules {
    has_next = false
    has_end_node = true
    rule_node {
      type = "Operation"
      value = ["OR"]
      children {
          type = "Condition"
          value = [
                                              "NotifyType",
                                              "in",
                                              "[\"1\"]"
                                          ]
          }
    }
    receiver_infos {
      receiver_type = "User"
      receiver_names = ["jonny"]
      receiver_channels = ["Email", "Sms"]
      start_time = "23:00:00"
      end_time = "23:59:59"
    }
  }
}