---
subcategory: "VEFAAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vefaas_timer"
sidebar_current: "docs-volcengine-resource-vefaas_timer"
description: |-
  Provides a resource to manage vefaas timer
---
# volcengine_vefaas_timer
Provides a resource to manage vefaas timer
## Example Usage
```hcl
resource "volcengine_vefaas_timer" "foo" {
  function_id = "35ybaxxx"
  name        = "test-1-tf"
  crontab     = "*/10 * * * *"
}
```
## Argument Reference
The following arguments are supported:
* `crontab` - (Required) Set the timing trigger time of the Timer trigger.
* `function_id` - (Required, ForceNew) The ID of Function.
* `name` - (Required, ForceNew) The name of the Timer trigger.
* `description` - (Optional) The description of the Timer trigger.
* `enable_concurrency` - (Optional) Whether the Timer trigger allows concurrency.
* `enabled` - (Optional) Whether the Timer trigger is enabled.
* `payload` - (Optional) The Timer trigger sends the content payload of the request.
* `retries` - (Optional) The retry count of the Timer trigger.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The creation time of the Timer trigger.
* `last_update_time` - The last update time of the Timer trigger.


## Import
VefaasTimer can be imported using the id, e.g.
```
$ terraform import volcengine_vefaas_timer.default FunctionId:Id
```

