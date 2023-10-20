---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_ha_vips"
sidebar_current: "docs-volcengine-datasource-ha_vips"
description: |-
  Use this data source to query detailed information of ha vips
---
# volcengine_ha_vips
Use this data source to query detailed information of ha vips
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

resource "volcengine_ha_vip" "foo" {
  ha_vip_name = "acc-test-ha-vip"
  description = "acc-test"
  subnet_id   = volcengine_subnet.foo.id
  #  ip_address = "172.16.0.5"
}

data "volcengine_ha_vips" "foo" {
  ids = [volcengine_ha_vip.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `ha_vip_name` - (Optional) The name of Ha Vip.
* `ids` - (Optional) A list of Ha Vip IDs.
* `ip_address` - (Optional) The ip address of Ha Vip.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of Ha Vip.
* `status` - (Optional) The status of Ha Vip.
* `subnet_id` - (Optional) The id of subnet.
* `vpc_id` - (Optional) The id of vpc.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ha_vips` - The collection of query.
    * `account_id` - The account id of the Ha Vip.
    * `associated_eip_address` - The associated eip address of the Ha Vip.
    * `associated_eip_id` - The associated eip id of the Ha Vip.
    * `associated_instance_ids` - The associated instance ids of the Ha Vip.
    * `associated_instance_type` - The associated instance type of the Ha Vip.
    * `created_at` - The create time of the Ha Vip.
    * `description` - The description of the Ha Vip.
    * `ha_vip_id` - The id of the Ha Vip.
    * `ha_vip_name` - The name of the Ha Vip.
    * `id` - The id of the Ha Vip.
    * `ip_address` - The ip address of the Ha Vip.
    * `master_instance_id` - The master instance id of the Ha Vip.
    * `project_name` - The project name of the Ha Vip.
    * `status` - The status of the Ha Vip.
    * `subnet_id` - The subnet id of the Ha Vip.
    * `updated_at` - The update time of the Ha Vip.
    * `vpc_id` - The vpc id of the Ha Vip.
* `total_count` - The total count of query.


