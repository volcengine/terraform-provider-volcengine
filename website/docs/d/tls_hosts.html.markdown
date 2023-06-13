---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_hosts"
sidebar_current: "docs-volcengine-datasource-tls_hosts"
description: |-
  Use this data source to query detailed information of tls hosts
---
# volcengine_tls_hosts
Use this data source to query detailed information of tls hosts
## Example Usage
```hcl
data "volcengine_tls_hosts" "default" {
  host_group_id = "527102e2-1e4f-45f4-a990-751152125da7"
}
```
## Argument Reference
The following arguments are supported:
* `host_group_id` - (Required) The id of host group.
* `heartbeat_status` - (Optional) The the heartbeat status.
* `ip` - (Optional) The ip address.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `host_infos` - The collection of query.
    * `heartbeat_status` - The the heartbeat status.
    * `host_group_id` - The id of host group.
    * `ip` - The ip address.
    * `log_collector_version` - The version of log collector.
* `total_count` - The total count of query.


