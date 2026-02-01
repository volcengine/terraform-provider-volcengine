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
data "volcengine_alb_listeners" "foo" {}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Listener IDs.
* `listener_name` - (Optional) The name of the Listener.
* `load_balancer_id` - (Optional) The id of the Alb.
* `name_regex` - (Optional) A Name Regex of Listener.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the listener.
* `protocol` - (Optional) The protocol of the Listener.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `listeners` - The collection of Listener query.
    * `access_log_record_customized_headers_enabled` - Whether the listener has enabled the "Log custom headers in the access log" feature.
    * `acl_ids` - The ID of the access control policy group bound to the listener, only returned when the AclStatus parameter is on.
    * `acl_status` - Whether to enable the access control function,valid value is on or off.
    * `acl_type` - The access control type.
    * `ca_certificate_id` - CA certificate ID associated with HTTPS listener.
    * `ca_certificate_source` - The source of the CA certificate associated with the HTTPS listener.
    * `cert_center_certificate_id` - The certificate id associated with the listener. Source is `cert_center`.
    * `certificate_id` - The certificate ID associated with the HTTPS listener.
    * `certificate_source` - The source of the certificate.
    * `create_time` - The create time of the Listener.
    * `customized_cfg_id` - The customized configuration ID, the value is empty string when not bound.
    * `description` - The description of listener.
    * `domain_extensions` - The HTTPS listener association list of extension domains for.
        * `cert_center_certificate_id` - The server certificate ID used by the domain name. It takes effect when the certificate source is cert_center.
        * `certificate_id` - The server certificate ID that domain used.
        * `certificate_source` - The source of the certificate.
        * `domain_extension_id` - The extension domain ID.
        * `domain` - The domain.
        * `listener_id` - The listener ID that domain belongs to.
        * `pca_leaf_certificate_id` - The server certificate ID used by the domain name. It takes effect when the certificate source is pca_leaf.
        * `san` - The CommonName, extended domain names, and IPs of the certificate are separated by ','.
    * `enable_http2` - The HTTP2 feature switch,valid value is on or off.
    * `enable_quic` - The QUIC feature switch,valid value is on or off.
    * `enabled` - The enable status of the Listener.
    * `id` - The ID of the Listener.
    * `listener_id` - The ID of the Listener.
    * `listener_name` - The name of the Listener.
    * `load_balancer_id` - The load balancer ID that the listener belongs to.
    * `pca_leaf_certificate_id` - The certificate ID associated with the HTTPS listener. Effective when the certificate source is pca_leaf.
    * `pca_root_ca_certificate_id` - The CA certificate ID associated with the HTTPS listener. It takes effect when the certificate source is pca_root.
    * `pca_sub_ca_certificate_id` - The CA certificate ID associated with the HTTPS listener. Effective when the certificate source is pca_sub.
    * `port` - The port receiving request of the Listener.
    * `project_name` - The project name of the listener.
    * `protocol` - The protocol of the Listener.
    * `server_group_id` - The ID of the backend server group which is associated with the Listener.
    * `server_groups` - The list of server groups with associated listeners.
        * `server_group_id` - The ID of server group.
        * `server_group_name` - The name of server group.
    * `status` - The status of the Listener.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of the Listener.
* `total_count` - The total count of Listener query.


