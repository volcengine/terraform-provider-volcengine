---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_ssl"
sidebar_current: "docs-volcengine-resource-rds_postgresql_instance_ssl"
description: |-
  Provides a resource to manage rds postgresql instance ssl
---
# volcengine_rds_postgresql_instance_ssl
Provides a resource to manage rds postgresql instance ssl
## Example Usage
```hcl
resource "volcengine_rds_postgresql_instance_ssl" "example" {
  instance_id      = "postgres-72715e0d9f58"
  ssl_enable       = true
  force_encryption = true
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the postgresql Instance.
* `force_encryption` - (Optional) Whether to enable force encryption. This only takes effect when the SSL encryption function of the instance is enabled.
* `reload_ssl_certificate` - (Optional) Update the validity period of the SSL certificate. This only takes effect when the SSL encryption function of the instance is enabled. It is not supported to pass in reload_ssl_certificate and ssl_enable at the same time.
* `ssl_enable` - (Optional) Whether to enable SSL.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsPostgresqlInstanceSsl can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_instance_ssl.default resource_id
```

