resource "volcengine_waf_cc_rule" "foo" {
  name = "tf-test"
  url = "/"
  field = "HEADER:User-Agemnt"
  single_threshold = "100"
  path_threshold = 101
  count_time = 102
  cc_type = 1
  effect_time = 200
  rule_priority = 2
  enable = 1
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
  host = "www.tf-test.com"
  exemption_time = 0
  cron_enable = 1
  cron_confs {
    crontab = "* 0 * * 1,2,3,4,5,6,0"
    path_threshold = 123
    single_threshold = 234
  }
  cron_confs {
    crontab = "* 3-8 * * 1,2,3,4,5,6,0"
    path_threshold = 345
    single_threshold = 456
  }
}