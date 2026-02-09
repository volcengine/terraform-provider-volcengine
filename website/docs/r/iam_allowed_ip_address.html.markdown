---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_allowed_ip_address"
sidebar_current: "docs-volcengine-resource-iam_allowed_ip_address"
description: |-
  Provides a resource to manage iam allowed ip address
---
# volcengine_iam_allowed_ip_address
Provides a resource to manage iam allowed ip address
## Example Usage
```hcl
resource "volcengine_iam_allowed_ip_address" "foo" {
  enable_ip_list = true
  ip_list {
    ip          = "your ip"
    description = "test1"
  }
  ip_list {
    ip          = "your ip"
    description = "test2"
  }
}

###执行terraform destroy时
# 设置enable_ip_list=false为删除
# 设置 ip_list = {}
```
## Argument Reference
The following arguments are supported:
* `enable_ip_list` - (Required) Whether to enable the IP whitelist.
* `ip_list` - (Required) The IP whitelist list.

The `ip_list` object supports the following:

* `ip` - (Required) The IP address.
* `description` - (Optional) The description of the IP address.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Iam AllowedIpAddress key don't support import

