---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_customized_cfg"
sidebar_current: "docs-volcengine-resource-alb_customized_cfg"
description: |-
  Provides a resource to manage alb customized cfg
---
# volcengine_alb_customized_cfg
Provides a resource to manage alb customized cfg
## Example Usage
```hcl
resource "volcengine_alb_customized_cfg" "foo" {
  customized_cfg_name    = "acc-test-cfg1"
  description            = "This is a test modify"
  customized_cfg_content = "proxy_connect_timeout 4s;proxy_request_buffering on;"
  project_name           = "default"
}
```
## Argument Reference
The following arguments are supported:
* `customized_cfg_content` - (Required) The content of CustomizedCfg. The length cannot exceed 4096 characters. Spaces and semicolons need to be escaped. Currently supported configuration items are `ssl_protocols`, `ssl_ciphers`, `client_max_body_size`, `keepalive_timeout`, `proxy_request_buffering` and `proxy_connect_timeout`.
* `customized_cfg_name` - (Required) The name of CustomizedCfg.
* `description` - (Optional) The description of CustomizedCfg.
* `project_name` - (Optional) The project name of the CustomizedCfg.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
AlbCustomizedCfg can be imported using the id, e.g.
```
$ terraform import volcengine_alb_customized_cfg.default ccfg-3cj44nv0jhhxc6c6rrtet****
```

