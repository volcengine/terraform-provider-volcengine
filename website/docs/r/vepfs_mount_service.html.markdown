---
subcategory: "VEPFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vepfs_mount_service"
sidebar_current: "docs-volcengine-resource-vepfs_mount_service"
description: |-
  Provides a resource to manage vepfs mount service
---
# volcengine_vepfs_mount_service
Provides a resource to manage vepfs mount service
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
```
## Argument Reference
The following arguments are supported:
* `mount_service_name` - (Required) The name of the mount service.
* `node_type` - (Required, ForceNew) The node type of the mount service. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `subnet_id` - (Required, ForceNew) The subnet id of the mount service.
* `project` - (Optional, ForceNew) The node type of the mount service.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account id of the mount service.
* `attach_file_systems` - The attached file system info of the mount service.
    * `account_id` - The account id of the vepfs file system.
    * `customer_path` - The id of the vepfs file system.
    * `file_system_id` - The id of the vepfs file system.
    * `file_system_name` - The name of the vepfs file system.
    * `status` - The status of the vepfs file system.
* `create_time` - The created time of the mount service.
* `nodes` - The nodes info of the mount service.
    * `default_password` - The default password of ecs instance.
    * `node_id` - The id of ecs instance.
* `region_id` - The region id of the mount service.
* `status` - The status of the mount service.
* `vpc_id` - The vpc id of the mount service.
* `zone_id` - The zone id of the mount service.
* `zone_name` - The zone name of the mount service.


## Import
VepfsMountService can be imported using the id, e.g.
```
$ terraform import volcengine_vepfs_mount_service.default resource_id
```

