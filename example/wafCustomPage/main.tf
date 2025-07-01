resource "volcengine_waf_custom_page" "foo" {
  host = "www.123.com"
  policy = 1
  client_ip = "ALL"
  name = "tf-test"
  description = "tf-test"
  url = "/tf-test"
  enable = 1
  code = 403
  page_mode = 1
  content_type = "text/html"
  body = "tf-test-body"
  advanced = 1
  redirect_url = "/test/tf/path"
  project_name = "default"
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