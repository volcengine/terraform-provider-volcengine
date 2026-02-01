---
subcategory: "VEPFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vepfs_mount_services"
sidebar_current: "docs-volcengine-datasource-vepfs_mount_services"
description: |-
  Use this data source to query detailed information of vepfs mount services
---
# volcengine_vepfs_mount_services
Use this data source to query detailed information of vepfs mount services
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = "cn-beijing-a"
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_vepfs_mount_service" "foo" {
  mount_service_name = "acc-test-mount-service"
  subnet_id          = volcengine_subnet.foo.id
  node_type          = "ecs.g1ie.large"
  project            = "default"
}

data "volcengine_vepfs_mount_services" "foo" {
  mount_service_id = volcengine_vepfs_mount_service.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `file_system_id` - (Optional) The id of Vepfs File System.
* `mount_service_id` - (Optional) The id of mount service.
* `mount_service_name` - (Optional) The name of mount service. This field support fuzzy query.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `status` - (Optional) The query status list of mount service.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `mount_services` - The collection of query.
    * `account_id` - The account id of the mount service.
    * `attach_file_systems` - The attached file system info of the mount service.
        * `account_id` - The account id of the vepfs file system.
        * `customer_path` - The id of the vepfs file system.
        * `file_system_id` - The id of the vepfs file system.
        * `file_system_name` - The name of the vepfs file system.
        * `status` - The status of the vepfs file system.
    * `create_time` - The created time of the mount service.
    * `id` - The id of the mount service.
    * `mount_service_id` - The id of the mount service.
    * `mount_service_name` - The name of the mount service.
    * `nodes` - The nodes info of the mount service.
        * `default_password` - The default password of ecs instance.
        * `node_id` - The id of ecs instance.
    * `project` - The project of the mount service.
    * `region_id` - The region id of the mount service.
    * `status` - The status of the mount service.
    * `subnet_id` - The subnet id of the mount service.
    * `vpc_id` - The vpc id of the mount service.
    * `zone_id` - The zone id of the mount service.
    * `zone_name` - The zone name of the mount service.
* `total_count` - The total count of query.


