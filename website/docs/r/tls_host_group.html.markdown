---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_host_group"
sidebar_current: "docs-volcengine-resource-tls_host_group"
description: |-
  Provides a resource to manage tls host group
---
# volcengine_tls_host_group
Provides a resource to manage tls host group
## Example Usage
```hcl
resource "volcengine_tls_host_group" "foo" {
  host_group_name = "tfgroup"
  host_group_type = "Label"
  host_identifier = "tf-controller"
  auto_update     = false
  service_logging = false
}
```
## Argument Reference
The following arguments are supported:
* `host_group_name` - (Required) The name of host group.
* `host_group_type` - (Required) The type of host group. The value can be IP or Label.
* `auto_update` - (Optional) Whether enable auto update.
* `host_identifier` - (Optional) The identifier of host.
* `host_ip_list` - (Optional) The ip list of host group.
* `iam_project_name` - (Optional) The project name of iam.
* `service_logging` - (Optional) Whether enable service logging.
* `update_end_time` - (Optional) The update end time of log collector.
* `update_start_time` - (Optional) The update start time of log collector.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `abnormal_heartbeat_status_count` - The abnormal heartbeat status count of host.
* `agent_latest_version` - The latest version of log collector.
* `create_time` - The create time of host group.
* `host_count` - The count of host.
* `modify_time` - The modify time of host group.
* `normal_heartbeat_status_count` - The normal heartbeat status count of host.
* `rule_count` - The rule count of host.


## Import
Tls Host Group can be imported using the id, e.g.
```
$ terraform import volcengine_tls_host_group.default edf052s21s*******dc15
```