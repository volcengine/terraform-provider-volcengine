---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_all_certificates"
sidebar_current: "docs-volcengine-datasource-alb_all_certificates"
description: |-
  Use this data source to query detailed information of alb all certificates
---
# volcengine_alb_all_certificates
Use this data source to query detailed information of alb all certificates
## Example Usage
```hcl
# Query all certificates (both regular and CA certificates)
data "volcengine_alb_all_certificates" "default" {
  # Optional filters
  ids          = ["cert-1pf4a8k8tokcg845wfariphc2", "cert-xoekc6lpu9s054ov5eohm3bj"]
  project_name = "default"
  tags {
    key   = "key1"
    value = "value2"
  }
}
```
## Argument Reference
The following arguments are supported:
* `certificate_name` - (Optional) The Name of Certificate.
* `certificate_type` - (Optional) The type of Certificate. Valid values: `CA`, `Server`.
* `ids` - (Optional) A list of IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of Certificate.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `certificates` - The collection of Certificate query.
    * `certificate_id` - The ID of the Certificate.
    * `certificate_name` - The name of the Certificate.
    * `certificate_type` - The type of the Certificate.
    * `create_time` - The create time of the Certificate.
    * `description` - The description of the Certificate.
    * `domain_name` - The domain name of the Certificate.
    * `expired_at` - The expire time of the Certificate.
    * `id` - The ID of the Certificate.
    * `listeners` - The ID list of the Listener.
    * `project_name` - The ProjectName of the Certificate.
    * `san` - The list of extended domain names for the certificate, separated by English commas ',', including (commonName, DnsName, IP).
    * `status` - The status of the Certificate.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
* `total_count` - The total count of query.


