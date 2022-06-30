---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_listeners"
sidebar_current: "docs-volcengine-datasource-listeners"
description: |-
  Use this data source to query detailed information of listeners
---
# volcengine_listeners
Use this data source to query detailed information of listeners
## Example Usage
```hcl
data "volcengine_listeners" "default" {
  ids = ["lsn-273yv0mhs5xj47fap8sehiiso", "lsn-273yw6zps6pz47fap8swa0q2z"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Listener IDs.
* `listener_name` - (Optional) The name of the Listener.
* `load_balancer_id` - (Optional) The id of the Clb.
* `name_regex` - (Optional) A Name Regex of Listener.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `listeners` - The collection of Listener query.
  * `acl_ids` - The acl ID list to which the Listener is bound.
  * `acl_status` - The acl status of the Listener.
  * `acl_type` - The acl type of the Listener.
  * `certificate_id` - The ID of the certificate which is associated with the Listener.
  * `create_time` - The create time of the Listener.
  * `enabled` - The enable status of the Listener.
  * `health_check_domain` - The domain of health check.
  * `health_check_enabled` - The enable status of health check function.
  * `health_check_healthy_threshold` - The healthy threshold of health check.
  * `health_check_http_code` - The normal http status code of health check.
  * `health_check_interval` - The interval executing health check.
  * `health_check_method` - The method of health check.
  * `health_check_timeout` - The response timeout of health check.
  * `health_check_un_healthy_threshold` - The unhealthy threshold of health check.
  * `health_check_uri` - The uri of health check.
  * `id` - The ID of the Listener.
  * `listener_id` - The ID of the Listener.
  * `listener_name` - The name of the Listener.
  * `port` - The port receiving request of the Listener.
  * `protocol` - The protocol of the Listener.
  * `server_group_id` - The ID of the backend server group which is associated with the Listener.
  * `status` - The status of the Listener.
  * `update_time` - The update time of the Listener.
* `total_count` - The total count of Listener query.


