---
subcategory: "NAT"
layout: "volcengine"
page_title: "Volcengine: volcengine_snat_entries"
sidebar_current: "docs-volcengine-datasource-snat_entries"
description: |-
  Use this data source to query detailed information of snat entries
---
# volcengine_snat_entries
Use this data source to query detailed information of snat entries
## Example Usage
```hcl
data "volcengine_snat_entries" "default" {
  ids = ["snat-274zl8b1kxzb47fap8u35uune"]
}
```
## Argument Reference
The following arguments are supported:
* `eip_id` - (Optional) An id of the public ip address used by the SNAT entry.
* `ids` - (Optional) A list of SNAT entry ids.
* `nat_gateway_id` - (Optional) An id of the nat gateway to which the entry belongs.
* `output_file` - (Optional) File name where to save data source results.
* `snat_entry_name` - (Optional) A name of SNAT entry.
* `source_cidr` - (Optional) The SourceCidr of SNAT entry.
* `subnet_id` - (Optional) An id of the subnet that is required to access the Internet.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `snat_entries` - The collection of snat entries.
    * `eip_address` - The public ip address used by the SNAT entry.
    * `eip_id` - The id of the public ip address used by the SNAT entry.
    * `id` - The id of the SNAT entry.
    * `nat_gateway_id` - The id of the nat gateway to which the entry belongs.
    * `snat_entry_id` - The id of the SNAT entry.
    * `snat_entry_name` - The name of the SNAT entry.
    * `source_cidr` - The SourceCidr of the SNAT entry.
    * `status` - The status of the SNAT entry.
    * `subnet_id` - The id of the subnet that is required to access the internet.
* `total_count` - The total count of snat entries query.


