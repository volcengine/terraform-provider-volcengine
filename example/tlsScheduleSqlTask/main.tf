resource "volcengine_tls_schedule_sql_task" "foo" {
  task_name = "tf-test"
  topic_id = "8ba48bd7-2493-4300-b1d0-cb760bxxxxxx"
  dest_region = "cn-beijing"
  dest_topic_id = "b966e41a-d6a6-4999-bd75-39962xxxxxx"
  process_start_time = 1751212980
  process_end_time = 1751295600
  process_time_window = "@m-15m,@m"
  query = "* | SELECT * limit 10000"
  request_cycle {
    cron_tab = "0 10 * * *"
    cron_time_zone = "GMT+08:00"
    time = 1
    type = "CronTab"
  }
  status = 1
  process_sql_delay = 60
  description = "tf-test"
}