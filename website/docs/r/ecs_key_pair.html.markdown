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
resource "volcengine_ecs_key_pair" "foo" {
  key_pair_name = "acc-test-key-name"
  description   = "acc-test"
  project_name  = "default"
  tags {
    key   = "tfk1"
    value = "tfv1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `key_pair_name` - (Required, ForceNew) The name of key pair.
* `description` - (Optional) The description of key pair.
* `key_file` - (Optional, ForceNew) Target file to save private key. It is recommended that the value not be empty. You only have one chance to download the private key, the volcengine will not save your private key, please keep it safe. In the TF import scenario, this field will not write the private key locally.
* `project_name` - (Optional) The project name of the key pair.
* `public_key` - (Optional, ForceNew) Public key string.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

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

