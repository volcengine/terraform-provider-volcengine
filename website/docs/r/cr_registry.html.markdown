---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_registry"
sidebar_current: "docs-volcengine-resource-cr_registry"
description: |-
  Provides a resource to manage cr registry
---
# volcengine_cr_registry
Provides a resource to manage cr registry
## Example Usage
```hcl
# create cr registry
resource "volcengine_cr_registry" "foo" {
  name               = "acc-test-cr"
  delete_immediately = false
  password           = "1qaz!QAZ"
  project            = "default"
}

# create cr namespace
resource "volcengine_cr_namespace" "foo" {
  registry = volcengine_cr_registry.foo.id
  name     = "acc-test-namespace"
  project  = "default"
}

# create cr repository
resource "volcengine_cr_repository" "foo" {
  registry     = volcengine_cr_registry.foo.id
  namespace    = volcengine_cr_namespace.foo.name
  name         = "acc-test-repository"
  description  = "A test repository created by terraform."
  access_level = "Public"
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required, ForceNew) The name of registry.
* `delete_immediately` - (Optional) Whether delete registry immediately. Only effected in delete action.
* `password` - (Optional) The password of registry user.
* `project` - (Optional) The ProjectName of the cr registry.
* `proxy_cache_enabled` - (Optional, ForceNew) Whether to enable proxy cache.
* `proxy_cache` - (Optional, ForceNew) The proxy cache of registry. This field is valid when proxy_cache_enabled is true.
* `resource_tags` - (Optional, ForceNew) Tags.
* `type` - (Optional, ForceNew) The type of registry. Valid values: `Enterprise`, `Micro`. Default is `Enterprise`.

The `proxy_cache` object supports the following:

* `type` - (Required, ForceNew) The type of proxy cache. Valid values: `DockerHub`, `DockerRegistry`.
* `endpoint` - (Optional, ForceNew) The endpoint of proxy cache.
* `password` - (Optional, ForceNew) The password of proxy cache.
* `skip_ssl_verify` - (Optional, ForceNew) Whether to skip ssl verify.
* `username` - (Optional, ForceNew) The username of proxy cache.

The `resource_tags` object supports the following:

* `key` - (Required, ForceNew) The Key of Tags.
* `value` - (Required, ForceNew) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `charge_type` - The charge type of registry.
* `create_time` - The creation time of registry.
* `domains` - The domain of registry.
    * `domain` - The domain of registry.
    * `type` - The domain type of registry.
* `status` - The status of registry.
    * `conditions` - The condition of registry.
    * `phase` - The phase status of registry.
* `user_status` - The status of user.
* `username` - The username of cr instance.


## Import
CR Registry can be imported using the name, e.g.
```
$ terraform import volcengine_cr_registry.default enterprise-x
```

