---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_allowlist_version_upgrade"
sidebar_current: "docs-volcengine-resource-rds_postgresql_allowlist_version_upgrade"
description: |-
  Provides a resource to manage rds postgresql allowlist version upgrade
---
# volcengine_rds_postgresql_allowlist_version_upgrade
Provides a resource to manage rds postgresql allowlist version upgrade
## Example Usage
```hcl
resource "volcengine_rds_postgresql_allowlist_version_upgrade" "example" {
  instance_id = "postgres-72715e0d9f58"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the postgresql instance to upgrade allowlist version.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



