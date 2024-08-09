---
subcategory: "PRIVATE_ZONE"
layout: "volcengine"
page_title: "Volcengine: volcengine_private_zone_user_vpc_authorization"
sidebar_current: "docs-volcengine-resource-private_zone_user_vpc_authorization"
description: |-
  Provides a resource to manage private zone user vpc authorization
---
# volcengine_private_zone_user_vpc_authorization
Provides a resource to manage private zone user vpc authorization
## Example Usage
```hcl
resource "volcengine_private_zone_user_vpc_authorization" "foo" {
  account_id = "2100278462"
}
```
## Argument Reference
The following arguments are supported:
* `account_id` - (Required, ForceNew) The account Id which authorizes the private zone resource.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
PrivateZoneUserVpcAuthorization can be imported using the id, e.g.
```
$ terraform import volcengine_private_zone_user_vpc_authorization.default resource_id
```

