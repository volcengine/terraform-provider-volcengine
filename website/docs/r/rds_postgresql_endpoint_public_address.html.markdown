---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_endpoint_public_address"
sidebar_current: "docs-volcengine-resource-rds_postgresql_endpoint_public_address"
description: |-
  Provides a resource to manage rds postgresql endpoint public address
---
# volcengine_rds_postgresql_endpoint_public_address
Provides a resource to manage rds postgresql endpoint public address
## Example Usage
```hcl
resource "volcengine_rds_postgresql_endpoint_public_address" "example" {
  instance_id = "postgres-ac541555dd74"
  endpoint_id = "postgres-ac541555dd74-cluster"
  eip_id      = "eip-1c0x0ehrbhb7k5e8j71k84ryd"
}
```
## Argument Reference
The following arguments are supported:
* `eip_id` - (Required) EIP ID to bind for public access.
* `endpoint_id` - (Required) Endpoint ID.
* `instance_id` - (Required) The ID of the RDS PostgreSQL instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsPostgresqlEndpointPublicAddress can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_endpoint_public_address.default resource_id
```

