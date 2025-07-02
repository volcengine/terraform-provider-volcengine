resource "volcengine_waf_custom_bot" "foo" {
  host = "www.tf-test.com"
  bot_type = "tf-test"
  description = "tf-test"
  project_name = "default"
  action = "observe"
  enable = 1
  accurate {
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