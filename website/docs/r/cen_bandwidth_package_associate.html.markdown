---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_bandwidth_package_associate"
sidebar_current: "docs-volcengine-resource-cen_bandwidth_package_associate"
description: |-
  Provides a resource to manage cen bandwidth package associate
---
# volcengine_cen_bandwidth_package_associate
Provides a resource to manage cen bandwidth package associate
## Example Usage
```hcl
resource "volcengine_cen_bandwidth_package_associate" "foo" {
  cen_bandwidth_package_id = "cbp-2bzeew3s8p79c2dx0eeohej4x"
  cen_id                   = "cen-2bzrl3srxsv0g2dx0efyoojn3"
}
```
## Argument Reference
The following arguments are supported:
* `cen_bandwidth_package_id` - (Required, ForceNew) The ID of the cen bandwidth package.
* `cen_id` - (Required, ForceNew) The ID of the cen.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Cen bandwidth package associate can be imported using the CenBandwidthPackageId:CenId, e.g.
```
$ terraform import volcengine_cen_bandwidth_package_associate.default cbp-4c2zaavbvh5fx****:cen-7qthudw0ll6jmc****
```

