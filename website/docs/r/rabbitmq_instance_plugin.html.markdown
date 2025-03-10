---
subcategory: "RABBITMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rabbitmq_instance_plugin"
sidebar_current: "docs-volcengine-resource-rabbitmq_instance_plugin"
description: |-
  Provides a resource to manage rabbitmq instance plugin
---
# volcengine_rabbitmq_instance_plugin
Provides a resource to manage rabbitmq instance plugin
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_rabbitmq_instance" "foo" {
  zone_ids             = [data.volcengine_zones.foo.zones[0].id, data.volcengine_zones.foo.zones[1].id, data.volcengine_zones.foo.zones[2].id]
  subnet_id            = volcengine_subnet.foo.id
  version              = "3.8.18"
  user_name            = "acc-test-user"
  user_password        = "93f0cb0614Aab12"
  compute_spec         = "rabbitmq.n3.x2.small"
  storage_space        = 300
  instance_name        = "acc-test-rabbitmq"
  instance_description = "acc-test"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_rabbitmq_instance_plugin" "foo" {
  instance_id = volcengine_rabbitmq_instance.foo.id
  plugin_name = "rabbitmq_shovel"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the rabbitmq instance..
* `plugin_name` - (Required, ForceNew) The name of the plugin.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `description` - The description of the plugin.
* `disable_prompt` - The disable prompt of the plugin.
* `enable_prompt` - The enable prompt of the plugin.
* `enabled` - Whether the plugin is enabled.
* `need_reboot_on_change` - Will changing the enabled state of the plugin cause a reboot of the rabbitmq instance.
* `port` - The port of the plugin.
* `version` - The version of the plugin.


## Import
RabbitmqInstancePlugin can be imported using the instance_id:plugin_name, e.g.
```
$ terraform import volcengine_rabbitmq_instance_plugin.default resource_id
```

