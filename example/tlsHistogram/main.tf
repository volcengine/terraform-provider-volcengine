provider "volcengine" {
  region = "cn-beijing"
}

resource "volcengine_tls_project" "foo" {
  project_name = "tf-test-project"
  description  = "tf-test-project"
}

resource "volcengine_tls_topic" "foo" {
  project_id      = volcengine_tls_project.foo.id
  topic_name      = "tf-test-topic"
  partition_count = 2
  shard_count     = 2
  ttl             = 30
}

resource "volcengine_tls_index" "foo" {
  topic_id = volcengine_tls_topic.foo.id
  full_text {
    case_sensitive = true
    include_chinese = false
    delimiter      = ","
  }
}

data "volcengine_tls_histograms" "foo" {
  topic_id   = volcengine_tls_topic.foo.id
  start_time = 1716307200000
  end_time   = 1716393600000
  query      = "*"
  interval   = 3600000
  depends_on = [volcengine_tls_index.foo]
}

output "histogram_total_count" {
  value = data.volcengine_tls_histograms.foo.total_count
}

output "histogram_infos" {
  value = data.volcengine_tls_histograms.foo.histogram_infos
}
