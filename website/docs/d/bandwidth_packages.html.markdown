---
subcategory: "BANDWIDTH_PACKAGE"
layout: "volcengine"
page_title: "Volcengine: volcengine_bandwidth_packages"
sidebar_current: "docs-volcengine-datasource-bandwidth_packages"
description: |-
  Use this data source to query detailed information of bandwidth packages
---
# volcengine_bandwidth_packages
Use this data source to query detailed information of bandwidth packages
## Example Usage
```hcl
resource "volcengine_bandwidth_package" "foo" {
  bandwidth_package_name    = "acc-test-bp"
  billing_type              = "PostPaidByBandwidth"
  isp                       = "BGP"
  description               = "acc-test"
  bandwidth                 = 2
  protocol                  = "IPv4"
  security_protection_types = ["AntiDDoS_Enhanced"]
  tags {
    key   = "k1"
    value = "v1"
  }
  count = 2
}

data "volcengine_bandwidth_packages" "foo" {
  ids = volcengine_bandwidth_package.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `bandwidth_package_name` - (Optional) Shared bandwidth package name to be queried.
* `ids` - (Optional) Shared bandwidth package instance ID to be queried.
* `isp` - (Optional) Line types for shared bandwidth packages.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the bandwidth package to be queried.
* `protocol` - (Optional) The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
* `security_protection_enabled` - (Optional) Security protection types for shared bandwidth packages.
* `tag_filters` - (Optional) A list of tags.

The `tag_filters` object supports the following:

* `key` - (Required) The key of the tag.
* `values` - (Required) The values of the tag.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `packages` - The collection of query.
    * `bandwidth_package_id` - The id of the bandwidth package.
    * `bandwidth_package_name` - The name of the bandwidth package.
    * `bandwidth` - The bandwidth of the bandwidth package.
    * `billing_type` - The billing type of the bandwidth package.
    * `business_status` - The business status of the bandwidth package.
    * `creation_time` - The creation time of the bandwidth package.
    * `deleted_time` - The deleted time of the bandwidth package.
    * `eip_addresses` - List of public IP information included in the shared bandwidth package.
        * `allocation_id` - The id of the eip.
        * `eip_address` - The eip address.
    * `expired_time` - The expiration time of the bandwidth package.
    * `id` - The id of the bandwidth package.
    * `isp` - The line type.
    * `overdue_time` - The overdue time of the bandwidth package.
    * `project_name` - The project name of the bandwidth package.
    * `protocol` - The protocol of the bandwidth package.
    * `security_protection_types` - Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced`.
    * `status` - The status of the bandwidth package.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of the bandwidth package.
* `total_count` - The total count of query.


