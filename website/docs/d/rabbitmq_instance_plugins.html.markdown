---
subcategory: "RABBITMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rabbitmq_instance_plugins"
sidebar_current: "docs-volcengine-datasource-rabbitmq_instance_plugins"
description: |-
  Use this data source to query detailed information of rabbitmq instance plugins
---
# volcengine_rabbitmq_instance_plugins
Use this data source to query detailed information of rabbitmq instance plugins
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

data "volcengine_rabbitmq_instance_plugins" "foo" {
  instance_id = volcengine_rabbitmq_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of rabbitmq instance.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `plugins` - The collection of query.
    * `description` - The description of plugin.
    * `disable_prompt` - The disable prompt of plugin.
    * `enable_prompt` - The enable prompt of plugin.
    * `enabled` - Whether plugin is enabled.
    * `need_reboot_on_change` - Will changing the enabled state of the plugin cause a reboot of the rabbitmq instance.
    * `plugin_name` - The name of plugin.
    * `port` - The port of plugin.
    * `version` - The version of plugin.
* `total_count` - The total count of query.


