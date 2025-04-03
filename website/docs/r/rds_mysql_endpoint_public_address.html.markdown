---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_endpoint_public_address"
sidebar_current: "docs-volcengine-resource-rds_mysql_endpoint_public_address"
description: |-
  Provides a resource to manage rds mysql endpoint public address
---
# volcengine_rds_mysql_endpoint_public_address
Provides a resource to manage rds mysql endpoint public address
## Example Usage
```hcl
resource "volcengine_rds_mysql_endpoint_public_address" "foo" {
  eip_id      = "eip-rrq618fo9c00v0x58s4r6ky"
  endpoint_id = "mysql-38c3d4f05f6e-custom-01b0"
  instance_id = "mysql-38c3d4f05f6e"
  domain      = "mysql-38c3d4f05f6e-test-01b0-public.rds.volces.com"
}
```
## Argument Reference
The following arguments are supported:
* `eip_id` - (Required, ForceNew) The id of the eip.
* `endpoint_id` - (Required, ForceNew) The id of the endpoint.
* `instance_id` - (Required, ForceNew) The id of mysql instance.
* `domain` - (Optional) The domain.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsMysqlEndpointPublicAddress can be imported using the instance id, endpoint id and eip id, e.g.
```
$ terraform import volcengine_rds_mysql_endpoint_public_address.default instanceId:endpointId:eipId
```

