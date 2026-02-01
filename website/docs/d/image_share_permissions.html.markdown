---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_image_share_permissions"
sidebar_current: "docs-volcengine-datasource-image_share_permissions"
description: |-
  Use this data source to query detailed information of image share permissions
---
# volcengine_image_share_permissions
Use this data source to query detailed information of image share permissions
## Example Usage
```hcl
data "volcengine_image_share_permissions" "foo" {
  image_id = "image-ydi2wozhozfu03z2****"
}
```
## Argument Reference
The following arguments are supported:
* `image_id` - (Required) The id of the image.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `accounts` - The collection of query.
    * `account_id` - The shared account id of the image.
* `total_count` - The total count of query.


