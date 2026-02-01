---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_download_tasks"
sidebar_current: "docs-volcengine-datasource-tls_download_tasks"
description: |-
  Use this data source to query detailed information of tls download tasks
---
# volcengine_tls_download_tasks
Use this data source to query detailed information of tls download tasks
## Example Usage
```hcl
data "volcengine_tls_download_tasks" "foo" {
  topic_id  = "8ba48bd7-2493-4300-b1d0-cb760b89e51b"
  task_name = "tf-test-download-task"
}

output "download_tasks" {
  value = data.volcengine_tls_download_tasks.foo.download_tasks
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of download task IDs.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).
* `task_name` - (Optional) The name of the download task.
* `topic_id` - (Optional) The ID of the log topic to which the download tasks belong.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `download_tasks` - The collection of download task results.
    * `allow_incomplete` - Whether to allow incomplete download.
    * `compression` - The compression format of the downloaded file.
    * `create_time` - The create time of the download task.
    * `data_format` - The data format of the downloaded file.
    * `download_url` - The download URL for the completed task.
    * `end_time` - The end time of the log data to download, in Unix timestamp format.
    * `limit` - The maximum number of log entries to download.
    * `log_context_infos` - The info of the log context.
        * `context_flow` - The context flow of the log.
        * `package_offset` - The package offset of the log.
        * `source` - The source of the log.
    * `log_count` - The number of the downloaded logs.
    * `log_size` - The size of the downloaded log data.
    * `query` - The query statement for the download task.
    * `sort` - The sorting order of the log data.
    * `start_time` - The start time of the log data to download, in Unix timestamp format.
    * `task_id` - The ID of the download task.
    * `task_name` - The name of the download task.
    * `task_status` - The status of the download task.
    * `task_type` - The type of the download task.
    * `topic_id` - The ID of the log topic to which the download task belongs.
* `total_count` - The total count of download tasks queried.


