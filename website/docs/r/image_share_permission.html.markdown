---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_image_share_permission"
sidebar_current: "docs-volcengine-resource-image_share_permission"
description: |-
  Provides a resource to manage image share permission
---
# volcengine_image_share_permission
Provides a resource to manage image share permission
## Example Usage
```hcl
resource "volcengine_image" "foo" {
  image_name         = "acc-test-image"
  description        = "acc-test"
  instance_id        = "i-ydi2q1s7wgqc6ild****"
  create_whole_image = false
  project_name       = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_image_share_permission" "foo" {
  image_id   = volcengine_image.foo.id
  account_id = "21000*****"
}
```
## Argument Reference
The following arguments are supported:
* `account_id` - (Required, ForceNew) The share account id of the image.
* `image_id` - (Required, ForceNew) The id of the image.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ImageSharePermission can be imported using the image_id:account_id, e.g.
```
$ terraform import volcengine_image_share_permission.default resource_id
```

