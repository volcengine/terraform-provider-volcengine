data "volcengine_tls_download_tasks" "foo" {
  topic_id = "3c57a110-399a-43b3-bc3c-5d60e065239a"
  task_name  = "tf-test-download"
}

output "download_tasks" {
  value = data.volcengine_tls_download_tasks.foo.download_tasks
}
