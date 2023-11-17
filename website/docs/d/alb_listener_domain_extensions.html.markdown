---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_listener_domain_extensions"
sidebar_current: "docs-volcengine-datasource-alb_listener_domain_extensions"
description: |-
  Use this data source to query detailed information of alb listener domain extensions
---
# volcengine_alb_listener_domain_extensions
Use this data source to query detailed information of alb listener domain extensions
## Example Usage
```hcl
data "volcengine_alb_listener_domain_extensions" "foo" {
  listener_id = "lsn-1g72yeyhrrj7k2zbhq5gp6xch"
}
```
## Argument Reference
The following arguments are supported:
* `listener_id` - (Required) A Listener ID.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `domain_extensions` - The collection of domain extensions query.
    * `certificate_id` - The server certificate ID that domain used.
    * `domain_extension_id` - The extension domain ID.
    * `domain` - The domain.
    * `id` - The ID of the Listener.
    * `listener_id` - The listener ID that domain belongs to.
* `total_count` - The total count of Listener query.


