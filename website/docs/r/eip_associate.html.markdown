---
subcategory: "EIP"
layout: "volcengine"
page_title: "Volcengine: volcengine_eip_associate"
sidebar_current: "docs-volcengine-resource-eip_associate"
description: |-
  Provides a resource to manage eip associate
---
# volcengine_eip_associate
Provides a resource to manage eip associate
## Example Usage
```hcl
resource "volcengine_eip_associate" "foo" {
  allocation_id = "eip-273ybrd0oeo007fap8t0nggtx"
  instance_id   = "i-cm9tjw9zp9j942mfkczp"
  instance_type = "EcsInstance"
}
```
## Argument Reference
The following arguments are supported:
* `allocation_id` - (Required, ForceNew) The allocation id of the EIP.
* `instance_id` - (Required, ForceNew) The instance id which be associated to the EIP.
* `instance_type` - (Required, ForceNew) The type of the associated instance.
* `private_ip_address` - (Optional, ForceNew) The private IP address of the instance will be associated to the EIP.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Eip associate can be imported using the eip allocation_id:instance_id, e.g.
```
$ terraform import volcengine_eip_associate.default eip-274oj9a8rs9a87fap8sf9515b:i-cm9t9ug9lggu79yr5tcw
```

