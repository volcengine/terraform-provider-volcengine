---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_listeners"
sidebar_current: "docs-volcengine-datasource-alb_listeners"
description: |-
  Use this data source to query detailed information of alb listeners
---
# volcengine_alb_listeners
Use this data source to query detailed information of alb listeners
## Example Usage
```hcl
data "volcengine_alb_listener" "foo" {}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Listener IDs.
* `listener_name` - (Optional) The name of the Listener.
* `load_balancer_id` - (Optional) The id of the Alb.
* `name_regex` - (Optional) A Name Regex of Listener.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the listener.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `listeners` - The collection of Listener query.
    * `acl_ids` - The ID of the access control policy group bound to the listener, only returned when the AclStatus parameter is on.
    * `acl_status` - Whether to enable the access control function,valid value is on or off.
    * `acl_type` - The access control type.
    * `ca_certificate_id` - CA certificate ID associated with HTTPS listener.
    * `certificate_id` - The certificate ID associated with the HTTPS listener.
    * `create_time` - The create time of the Listener.
    * `customized_cfg_id` - The customized configuration ID, the value is empty string when not bound.
    * `description` - The description of listener.
    * `domain_extensions` - The HTTPS listener association list of extension domains for.
        * `certificate_id` - The server certificate ID that domain used.
        * `domain_extension_id` - The extension domain ID.
        * `domain` - The domain.
        * `listener_id` - The listener ID that domain belongs to.
    * `enable_http2` - The HTTP2 feature switch,valid value is on or off.
    * `enable_quic` - The QUIC feature switch,valid value is on or off.
    * `enabled` - The enable status of the Listener.
    * `id` - The ID of the Listener.
    * `listener_id` - The ID of the Listener.
    * `listener_name` - The name of the Listener.
    * `load_balancer_id` - The load balancer ID that the listener belongs to.
    * `port` - The port receiving request of the Listener.
    * `project_name` - The project name of the listener.
    * `protocol` - The protocol of the Listener.
    * `server_group_id` - The ID of the backend server group which is associated with the Listener.
    * `server_groups` - The list of server groups with associated listeners.
        * `server_group_id` - The ID of server group.
        * `server_group_name` - The name of server group.
    * `status` - The status of the Listener.
    * `update_time` - The update time of the Listener.
* `total_count` - The total count of Listener query.


