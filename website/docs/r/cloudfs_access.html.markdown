---
subcategory: "CLOUDFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloudfs_access"
sidebar_current: "docs-volcengine-resource-cloudfs_access"
description: |-
  Provides a resource to manage cloudfs access
---
# volcengine_cloudfs_access
Provides a resource to manage cloudfs access
## Example Usage
```hcl
resource "volcengine_cloudfs_access" "foo1" {
  fs_name = "tftest2"

  subnet_id         = "subnet-13fca1crr5d6o3n6nu46cyb5m"
  security_group_id = "sg-rrv1klfg5s00v0x578mx14m"
  vpc_route_enabled = false
}
```
## Argument Reference
The following arguments are supported:
* `fs_name` - (Required, ForceNew) The name of file system.
* `security_group_id` - (Required, ForceNew) The id of security group.
* `subnet_id` - (Required, ForceNew) The id of subnet.
* `access_account_id` - (Optional, ForceNew) The account id of access.
* `access_iam_role` - (Optional, ForceNew) The iam role of access. If the VPC of another account is attached, the other account needs to create a role with CFSCacheAccess permission, and enter the role name as a parameter.
* `vpc_route_enabled` - (Optional) Whether enable all vpc route.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `access_id` - The id of access.
* `access_service_name` - The service name of access.
* `created_time` - The creation time.
* `is_default` - Whether is default access.
* `status` - Status of access.
* `vpc_id` - The id of vpc.


## Import
CloudFs Access can be imported using the FsName:AccessId, e.g.
```
$ terraform import volcengine_cloudfs_file_system.default tfname:access-**rdgmedx3fow
```

