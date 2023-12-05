---
subcategory: "BANDWIDTH_PACKAGE"
layout: "volcengine"
page_title: "Volcengine: volcengine_bandwidth_package_attachment"
sidebar_current: "docs-volcengine-resource-bandwidth_package_attachment"
description: |-
  Provides a resource to manage bandwidth package attachment
---
# volcengine_bandwidth_package_attachment
Provides a resource to manage bandwidth package attachment
## Example Usage
```hcl
resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth    = 1
  isp          = "BGP"
  name         = "acc-eip"
  description  = "acc-test"
  project_name = "default"
}

resource "volcengine_bandwidth_package" "foo" {
  bandwidth_package_name = "acc-test"
  billing_type           = "PostPaidByBandwidth"
  isp                    = "BGP"
  description            = "tftest-description"
  bandwidth              = 10
  protocol               = "IPv4"
  tags {
    key   = "tftest"
    value = "tftest"
  }
}

resource "volcengine_bandwidth_package_attachment" "foo" {
  allocation_id        = volcengine_eip_address.foo.id
  bandwidth_package_id = volcengine_bandwidth_package.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `allocation_id` - (Required, ForceNew) The ID of the public IP or IPv6 public bandwidth to be added to the shared bandwidth package instance.
* `bandwidth_package_id` - (Required, ForceNew) The bandwidth package id.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
BandwidthPackageAttachment can be imported using the bandwidth package id and eip id, e.g.
```
$ terraform import volcengine_bandwidth_package_attachment.default BandwidthPackageId:EipId
```

