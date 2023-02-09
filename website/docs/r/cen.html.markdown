---
subcategory: "CEN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen"
sidebar_current: "docs-volcengine-resource-cen"
description: |-
  Provides a resource to manage cen
---
# volcengine_cen
Provides a resource to manage cen
## Example Usage
```hcl
resource "volcengine_cen" "foo" {
  cen_name    = "tf-test"
  description = "tf-test"
}
```
## Argument Reference
The following arguments are supported:
* `cen_name` - (Optional) The name of the cen.
* `description` - (Optional) The description of the cen.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account ID of the cen.
* `cen_bandwidth_package_ids` - A list of bandwidth package IDs of the cen.
* `cen_id` - The ID of the cen.
* `creation_time` - The create time of the cen.
* `status` - The status of the cen.
* `update_time` - The update time of the cen.


## Import
Cen can be imported using the id, e.g.
```
$ terraform import volcengine_cen.default cen-7qthudw0ll6jmc****
```

