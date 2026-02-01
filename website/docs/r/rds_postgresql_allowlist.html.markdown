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
  allow_list      = ["10.0.0.0/24"]
  security_group_bind_infos {
    security_group_id = "sg-1jojfhw8rca9s1n7ampztrq6w"
    bind_mode         = "IngressDirectionIp"
  }
}
resource "volcengine_rds_postgresql_allowlist" "example" {
  instance_ids    = ["postgres-72715e0d9f58", "postgres-eb3a578a6d73"]
  allow_list_name = "unify_new"
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_name` - (Required) The name of the postgresql allow list.
* `allow_list_category` - (Optional) The category of the allow list. Valid values: Ordinary, Default. When this parameter is used as a request parameter, there is no default value.
* `allow_list_desc` - (Optional) The description of the postgresql allow list.
* `allow_list_type` - (Optional, ForceNew) The type of IP address in the whitelist. Currently only `IPv4` addresses are supported.
* `allow_list` - (Optional) Enter an IP address or a range of IP addresses in CIDR format. This field cannot be used together with the user_allow_list field.
* `instance_ids` - (Optional) IDs of PostgreSQL instances to unify allowlists. When set, creation uses UnifyNewAllowList to merge existing instance allowlists into a new one. Supports merging and generating allowlists of up to 300 instances.
* `security_group_bind_infos` - (Optional) The information of security groups to bind with the allow list.
* `update_security_group` - (Optional) Whether to update the security groups bound to the allowlist when modifying.
* `user_allow_list` - (Optional) IP addresses outside security groups to be added to the allowlist. Cannot be used with allow_list.

The `security_group_bind_infos` object supports the following:

* `bind_mode` - (Required) The binding mode of the security group. Valid values: IngressDirectionIp, AssociateEcsIp.
* `security_group_id` - (Required) The ID of the security group.
* `ip_list` - (Optional) IP addresses in the security group.
* `security_group_name` - (Optional) The name of the security group.

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

