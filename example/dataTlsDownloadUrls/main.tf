resource "volcengine_tls_download_task" "foo" {
  topic_id         = "36be6c75-0733-4bee-b63d-48e0eae37f87"
  task_name        = "tf-test-download-mm"
  query            = "*"
  start_time       = 1740426022
  end_time         = 1740626022
  compression      = "gzip"
  data_format      = "json"
  limit            = 10000000
  sort             = "desc"
  allow_incomplete = false
  task_type        = 1
  log_context_infos {
  }
}

output "tls_download_task_id" {
  value = volcengine_tls_download_task.foo.task_id
}

data "volcengine_tls_download_urls" "default" {
  task_id = resource.volcengine_tls_download_task.foo.task_id
}