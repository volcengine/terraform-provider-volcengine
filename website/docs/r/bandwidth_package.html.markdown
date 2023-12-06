---
subcategory: "BANDWIDTH_PACKAGE"
layout: "volcengine"
page_title: "Volcengine: volcengine_bandwidth_package"
sidebar_current: "docs-volcengine-resource-bandwidth_package"
description: |-
  Provides a resource to manage bandwidth package
---
# volcengine_bandwidth_package
Provides a resource to manage bandwidth package
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
## Example Usage
```hcl
resource "volcengine_bandwidth_package" "foo" {
  bandwidth_package_name = "tf-test"
  billing_type           = "PostPaidByBandwidth"
  //billing_type = "PrePaid"
  isp         = "BGP"
  description = "tftest-description"
  bandwidth   = 10
  protocol    = "IPv4"
  //period = 1
  security_protection_types = ["AntiDDoS_Enhanced"]
  tags {
    key   = "tftest"
    value = "tftest"
  }
}
```
## Argument Reference
The following arguments are supported:
* `bandwidth` - (Required) Bandwidth upper limit of shared bandwidth package, unit: Mbps. Valid values: 2 to 5000.
* `bandwidth_package_name` - (Optional) The name of the bandwidth package.
* `billing_type` - (Optional, ForceNew) BillingType of the Ipv6 bandwidth. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`.
* `description` - (Optional) The description of the bandwidth package.
* `isp` - (Optional, ForceNew) Route type, default to BGP.
* `period` - (Optional, ForceNew) Duration of purchasing shared bandwidth package on an annual or monthly basis. The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.
* `project_name` - (Optional) The project name of the bandwidth package.
* `protocol` - (Optional, ForceNew) The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
* `security_protection_types` - (Optional, ForceNew) Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection, which supports adding basic protection type public IP addresses.If left blank, it indicates a shared bandwidth package with basic protection, which supports the addition of public IP addresses with enhanced protection.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
BandwidthPackage can be imported using the id, e.g.
```
$ terraform import volcengine_bandwidth_package.default bwp-2zeo05qre24nhrqpy****
```

