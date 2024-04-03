---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_attach_instance"
sidebar_current: "docs-volcengine-resource-cen_attach_instance"
description: |-
  Provides a resource to manage cen attach instance
---
# volcengine_cen_attach_instance
Provides a resource to manage cen attach instance
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_cen" "foo" {
  cen_name     = "acc-test-cen"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_cen_attach_instance" "foo" {
  cen_id             = volcengine_cen.foo.id
  instance_id        = volcengine_vpc.foo.id
  instance_region_id = "cn-beijing"
  instance_type      = "VPC"
}
```
## Argument Reference
The following arguments are supported:
* `cen_id` - (Required, ForceNew) The ID of the cen.
* `instance_id` - (Required, ForceNew) The ID of the instance.
* `instance_region_id` - (Required, ForceNew) The region ID of the instance.
* `instance_type` - (Required, ForceNew) The type of the instance. Valid values: `VPC`, `DCGW`.
* `instance_owner_id` - (Optional, ForceNew) The owner ID of the instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The create time of the cen attaching instance.
* `status` - The status of the cen attaching instance.


## Import
Cen attach instance can be imported using the CenId:InstanceId:InstanceType:RegionId, e.g.
```
$ terraform import volcengine_cen_attach_instance.default cen-7qthudw0ll6jmc***:vpc-2fexiqjlgjif45oxruvso****:VPC:cn-beijing
```

