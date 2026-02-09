---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_download_task"
sidebar_current: "docs-volcengine-resource-tls_download_task"
description: |-
  Provides a resource to manage tls download task
---
# volcengine_tls_download_task
Provides a resource to manage tls download task
## Example Usage
```hcl
resource "volcengine_tls_download_task" "foo" {
  topic_id         = "3c57a110-399a-43b3-bc3c-5d60e065239a"
  task_name        = "tf-test-download"
  query            = "*"
  start_time       = 1768448896
  end_time         = 1768450896
  compression      = "gzip"
  data_format      = "json"
  limit            = 1000000
  sort             = "asc"
  allow_incomplete = false
  task_type        = 1
  log_context_infos {
    source         = "your ip"
    context_flow   = "1768450893021#4258909d8fc97e7d-286d6d5f6966623c-6943"
    package_offset = "4833728523"
  }
}

output "tls_download_task_id" {
  value = volcengine_tls_download_task.foo.task_id
}
```
## Argument Reference
The following arguments are supported:
* `end_time` - (Required, ForceNew) The end time of the log data to download, in Unix timestamp format.
* `start_time` - (Required, ForceNew) The start time of the log data to download, in Unix timestamp format.
* `task_name` - (Required, ForceNew) The name of the download task.
* `task_type` - (Required, ForceNew) The type of the download task.
* `topic_id` - (Required, ForceNew) The ID of the log topic to which the download task belongs.
* `allow_incomplete` - (Optional, ForceNew) Whether to allow incomplete download.
* `compression` - (Optional, ForceNew) The compression format of the downloaded file. Valid values: gzip.
* `data_format` - (Optional, ForceNew) The data format of the downloaded file. Valid values: csv, json.
* `limit` - (Optional, ForceNew) The maximum number of log entries to download.
* `log_context_infos` - (Optional, ForceNew) The info of the log context.
* `query` - (Optional, ForceNew) The query statement for the download task.
* `sort` - (Optional, ForceNew) The sorting order of the log data. Valid values: asc, desc.

The `log_context_infos` object supports the following:

* `context_flow` - (Optional, ForceNew) The context flow of the log.
* `package_offset` - (Optional, ForceNew) The package offset of the log.
* `source` - (Optional, ForceNew) The source of the log.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `download_url` - The download URL for the completed task.
* `task_id` - The ID of the download task.
* `task_status` - The status of the download task.


## Import
tls download task can be imported using the topic_id and task_id, e.g.
```
$ terraform import volcengine_tls_download_task.default topic-123456:task-1234567890
```

