resource "volcengine_waf_bot_analyse_protect_rule" "foo" {
  project_name = "default"
  statistical_type = 2
  statistical_duration = 50
  single_threshold = 100
  single_proportion = 0.25
  rule_priority = 3
  path_threshold = 1000
  path = "/mod"
  name = "tf-test-mod"
  host = "www.tf-test.com"
  field = "HEADER:User-Agent"
  exemption_time = 60
  enable = 1
  effect_time = 1000
  action_type = 1
  action_after_verification = 1
  accurate_group {
    accurate_rules {
      http_obj = "request.uri"
      obj_type = 1
      opretar = 2
      property = 0
      value_string = "tf"
    }
    accurate_rules {
      http_obj = "request.schema"
      obj_type = 0
      opretar = 2
      property = 0
      value_string = "tf-2"
    }
    logic = 2
  }
}
