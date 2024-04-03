---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_attach_instances"
sidebar_current: "docs-volcengine-datasource-cen_attach_instances"
description: |-
  Use this data source to query detailed information of cen attach instances
---
# volcengine_cen_attach_instances
Use this data source to query detailed information of cen attach instances
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

data "volcengine_cen_attach_instances" "foo" {
  cen_id = volcengine_cen_attach_instance.foo.cen_id
}
```
## Argument Reference
The following arguments are supported:
* `cen_id` - (Optional) A cen ID.
* `instance_id` - (Optional) An instance ID.
* `instance_region_id` - (Optional) A region id of instance.
* `instance_type` - (Optional) An instance type.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `attach_instances` - The collection of cen attach instance query.
    * `cen_id` - The ID of the cen.
    * `creation_time` - The create time of the cen attaching instance.
    * `instance_id` - The ID of the instance.
    * `instance_owner_id` - The owner ID of the instance.
    * `instance_region_id` - The region id of the instance.
    * `instance_type` - The type of the instance.
    * `status` - The status of the cen attaching instance.
* `total_count` - The total count of cen attach instance query.


