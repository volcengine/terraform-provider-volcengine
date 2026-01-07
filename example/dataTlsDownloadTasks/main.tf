data "volcengine_tls_download_tasks" "foo" {
  topic_id = "8ba48bd7-2493-4300-b1d0-cb760b89e51b"
  task_name  = "tf-test-download-task"
}

output "download_tasks" {
  value = data.volcengine_tls_download_tasks.foo.download_tasks
}
