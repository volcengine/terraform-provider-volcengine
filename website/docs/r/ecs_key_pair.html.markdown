---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_key_pair"
sidebar_current: "docs-volcengine-resource-ecs_key_pair"
description: |-
  Provides a resource to manage ecs key pair
---
# volcengine_ecs_key_pair
Provides a resource to manage ecs key pair
## Example Usage
```hcl
resource "volcengine_ecs_key_pair" "default" {
  key_pair_name = "tf-test-key-name1"
  description   = "tftest21111"
}
```
## Argument Reference
The following arguments are supported:
* `key_pair_name` - (Required, ForceNew) The name of key pair.
* `description` - (Optional) The description of key pair.
* `key_file` - (Optional, ForceNew) Target file to save info.
* `public_key` - (Optional, ForceNew) Public key string.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `finger_print` - The finger print info.
* `key_pair_id` - The id of key pair.


## Import
ECS key pair can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_key_pair.default kp-mizl7m1kqccg5smt1bdpijuj
```

