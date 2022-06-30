---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_certificates"
sidebar_current: "docs-volcengine-datasource-certificates"
description: |-
  Use this data source to query detailed information of certificates
---
# volcengine_certificates
Use this data source to query detailed information of certificates
## Example Usage
```hcl
data "volcengine_certificates" "default" {
  ids = ["cert-274scdwqufwg07fap8u5fu8pi"]
}
```
## Argument Reference
The following arguments are supported:
* `certificate_name` - (Optional) The name of the Certificate.
* `ids` - (Optional) The list of Certificate IDs.
* `name_regex` - (Optional) The Name Regex of Certificate.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `certificates` - The collection of Certificate query.
  * `certificate_id` - The ID of the Certificate.
  * `certificate_name` - The name of the Certificate.
  * `create_time` - The create time of the Certificate.
  * `description` - The description of the Certificate.
  * `domain_name` - The domain name of the Certificate.
  * `expired_at` - The expire time of the Certificate.
  * `id` - The ID of the Certificate.
  * `listeners` - The ID list of the Listener.
* `total_count` - The total count of Certificate query.


