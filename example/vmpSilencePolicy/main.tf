resource "volcengine_vmp_silence_policy" "example" {
  name        = "tf-acc-silence"
  description = "terraform silence policy"
  time_range_matchers {
    location = "Asia/Shanghai"
    periodic_date  {
      time         = "20:00~21:12"
      weekday      = "1,5"
    }
  }
  metric_label_matchers {
    matchers {
      label    = "app"
      value    = "test"
      operator = "NotEqual"
    }
    matchers {
      label    = "env"
      value    = "prod"
      operator = "Equal"
    }
  }
}