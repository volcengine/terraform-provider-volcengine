resource "volcengine_waf_system_bot" "foo" {
  bot_type = "feed_fetcher"
  project_name = "default"
  action = "observe"
  enable = 0
  host = "www.tf-test.com"
}