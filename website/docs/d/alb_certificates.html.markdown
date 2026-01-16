---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_certificates"
sidebar_current: "docs-volcengine-datasource-alb_certificates"
description: |-
  Use this data source to query detailed information of alb certificates
---
# volcengine_alb_certificates
Use this data source to query detailed information of alb certificates
## Example Usage
```hcl
data "volcengine_alb_certificates" "default" {
  certificate_name = "tf-test"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `certificate_name` - (Optional) The Name of Certificate.
* `ids` - (Optional) The list of Certificate IDs.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name to which the certificate belongs.
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
    * `san` - The san extension of the Certificate.
    * `status` - The status of the Certificate.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
* `total_count` - The total count of Certificate query.


