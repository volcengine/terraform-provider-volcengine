---
subcategory: "NAT"
layout: "vestack"
page_title: "Vestack: vestack_snat_entry"
sidebar_current: "docs-vestack-resource-snat_entry"
description: |-
  Provides a resource to manage snat entry
---
# vestack_snat_entry
Provides a resource to manage snat entry
## Example Usage
```hcl
resource "vestack_snat_entry" "foo" {
  nat_gateway_id  = "ngw-2743w1f6iqby87fap8tvm9kop"
  subnet_id       = "subnet-2744i7u9alnnk7fap8tkq8aft"
  eip_id          = "eip-274zlae117nr47fap8tzl24v4"
  snat_entry_name = "tf-test-up"
}
```
## Argument Reference
The following arguments are supported:
* `eip_id` - (Required) The id of the public ip address used by the SNAT entry.
* `nat_gateway_id` - (Required, ForceNew) The id of the nat gateway to which the entry belongs.
* `subnet_id` - (Required, ForceNew) The id of the subnet that is required to access the internet.
* `snat_entry_name` - (Optional) The name of the SNAT entry.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `status` - The status of the SNAT entry.


## Import
Snat entry can be imported using the id, e.g.
```
$ terraform import vestack_snat_entry.default snat-3fvhk47kf56****
```

