---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_allowlist"
sidebar_current: "docs-volcengine-resource-rds_postgresql_allowlist"
description: |-
  Provides a resource to manage rds postgresql allowlist
---
# volcengine_rds_postgresql_allowlist
Provides a resource to manage rds postgresql allowlist
## Example Usage
```hcl
resource "volcengine_rds_postgresql_allowlist" "foo" {
  allow_list_name = "acc-test-allowlist"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = ["192.168.0.0/24", "192.168.1.0/24"]
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_name` - (Required) The name of the postgresql allow list.
* `allow_list` - (Required) Enter an IP address or a range of IP addresses in CIDR format.
* `allow_list_desc` - (Optional) The description of the postgresql allow list.
* `allow_list_type` - (Optional, ForceNew) The type of IP address in the whitelist. Currently only `IPv4` addresses are supported.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `associated_instance_num` - The total number of instances bound under the whitelist.
* `associated_instances` - The list of postgresql instances.
    * `instance_id` - The id of the postgresql instance.
    * `instance_name` - The name of the postgresql instance.
    * `vpc` - The id of the vpc.


## Import
RdsPostgresqlAllowlist can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_allowlist.default resource_id
```

