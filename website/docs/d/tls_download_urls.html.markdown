---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_download_urls"
sidebar_current: "docs-volcengine-datasource-tls_download_urls"
description: |-
  Use this data source to query detailed information of tls download urls
---
# volcengine_tls_download_urls
Use this data source to query detailed information of tls download urls
## Example Usage
```hcl
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
```
## Argument Reference
The following arguments are supported:
* `task_id` - (Required) The ID of the download task.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `download_url` - The download URL of the download task.


