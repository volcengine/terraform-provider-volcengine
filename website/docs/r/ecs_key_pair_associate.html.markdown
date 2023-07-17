---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_key_pair_associate"
sidebar_current: "docs-volcengine-resource-ecs_key_pair_associate"
description: |-
  Provides a resource to manage ecs key pair associate
---
# volcengine_ecs_key_pair_associate
Provides a resource to manage ecs key pair associate
## Example Usage
```hcl
resource "volcengine_ecs_key_pair_associate" "default" {
  key_pair_id = "kp-ybvyy1e5msl8u258ovrv"
  instance_id = "i-ybskpw36rul8u1yekckh"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The ID of ECS Instance.
* `key_pair_id` - (Required, ForceNew) The ID of ECS KeyPair Associate.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ECS key pair associate can be imported using the id, e.g.

After binding the key pair, the instance needs to be restarted for the key pair to take effect.

After the key pair is bound, the password login method will automatically become invalid. If your instance has been set for password login, after the key pair is bound, you will no longer be able to use the password login method.

```
$ terraform import volcengine_ecs_key_pair_associate.default kp-ybti5tkpkv2udbfolrft:i-mizl7m1kqccg5smt1bdpijuj
```

