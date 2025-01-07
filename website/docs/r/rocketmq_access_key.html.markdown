---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_access_key"
sidebar_current: "docs-volcengine-resource-rocketmq_access_key"
description: |-
  Provides a resource to manage rocketmq access key
---
# volcengine_rocketmq_access_key
Provides a resource to manage rocketmq access key
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

resource "volcengine_rocketmq_instance" "foo" {
  zone_ids             = [data.volcengine_zones.foo.zones[0].id]
  subnet_id            = volcengine_subnet.foo.id
  version              = "4.8"
  compute_spec         = "rocketmq.n1.x2.micro"
  storage_space        = 300
  auto_scale_queue     = true
  file_reserved_time   = 10
  instance_name        = "acc-test-rocketmq"
  instance_description = "acc-test"
  project_name         = "default"
  charge_info {
    charge_type = "PostPaid"
  }
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_rocketmq_access_key" "foo" {
  instance_id   = volcengine_rocketmq_instance.foo.id
  description   = "acc-test-key"
  all_authority = "SUB"
}
```
## Argument Reference
The following arguments are supported:
* `all_authority` - (Required) The default authority of the rocketmq topic. Valid values: `ALL`, `PUB`, `SUB`, `DENY`. Default is `DENY`.
* `description` - (Required, ForceNew) The description of the rocketmq topic. The description is used to effectively distinguish and manage keys. Please use different descriptions for each key.
* `instance_id` - (Required, ForceNew) The id of rocketmq instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `access_key` - The access key id of the rocketmq key.
* `acl_config_json` - The acl config of the rocketmq key.
* `actived` - The active status of the rocketmq key.
* `create_time` - The create time of the rocketmq key.
* `secret_key` - The secret key of the rocketmq key.
* `topic_permissions` - The custom authority of the rocketmq key.
    * `permission` - The custom authority for the topic.
    * `topic_name` - The name of the rocketmq topic.


## Import
RocketmqAccessKey can be imported using the instance_id:access_key, e.g.
```
$ terraform import volcengine_rocketmq_access_key.default resource_id
```

